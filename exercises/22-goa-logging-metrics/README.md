# Add Logging and Metrics

Add logging and metrics reporting to the goa app built in the previous exercise.

## Logging Packages

Use one of:

* [log15](https://github.com/inconshreveable/log15)
* [go-kit](https://github.com/go-kit/kit/tree/master/log)
* [logrus](https://github.com/Sirupsen/logrus)

Setup the service using `WithLogger`. See
[http://goa.design/implement/logging/](http://goa.design/implement/logging/).

## Metrics

Use graphite and grafana to trace metrics reported by statsd.
[https://github.com/kamon-io/docker-grafana-graphite](https://github.com/kamon-io/docker-grafana-graphite)
contains a nice self-contained docker image:

```
docker run \
  --detach \
   --publish=80:80 \
   --publish=81:81 \
   --publish=8125:8125/udp \
   --publish=8126:8126 \
   --name kamon-grafana-dashboard \
   kamon/grafana_graphite
```

Setup the metrics sink to use `statsd` running on localhost on port `8125` if using the docker image
above. See
[https://github.com/goadesign/goa/blob/master/metrics.go](https://github.com/goadesign/goa/blob/master/metrics.go).

goa sends metrics for each error and response by default.
