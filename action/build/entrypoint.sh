#!/bin/sh
# Build the application and test output function

GOOS=linux CGO_ENABLED=0 GOARCH=amd64
go mod download
go build -ldflags "-X github.com/troy0820/secretkube/version.Version=v0.0.1" -o secretkube

./secretkube output -f "testdata/json.json" -o output.yaml -n secret -s secret

./secretkube output -f "testdata/json.json" -n secret -s secret

