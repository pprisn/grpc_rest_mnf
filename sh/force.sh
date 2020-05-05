#!/bin/sh
migrate -path ../migrations -database "postgres://localhost/mnf?sslmode=disable" force 20200504153203

