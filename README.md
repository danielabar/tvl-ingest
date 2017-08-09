# tvl-ingest

> TTC Vehicle Locations ingest from real-time API

## Basic idea
- Query `vehicleLocations` command every x minutes (TBD), will require one query for every route we're interested in (`lastTime` from `nth` request becomes the value of `t` parameter for `nth+1` request, for first request `t` is `0`)
- Generate ES docs for each result, making sure to stamp each with `lastTime`, this is critical because it represents a "batch" of vehicles that were in particular locations at the SAME point in time
- Save above generated docs to the `raw` index

Elasticsearch with [official docker image](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html#docker-cli-run-prod-mode)

Note official image comes with basic auth elastic/changeme, for development/hacking purposes, disable it or get a different image that doesn't have it.

```
docker pull docker.elastic.co/elasticsearch/elasticsearch:5.5.1
docker volume create --name esdata
docker run -p 9200:9200 -e "http.host=0.0.0.0" -e "transport.host=127.0.0.1" -v esdata:/user/share/elasticsearch/data docker.elastic.co/elasticsearch/elasticsearch:5.5.1
```

TODO: Save ingested data with container so it can be shared for development purposes.
