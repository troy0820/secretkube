SHELL:=/bin/bash

all: secret

test: secret
	go test ./... -v

secret:
	go build -mod=vendor -ldflags "-X github.com/troy0820/secretkube/version.Version=0.0.1" -o secret

clean:
	rm -rf secret* output*

.PHONY: all test secret clean
