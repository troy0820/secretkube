#!/bin/sh
# Build the application and test output function

go mod download

go build -o secretkube ./cmd/secretkube

./secretkube output -f "testdata/json.json" -o output.yaml -n secret -s secret

./secretkube output -f "testdata/json.json" -n secret -s secret

