#!/bin/sh
migrate -path ../migrations -database "postgres://localhost/mnf?sslmode=disable" down

