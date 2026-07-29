# Message Broker Structure

## Example message published to RabbitMQ

```json
{
  "driver_id": "<random driver id from PostgreSQL db>",
  "coordinates": {
    "lat": "<random float64 for example 47.842658>",
    "lon": "<random float64 for example 34.811989>",
    "created_at": "<example 0001-01-01 00:00:00>"
  },
  "order": {
    "id": "<random order id from db>",
    "status": "<created, in_progress, done>"
  }
}
```

    