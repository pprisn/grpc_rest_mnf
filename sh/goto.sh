#!/bin/sh
migrate -path ../migrations -database "postgres://localhost/mnf?sslmode=disable" force 2020054121916

