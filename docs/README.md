You need to create db with docker

```
docker run -d \
  --name taxi-service-db \
  -e POSTGRES_USER=ukique \
  -e POSTGRES_PASSWORD=taxi \
  -e POSTGRES_DB=taxi_service_db \
  -p 5435:5432 \
  postgres
```