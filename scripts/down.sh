#!/bin/bash

source ../.env
cd ../sql/schema
goose postgres "$DATABASE_URL" down

docker exec -it redis-stack redis-cli flushdb