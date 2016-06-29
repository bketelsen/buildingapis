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

9. **BREAK**

9. Middleware as a concept - Brian
  - Build sample middleware for logging request headers (e.g.)
10. Alternative Muxers - Brian
11. Middleware packages (Alice) - Brian
12. Authorization - Raphael *(IN PROGRESS)*
  - Basic Auth, Shared secrets, JWT, OAuth2
  - Add basic auth to sample
13. Security - Raphael
  - CSRF
  - XSS
  - SQL injection

13. **LUNCH**

14. Frameworks - Raphael *(IN PROGRESS)*
  - Why - helps at scale
  - Costs
  - Echo, Gin, Goji <- "standard frameworks"
15. Replicate API built in step 6 in any of the above framework - Together
16. Codegen - why? - Brian
17. DIY codegen, goa, go-swagger - Brian
18. (re)build app using goa - Together

18. **BREAK**

19. Running APIs in production - Brian
  - Logging! (principles)
  - Metrics
  - Tracing
20. Add logging/metric data to API built with goa - Together
21. Load Balancing w/ nginx - Brian
  - Show nginx config for app
22. Backplane LBs - Brian
23. Running in containers - Brian
  - Logging (docker logs, syslog)
  - Monitoring (collectd, 
24. "Dockerize" app built in previous steps - Together
  - With logging
  - And nginx
25. Microservices - Raphael
  - What are they?
  - Why and why not?
26. Advanced topics: Vagrant flux - linkerd - Raphael
27. Open Q&A - Together

TIMING

0. 9:00-9:05 B
1. 9:05-9:15 B
2. 9:15-9:20 R
3. 9:20-9:35 B
4. 9:35-9:50 R
5. 9:50-10:00 R
6. 10:00-10:15 *Exercise* R
7. 10:15-10:25 B
8. 10:25-10:45 *Exercise* B
9. BREAK 10:45-11:00
9. 11:00-11:15 *Exercise* B
10. 11:15-11:25 B
11. 11:25-11:35 B
12. 11:35-11:45 *Execercise* R
13. 11:45:11:55 R
13. LUNCH 12:00-13:00
14. 13:00-13:05 R
15. 13:05-13:25 *Exercise* B+R
16. 13:25-13:35 B
17. 13:35-13:45 B
18. 13:45-14:10 *Exercise* R
19. 14:10-14:25 B
20. 14:25-14:40 *Exercise* B
20. BREAK 14:40-15:00
21. 15:05:15:15 B
22. 15:15-15:25 B
23. 15:25-15:40 B
24. 15:40-16:05 *Exercise* R
25. 16:05-16:20 R
26. 16:20-16:35 R
27. 16:35-17:00 B+R
