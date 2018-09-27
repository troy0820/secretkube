SHELL:=/bin/bash

.PHONY: test
test: secret
	go test ./... -v

.PHONY: secret
secret:
	go build -o secret

clean:
	rm -rf secret*

