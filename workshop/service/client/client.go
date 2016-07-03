package client

import (
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
)

// Client is the GoWorkshop service client.
type Client struct {
	*goaclient.Client
	JWTAuthSigner   goaclient.Signer
	BasicAuthSigner goaclient.Signer
	Encoder         *goa.HTTPEncoder
	Decoder         *goa.HTTPDecoder
}

// New instantiates the client.
func New(c goaclient.Doer) *Client {
	client := &Client{
		Client:  goaclient.New(c),
		Encoder: goa.NewHTTPEncoder(),
		Decoder: goa.NewHTTPDecoder(),
	}

	// Setup encoders and decoders
	client.Encoder.Register(goa.NewJSONEncoder, "application/json")
	client.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	client.Encoder.Register(goa.NewJSONEncoder, "*/*")
	client.Decoder.Register(goa.NewJSONDecoder, "*/*")

	return client
}

// SetJWTAuthSigner sets the request signer for the JWTAuth security scheme.
func (c *Client) SetJWTAuthSigner(signer goaclient.Signer) {
	c.JWTAuthSigner = signer
}

// SetBasicAuthSigner sets the request signer for the BasicAuth security scheme.
func (c *Client) SetBasicAuthSigner(signer goaclient.Signer) {
	c.BasicAuthSigner = signer
}
