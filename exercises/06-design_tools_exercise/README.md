# Exercise: Design the GoWorkshop API using Swagger

The GoWorkshop API consists of two resources:

* `courses` represent a specific workshop course with start and end times and a location.
* `registrations` represent a registration to a course with details about the attendee.

It should be possible to create, list and delete courses and create and list registrations.

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
