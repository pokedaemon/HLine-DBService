#!/bin/sh

migrate -database "postgres://localhost/reverseshoptest?sslmode=disable" -path migrations up
