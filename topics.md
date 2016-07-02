### Topics

9:00-9:05 Intro - Agenda - Mode of operation (ask questions etc.)
 
1. 9:05-9:15 HTTP Basics (+https, standard headers) - Brian
2. 9:15-9:20 HTTP2 - Raphael *(DONE)*
3. 9:20-9:35 Different API styles -> REST - Brian *(Mostly Done)* needs examples
4. 9:35-9:50 REST Design consideration - Raphael *(DONE)*
5. 9:50-10:00 Tools for designing APIs, Swagger, RAML, JSONSchema, etc. - Raphael *(DONE)*
6. 10:00-10:15 Design API that we will build later using Swagger editor - Raphael/Together *(IN PROGRESS)*
7. 10:15-10:25 HTTP Go server (brief) - Brian
  - HTTP server "architecture" one goroutine per request, request flow
  - Routing
  - Error Handling
  - Encoding/Decoding + MIME
8. 10:25-10:45 Build a REST server in Go using 100% stdlib - Brian
  - Error handling
  - Testing
  ....

9. 10:45-11:00 **BREAK**

10. 11:00-11:15 Middleware as a concept - Brian
  - Build sample middleware for logging request headers (e.g.)
11. 11:15-11:25 Alternative Muxers - Brian
12. 11:25-11:35 Middleware packages (Alice) - Brian
13. 11:35-11:45 Authorization - Raphael *(IN PROGRESS)*
  - Basic Auth, Shared secrets, JWT, OAuth2
  - Add basic auth to sample
14. 11:45-11:50 Security - Raphael
  - CSRF
  - XSS
  - SQL injection

15. 12:00-13:00 **LUNCH**

16. 13:00-13:05 Frameworks - Raphael *(IN PROGRESS)*
  - Why - helps at scale
  - Costs
  - Echo, Gin, Goji <- "standard frameworks"
17. 13:05-13:25 Replicate API built in step 6 in any of the above framework - ?/Together
18. 13:25-13:35 Codegen - why? - Brian
20. 13:35-13:45 DIY codegen, goa, go-swagger - Brian
21. 13:45-14:10 (re)build app using goa - Raphael/Together
22. 14:10-14:25 Running APIs in production - Brian
  - Logging! (principles)
  - Metrics
  - Tracing
23. 14:25-14:40 Add logging/metric data to API built with goa - Raphael/Together

24. 14:40-15:00 **BREAK**

25. 15:05-15:15 Load Balancing w/ nginx - Brian
  - Show nginx config for app
26. 15:15-15:25 Backplane LBs - Brian
27. 15:25-15:40 Running in containers - Brian
  - Logging (docker logs, syslog)
  - Monitoring (collectd, 
28. 15:40-16:05 "Dockerize" app built in previous steps - Brian/Together
  - With logging
  - And nginx
29. 16:05-16:20 Microservices - Raphael
  - What are they?
  - Why and why not?
30. 16:20-16:35 Advanced topics: Vagrant flux - linkerd - Raphael
31. 16:35-17:00 Open Q&A - Together
