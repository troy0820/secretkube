SHELL:=/bin/bash

VERSION := $(shell cat VERSION.txt)

all: secret

test: secret
	go test ./... -v

secret-local:
	go build -mod=vendor -ldflags "-X github.com/troy0820/secretkube/version.Version=$(VERSION)" -o secretkube ./cmd/secretkube

release:
	git archive --format=tar -v HEAD | gzip >secretkube-$(VERSION).tar.gzip
	shasum -a 256 secretkube-$(VERSION).tar.gzip > secretkube-$(VERSION).tar.gzip.sha256
	cat secretkube-$(VERSION).tar.gzip.sha256 | shasum -c

secret:
	go build -ldflags "-X github.com/troy0820/secretkube/version.Version=$(VERSION)" -o secretkube ./cmd/secretkube

clean:
	rm -rf secret* output* *.tar

.PHONY: all test secret-local secret release clean
