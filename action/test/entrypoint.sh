#!/bin/bash
# Test the Golang application
rm go.sum

go get ./...

go test ./... -v

