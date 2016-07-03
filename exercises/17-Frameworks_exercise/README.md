# Exercise: Implement the GoWorkshop API

## Choice of Frameworks

* [Gin](https://github.com/gin-gonic/gin)
* [Echo](https://github.com/labstack/echo)
* [Goji](https://github.com/goji/goji)

Endpoints are:

courses:
* create a course `POST /courses`
* show a course `GET /course`
* delete a course `DELETE /courses/:id`

registrations:
* create a registration `POST /registrations`
* show a registration `GET /registration/:id`
* list all registrations `GET /registrations`

## Database

The `library` package implements a in-memory database:

```go
	import "github.com/bketelsen/buildingapis/exercises/library"
```

Create a new database with:

```go
	db := library.NewDB()
```

The database contains pre-seeded data. To create an empty database use:

```go
	db := library.NewEmptyDB()
```

### Insert

```go
	model := &library.RegistrationModel{
		ID:        library.NewID(),
		CourseID:  courseID,
		FirstName: firstName,
		LastName:  lastName,
		Address:   address,
	}
	err := db.Insert("registrations", model)
```

### Get

```go
	list, err := db.List("registrations", "id", strconv.Itoa(id))
```

### List

```go
	list, err := db.List("registrations", "id", nil)
```

### Delete

```go
	err := db.Delete("courses", "id", strconv.Itoa(id))
```

## Reminder:

The GoWorkshop API consists of two resources:

* `courses` represent a specific workshop course with start and end times and a location.
* `registrations` represent a registration to a course with details about the attendee.

It should be possible to create, show and delete courses and create, show and list registrations.

The course type should have the following fields:

|   Name        |Type      |Description          |
|---------------|----------|---------------------|
| `id`          | Integer  | Course identifier   |
| `href`        | String   | API relative URI    |
| `name`        | String   | Course name         |
| `description` | String   | Description         |
| `start_time`  | DateTime | Start date and time |
| `end_time`    | DateTime | End date and time   |
| `location`    | String   | Location (city)     |

The registration type should have the following fields:

|   Name       |Type      |Description          |
|--------------|----------|---------------------|
| `id`         | Integer  | Identifier          |
| `href`       | String   | API relative URI    |
| `first_name` | String   | Course name         |
| `last_name`  | String   | Description         |
| `address`    | Address  | Start date and time |

The Address type should have the following fields:

|   Name   |Type      |Description    |
|----------|----------|---------------|
| `number` | Integer  | Street Number |
| `street` | String   | Street Name   |
| `city`   | String   | City name     |
| `state`  | String   | US State code |
| `zip`    | Integer  | US Zip code   |
