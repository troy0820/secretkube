# Secret Kube [![Build Status](https://travis-ci.org/troy0820/secretkube.svg?branch=master)](https://travis-ci.org/troy0820/secretkube)

_Name is a work in progress_

This tool will allow you to keep the keys in your json file and base64 the contents and apply them to your Kubernetes cluster.

Usually when you create a secret in Kubernetes, you will do the following:

`kubectl create secret generic <secret-name> --from-file=<file-name>`

The key of your secret will be the file name and everything will be base64 encoded.  This becomes a problem when you want to take a JSON file and have the keys for your secret, be the keys not base64 encoded but the contents base64 encoded.

### Dependencies vendored with `dep`

