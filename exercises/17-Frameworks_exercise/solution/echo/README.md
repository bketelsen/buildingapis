# GoWorkshop - Echo Version

## Usage:

Compile and run the service:
```
go build -o gowo
./gowo
```

The commands below use the [httpie](https://github.com/jkbrzt/httpie) tool to call the service.

This set of commands creates a new course, shows it and then deletes it. The last command shows the
course one more time to prove it was deleted.

```
http POST localhost:8000/courses name=new_course description="awesome content" start_time=2016-11-02T09:00:00-08:00 end_time=2016-11-02T17:00:00-08:00 location="Santa Barbara"
http localhost:8000/courses/10 -v
http DELETE localhost:8000/courses/10 -v
http localhost:8000/courses/10 -v
```

This set of commands list all existing registrations, create a new one, list all registrations again
to show the newly created registration and finally show the newly created registration.

```
http localhost:8000/registrations -v
http POST localhost:8000/registrations first_name=me last_name=here address:='{"number":43,"street":"Some Street","city":"Denver","state":"CO","zip":80205}' course_href=/courses/1 -v
http localhost:8000/registrations -v
http localhost:8000/registrations/11 -v
```
