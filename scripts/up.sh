#!/bin/bash

cd ../sql/schema

goose postgres "postgres://postgres:admin@localhost:5432/test?sslmode=disable" up