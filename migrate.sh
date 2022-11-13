#!/bin/bash

export $(xargs < .env)
migrate -path ./schema -database "postgres://postgres:${POSTGRES_PASSWORD}@localhost:5436/postgres?sslmode=disable" up
