package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// LoginSessionPath computes a request path to the login action of session.
func LoginSessionPath() string {
	return fmt.Sprintf("/api/token")
}

// Login uses basic auth and on successful auth returns a JWT in the response "Authorization" header
func (c *Client) LoginSession(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewLoginSessionRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewLoginSessionRequest create the request corresponding to the login action endpoint of the session resource.
func (c *Client) NewLoginSessionRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.BasicAuthSigner != nil {
		c.BasicAuthSigner.Sign(req)
	}
	return req, nil
}
