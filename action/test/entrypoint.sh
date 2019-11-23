#!/bin/sh
# Build the application and test output function
go mod download

go test ./... -v
