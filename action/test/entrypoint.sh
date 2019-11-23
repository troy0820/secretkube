#!/bin/sh
# Test the Golang application
go mod download
go test ./... -v

