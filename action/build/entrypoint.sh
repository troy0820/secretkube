#!/bin/bash
# Build the application and test output function
rm go.sum

go get ./...

go build -ldflags "-X github.com/troy0820/secretkube/version.Version=v0.0.1" -o secretkube

./secretkube output -f "testdata/json.json" -o output.yaml -n secret -s secret

./secretkube output -f "testdata/json.json" -n secret -s secret

