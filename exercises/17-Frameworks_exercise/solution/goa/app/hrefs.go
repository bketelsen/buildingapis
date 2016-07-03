//************************************************************************//
// API "GoWorkshop": Application Resource Href Factories
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/bketelsen/buildingapis/exercises/16-Frameworks_exercise/solution/goa/design
// --out=$(GOPATH)/src/github.com/bketelsen/buildingapis/exercises/16-Frameworks_exercise/solution/goa
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "fmt"

// CourseHref returns the resource href.
func CourseHref(id interface{}) string {
	return fmt.Sprintf("/courses/%v", id)
}

// RegistrationHref returns the resource href.
func RegistrationHref(id interface{}) string {
	return fmt.Sprintf("/registrations/%v", id)
}
