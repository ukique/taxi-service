# Message Broker Structure

## Example message published to RabbitMQ

```json
{
  "driver_id": "<random driver id from MySQL db>",
  "coordinates": {
    "lat": "<random float64 for example 47.842658>",
    "lon": "<random float64 for example 34.811989>"
  },
  "order": {
    "id": "<random order id from MySQL db>",
    "status": "<created, in_progress, done>"
  }
}
```

