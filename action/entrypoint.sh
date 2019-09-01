#!/bin/bash
# Test the Golang application
rm go.sum

go get ./...

go test ./... -v

go build -ldflags "-X github.com/troy0820/secretkube/version.Version=v0.0.1" -o secretkube

./secretkube output -f "testdata/json.json" -o output.yaml -n troy -s troyboy
