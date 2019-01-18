SHELL:=/bin/bash

all: secret

test: secret
	go test ./... -v

secret-local:
	go build -mod=vendor -ldflags "-X github.com/troy0820/secretkube/version.Version=0.0.1" -o secret


secret:
	go build -ldflags "-X github.com/troy0820/secretkube/version.Version=0.0.1" -o secret

clean:
	rm -rf secret* output*

.PHONY: all test secret-local secret clean
