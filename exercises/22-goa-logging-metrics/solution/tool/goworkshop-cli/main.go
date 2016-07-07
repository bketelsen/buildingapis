package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/context"

	"github.com/bketelsen/buildingapis/exercises/22-goa-logging-metrics/solution/client"
	"github.com/bketelsen/buildingapis/exercises/22-goa-logging-metrics/solution/tool/cli"
	"github.com/spf13/cobra"
)

func main() {
	// Create command line parser
	app := &cobra.Command{
		Use:   "GoWorkshop-cli",
		Short: `CLI client for the GoWorkshop service`,
	}

	// Create client struct
	httpClient := newHTTPClient()
	c := client.New(httpClient)

	// Register global flags
	app.PersistentFlags().StringVarP(&c.Scheme, "scheme", "s", "", "Set the requests scheme")
	app.PersistentFlags().StringVarP(&c.Host, "host", "H", "localhost:8080", "API hostname")
	app.PersistentFlags().DurationVarP(&httpClient.Timeout, "timeout", "t", time.Duration(20)*time.Second, "Set the request timeout")
	app.PersistentFlags().BoolVar(&c.Dump, "dump", false, "Dump HTTP request and response.")

	// Initialize API client
	c.UserAgent = "GoWorkshop-cli/1.0"

	// Register API commands
	cli.RegisterCommands(app, c)

	// Register siege command
	siegeCmd := cobra.Command{
		Use:   "siege",
		Short: `Sends lots of traffic to service to showcase logging and metrics`,
		Run:   func(*cobra.Command, []string) { siege(c) },
	}
	app.AddCommand(&siegeCmd)

	// Execute!
	if err := app.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}
}

// siege continuously sends traffic to the service until the process is killed.
func siege(c *client.Client) {
	fmt.Printf("Sending requests continuously to %s, CTRL+C to stop.", c.Host)
	for {
		candidates[rand.Intn(len(candidates))](c)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

var candidates = []func(c *client.Client){createCourse, showCourse, deleteCourse, listRegistrations}

func createCourse(c *client.Client) {
	desc := "course"
	startTime := time.Now().Add(time.Duration(10) * time.Hour)
	payload := client.CreateCoursePayload{
		Name:        "name",
		Description: &desc,
		Location:    "location",
		StartTime:   startTime,
		EndTime:     startTime.Add(time.Duration(8) * time.Hour),
	}
	_, err := c.CreateCourse(context.Background(), client.CreateCoursePath(), &payload, "")
	if err != nil {
		fail("create course failed: %s", err)
	}
}

func showCourse(c *client.Client) {
	p := client.ShowCoursePath(int(math.Mod(float64(rand.Int()), 100)))
	_, err := c.ShowCourse(context.Background(), p)
	if err != nil {
		fail("create course failed: %s", err)
	}
}

func deleteCourse(c *client.Client) {
	p := client.ShowCoursePath(int(math.Mod(float64(rand.Int()), 100)))
	_, err := c.DeleteCourse(context.Background(), p)
	if err != nil {
		fail("create course failed: %s", err)
	}
}

func listRegistrations(c *client.Client) {
	_, err := c.ListRegistration(context.Background(), client.ListRegistrationPath(), nil)
	if err != nil {
		fail("list registrations failed: %s", err)
	}
}

func fail(format string, vals ...interface{}) {
	fmt.Fprintf(os.Stderr, format, vals)
	os.Exit(1)
}

// newHTTPClient returns the HTTP client used by the API client to make requests to the service.
func newHTTPClient() *http.Client {
	// TBD: Change as needed (e.g. to use a different transport to control redirection policy or
	// disable cert validation or...)
	return http.DefaultClient
}
