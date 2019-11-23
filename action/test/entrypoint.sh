#!/bin/sh
# Test the Golang application
GOOS=linux CGO_ENABLED=1 GOARCH=amd64
go mod download
go test ./... -v

