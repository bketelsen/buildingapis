# GoWorkshop

GoWorkshop is a RESTful service that exposes a couple of related resources:

* `courses` represent a specific workshop course with start and end times and a location.
* `registrations` represent a registration to a course with details about the attendee.

The service implementation makes use of the HashiCorp
[memdb](https://github.com/hashicorp/go-memdb) package to store data in memory.

## Security

The APIs are secured with JWT, the `session` resource exposes a `Login` action that can be used to
retrieve tokens in exchange for email and password. The flow is:

1. Client makes request to `/token` using basic auth
2. Client reads content of `Authorization` response header
3. Client sets `Authorization` header of requests to `courses` and `registrations` resources

The database fixtures include a user with email `email` and password `password`.

## Swagger

The `public` resource serves the Swagger definition of the service at `/swagger.json`. It also
defines a CORS policy which lets browsers from any origin retrieve the corresponding JSON file.

This makes it possible to load the Swagger spec in Swagger UI on a development laptop for example.

## Usage

### Start the Service

Build and run the service:

```
go generate
go build -o gow
./gow
```

This should print something like:

```
INFO[06-26|17:19:45] mount                                    ctrl=Course action=Create route="POST /api/courses" security=JWTAuth
INFO[06-26|17:19:45] mount                                    ctrl=Course action=Delete route="DELETE /api/courses/:id" security=JWTAuth
INFO[06-26|17:19:45] mount                                    ctrl=Course action=List route="GET /api/courses" security=JWTAuth
INFO[06-26|17:19:45] mount                                    ctrl=Course action=Patch route="PATCH /api/courses/:id" security=JWTAuth
INFO[06-26|17:19:45] mount                                    ctrl=Course action=Show route="GET /api/courses/:id" security=JWTAuth
INFO[06-26|17:19:45] mount                                    ctrl=Public files=public/ route="GET /swagger/*file"
INFO[06-26|17:19:45] mount                                    ctrl=Public files=swagger/swagger.json route="GET /swagger.json"
INFO[06-26|17:19:45] mount                                    ctrl=Public files=swagger/swagger.yaml route="GET /swagger.yaml"
INFO[06-26|17:19:45] mount                                    ctrl=Public files=public/index.html route="GET /swagger/"
INFO[06-26|17:19:45] mount                                    ctrl=Registration action=Create route="POST /api/registrations" security=JWTAuth
INFO[06-26|17:19:45] mount                                    ctrl=Registration action=List route="GET /api/registrations" security=JWTAuth
INFO[06-26|17:19:45] mount                                    ctrl=Registration action=Patch route="PATCH /api/registrations/:id" security=JWTAuth
INFO[06-26|17:19:45] mount                                    ctrl=Registration action=Show route="GET /api/registrations/:id" security=JWTAuth
INFO[06-26|17:19:45] listen                                   transport=http addr=:8080
```

### Use the Generated Client

Use the client to make requests:

```
cd ./tool/goworkshop-cli
go build -o gowc
./gowc --help
```

Login with:

```
gowc --user=email --password=password login session -v
```

Note the value of the `Authorization` header, use it with:

```
gowc --token=$token list courses
```

See the tool contextual help for calling other APIs, for example:

```
gowc list --help
```

### Read the Docs

Open a browser and point it at `http://localhost:8080/swagger/`. This will serve the swagger
specification of the API.
