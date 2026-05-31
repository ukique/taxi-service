# Taxi Service
A real-time taxi dispatch web service that allows operators to monitor driver locations and manage ride history.

The system tracks driver coordinates as they come in — location updates are saved to the database and pushed to the frontend simultaneously via WebSocket, without one blocking the other. Operators can browse the full coordinates history for any ride as a table directly in the browser.

## Quick Start
1. Clone the repository
```bash
git clone https://github.com/ukique/taxi-service
cd taxi-service/taxi-service
```
2. Run
```bash
docker compose up --build
```
3. Open
[http://localhost](http://localhost)

## Message Broker Structure
#### example of Message Broker Structure you can read [here](https://github.com/ukique/taxi-service/blob/main/docs/examples/message_broker_structure.md)
