### Topics

1. HTTP Basics (+https, standard headers) - Brian
2. HTTP2 - Raphael
3. Different API styles -> REST - Brian
4. REST Design consideration (brief?) - Raphael
5. Tools for designing APIs, Swagger, RAML, JSONSchema, etc. - Raphael
6. Design API that we will build later using Swagger editor - Raphael
5. HTTP Go server (brief) - Brian
  0. HTTP server "architecture" one goroutine per request, request flow
  a. Routing
  b. Error Handling
  c. Encoding/Decoding + MIME
6. Build a REST server in Go using 100% stdlib - Together
  a. Error handling
  b. Testing
  ....
7. Middleware as a concept - Brian
  a. Build sample middleware for logging request headers (e.g.)
8. Alternative Muxers - Brian
9. Middleware packages (Alice) - Brian
10. Frameworks - Raphael
  a. Why - helps at scale
  b. Costs
  c. Echo, Gin, Goji <- "standard frameworks", Iris? HTTP2
11. Replicate API built in step 6 in any of the above framework - Together
12. Codegen - why? - Brian
13. DIY codegen, goa, go-swagger - Brian
14. (re)build app using goa - Together
14. Documenting APIs - Raphael
15. Running APIs in production - Raphael/Brian
  a. Logging!
  b. Metrics
  c. Load balancing
16. Add logging/metric data to API built with goa - Together
17. Backplane LBs - Brian



17. Advanced topics: Vagrant flux - linkerd - Raphael

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

