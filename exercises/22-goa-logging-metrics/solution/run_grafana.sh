#!/bin/bash

docker run \
  --detach \
   --publish=80:80 \
   --publish=81:81 \
   --publish=8125:8125/udp \
   --publish=8126:8126 \
   --name kamon-grafana-dashboard \
   kamon/grafana_graphite

echo "grafana should now be running, open your browser on http://localhost"
echo "Add a data source using Graphite and point it at http://localhost:81"
echo "Add the metric stats.counters.GoWorkshop.goa.response.*.*"
