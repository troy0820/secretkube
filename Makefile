SHELL:=/bin/bash

all: secret

test: secret
	go test ./... -v

secret:
	go build -o secret

clean:
	rm -rf secret* output*

.PHONY: all test secret clean
