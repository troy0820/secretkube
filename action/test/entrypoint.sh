#!/bin/sh
# Build the application and test output function

rm go.sum

go get ./...

go test ./... -v
