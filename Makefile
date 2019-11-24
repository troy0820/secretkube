SHELL:=/bin/bash

VERSION := $(shell cat VERSION.txt)

all: secret

test: secret
	go test ./... -v

secret-local:
	go build -mod=vendor -ldflags "-X github.com/troy0820/secretkube/version.Version=$(VERSION)" -o secretkube ./cmd/secretkube


secret:
	go build -ldflags "-X github.com/troy0820/secretkube/version.Version=$(VERSION)" -o secretkube ./cmd/secretkube

clean:
	rm -rf secret* output*

.PHONY: all test secret-local secret clean
