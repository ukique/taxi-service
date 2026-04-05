# You need to create db with docker

```
docker run -d \
  --name taxi-service-db \
  -e POSTGRES_USER=ukique \
  -e POSTGRES_PASSWORD=taxi \
  -e POSTGRES_DB=taxi_service_db \
  -p 5435:5432 \
  postgres
```

# Message Broker Structure
#### example of Message Broker Structure you can read [here](https://github.com/ukique/taxi-service/blob/main/docs/examples/message_broker_structure.md)