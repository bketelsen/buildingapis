### Topics
* API Design Considerations
* Routing
* Error Handling
* Testing
* Code Generation
* Advanced API Options

* Intro to REST
I'm thinking a discussion around rest and routing. This would start with the basics then progress to show the challenges (flat vs subresource, actions that are not crud)
I guess that all could fall into a larger "API design section"

Would also be good to cover HTTP and the three interesting components: path, headers and body. For both requests and responses. Also explaining some HTTP statuses and common headers.

HTTPS and certs
Let's encrypt

Quick overview of the Golang HTTP server, how it's awesome. Concurrent but still simple.

Asynchronous vs synchronous APIs

Tracing across systems

Logging

Metrics, monitoring etc

Encoding
How content negotiation works

Then all the topics about the frameworks: Request contexts, path parameters, "binding" the body, contrast that with goa's approach in the goa section.

World also be good to talk about what's around HTTP servers: load balancers, SSL termination, proxies and reverse proxies, talk about DNS and DNS round robbin, service discovery, API gateways is a topic in itself

Ah auth all sorts of auth
Might be nice to go through the oauth2 flow as it can be confusing at first but it's actually not that complicated.

In the frameworks section talk about middlewares and go over common ones

HTTP2
Websocket
gRPC
Maybe in encoding cover JSON quickly since it's so omnipresent
JSONP

Describe a SPA architecture

Differences between backend facing and UI facing APIs

