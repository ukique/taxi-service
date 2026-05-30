#!/bin/sh
goose -dir ./migrations postgres "$DATABASE_URL" up
./service