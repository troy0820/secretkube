# Secret Kube
[![Travis CI](https://img.shields.io/travis/troy0820/secretkube.svg?style=flat-square)](https://travis-ci.org/troy0820/secretkube)
![](https://img.shields.io/github/issues/troy0820/secretkube.svg?style=flat-square)
[![Go Report Card](https://goreportcard.com/badge/github.com/troy0820/secretkube?style=flat-square)](https://goreportcard.com/report/github.com/troy0820/secretkube)

_Name is a work in progress_
## Table of Contents

_(Table of contents goes here)_


This tool will allow you to keep the keys in your json file and base64 the contents and apply them to your Kubernetes cluster.

Usually when you create a secret in Kubernetes, you will do the following:

`kubectl create secret generic <secret-name> --from-file=<file-name>`

The key of your secret will be the file name and everything will be base64 encoded.  This becomes a problem when you want to take a JSON file and have the keys for your secret, be the keys not base64 encoded but the contents base64 encoded.

### How to install

`go get -u github.com/troy0820/secretkube`

### Dependencies vendored with ~~dep~~ `go 1.11 modules`

