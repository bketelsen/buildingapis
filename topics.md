### Topics

1. HTTP Basics (+https, standard headers) - Brian
2. HTTP2 - Raphael *(DONE)*
3. Different API styles -> REST - Brian *(Mostly Done)* needs examples
4. REST Design consideration (brief?) - Raphael *(DONE)*
5. Tools for designing APIs, Swagger, RAML, JSONSchema, etc. - Raphael *(IN PROGRESS)*
6. Design API that we will build later using Swagger editor - Raphael *(NEED INPUT)*
5. HTTP Go server (brief) - Brian
  - HTTP server "architecture" one goroutine per request, request flow
  - Routing
  - Error Handling
  - Encoding/Decoding + MIME
6. Build a REST server in Go using 100% stdlib - Together *(NEED INPUT)*
  - Error handling
  - Testing
  ....
7. Middleware as a concept - Brian
  - Build sample middleware for logging request headers (e.g.)
8. Alternative Muxers - Brian
9. Middleware packages (Alice) - Brian
10. Authorization - Raphael *(IN PROGRESS)*
  - Basic Auth, Shared secrets, JWT, OAuth2
  - Add basic auth to sample
11. Security
  - CSRF
  - XSS
  - SQL injection
11. Frameworks - Raphael *(IN PROGRESS)*
  - Why - helps at scale
  - Costs
  - Echo, Gin, Goji <- "standard frameworks", Iris? HTTP2
12. Replicate API built in step 6 in any of the above framework - Together
13. Codegen - why? - Brian
14. DIY codegen, goa, go-swagger - Brian
15. (re)build app using goa - Together
16. Documenting APIs - Raphael
17. Running APIs in production - Raphael/Brian
  a. Logging!
  b. Metrics
  c. Load balancing
18. Add logging/metric data to API built with goa - Together
19. Backplane LBs - Brian
20. Advanced topics: Vagrant flux - linkerd - Raphael
