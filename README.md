# tvl-ingest

TTC Vehicle Locations ingest.

Elasticsearch with [official docker image](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html#docker-cli-run-prod-mode)

```
docker pull docker.elastic.co/elasticsearch/elasticsearch:5.5.1
docker volume create --name esdata
docker run -p 9200:9200 -e "http.host=0.0.0.0" -e "transport.host=127.0.0.1" -v esdata:/user/share/elasticsearch/data docker.elastic.co/elasticsearch/elasticsearch:5.5.1
```

TODO: Save ingested data with container so it can be shared for development purposes.
