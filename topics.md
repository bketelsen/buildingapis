### Topics

0. Intro - Agenda - Mode of operation (ask questions etc.)
1. HTTP Basics (+https, standard headers) - Brian
2. HTTP2 - Raphael *(DONE)*
3. Different API styles -> REST - Brian *(Mostly Done)* needs examples
4. REST Design consideration - Raphael *(DONE)*
5. Tools for designing APIs, Swagger, RAML, JSONSchema, etc. - Raphael *(IN PROGRESS)*
6. Design API that we will build later using Swagger editor - Raphael *(NEED INPUT)*
7. HTTP Go server (brief) - Brian
  - HTTP server "architecture" one goroutine per request, request flow
  - Routing
  - Error Handling
  - Encoding/Decoding + MIME
8. Build a REST server in Go using 100% stdlib - Together *(NEED INPUT)*
  - Error handling
  - Testing
  ....

BREAK

9. Middleware as a concept - Brian
  - Build sample middleware for logging request headers (e.g.)
10. Alternative Muxers - Brian
11. Middleware packages (Alice) - Brian
12. Authorization - Raphael *(IN PROGRESS)*
  - Basic Auth, Shared secrets, JWT, OAuth2
  - Add basic auth to sample
13. Security
  - CSRF
  - XSS
  - SQL injection

LUNCH

14. Frameworks - Raphael *(IN PROGRESS)*
  - Why - helps at scale
  - Costs
  - Echo, Gin, Goji <- "standard frameworks"
15. Replicate API built in step 6 in any of the above framework - Together
16. Codegen - why? - Brian
17. DIY codegen, goa, go-swagger - Brian
18. (re)build app using goa - Together

BREAK

19. Running APIs in production - Raphael/Brian
  a. Logging! (principles)
  b. Metrics
  c. Tracing
20. Add logging/metric data to API built with goa - Together
21. Load Balancing w/ nginx - Brian
  a. Show nginx config for app
22. Backplane LBs - Brian
23. Running in containers - Brian
  a. Logging (docker logs, syslog)
  b. Monitoring (collectd, 
24. "Dockerize" app built in previous steps - Together
  a. With logging
  b. And nginx
25. Microservices - Raphael
  a. What are they?
  b. Why and why not?
26. Advanced topics: Vagrant flux - linkerd - Raphael
27. Open Q&A - Together

0. 9:00-9:05
1. 9:05-9:15
2. 9:15-9:20
3. 9:20-9:35
4. 9:35-9:50
5. 9:50-10:00
6. 10:00-10:15 *Exercise*
7. 10:15-10:25
8. 10:25-10:45 *Exercise*

BREAK 10:45-11:00

9. 11:00-11:15 *Exercise*
10. 11:15-11:25
11. 11:25-11:35
12. 11:35-11:45 *Execercise*
13. 11:45:11:55

LUNCH 12:00-13:00

14. 13:00-13:05
15. 13:05-13:25 *Exercise*
16. 13:25-13:35
17. 13:35-13:45
18. 13:45-14:10 *Exercise*
19. 14:10-14:25
20. 14:25-14:40 *Exercise*

BREAK 14:40-15:00

21. 15:05:15:15
22. 15:15-15:25
23. 15:25-15:40
24. 15:40-16:05 *Exercise*
25. 16:05-16:20
26. 16:20-16:35
27. 16:35-17:00
