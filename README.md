# Secret Kube

![](https://img.shields.io/github/issues/troy0820/secretkube.svg?style=flat-square)
[![Go Report Card](https://goreportcard.com/badge/github.com/troy0820/secretkube?style=flat-square)](https://goreportcard.com/report/github.com/troy0820/secretkube)
![](https://github.com/troy0820/secretkube/workflows/Go%20build%20and%20Output/badge.svg)
![](https://github.com/troy0820/secretkube/workflows/Go%20test/badge.svg)

Create Kubernetes secrets from JSON files with automatic base64 encoding. This tool preserves your JSON keys as secret keys while automatically encoding the values.

## Table of Contents

- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)
  - [Create a Secret](#create-a-secret)
  - [Output a Secret](#output-a-secret)
  - [Get a Secret](#get-a-secret)
  - [Watch Secrets](#watch-secrets)
  - [Inform about Secrets](#inform-about-secrets)
- [Examples](#examples)
- [Building from Source](#building-from-source)
- [Contributing](#contributing)
- [License](#license)

## Overview

This tool allows you to keep the keys in your JSON file and base64 encode the contents before applying them to your Kubernetes cluster.

Usually when you create a secret in Kubernetes, you will do the following:

```bash
kubectl create secret generic <secret-name> --from-file=<file-name>
```

The key of your secret will be the file name and everything will be base64 encoded. This becomes a problem when you want to take a JSON file and have the keys for your secret be the JSON keys (not base64 encoded) but with the contents base64 encoded.

SecretKube solves this by:
- Reading JSON files with key/value pairs
- Preserving the JSON keys as secret keys
- Automatically base64 encoding only the values
- Applying the secrets directly to your Kubernetes cluster

## Installation

### Using Go Install

```bash
go install github.com/troy0820/secretkube/cmd/secretkube@latest
```

### Using Go Get (deprecated but still works)

```bash
go get -u github.com/troy0820/secretkube/cmd/secretkube
```

### From Source

```bash
git clone https://github.com/troy0820/secretkube.git
cd secretkube
make secret
```

## Usage

### Create a Secret

Create a Kubernetes secret from a JSON file and apply it to your cluster:

```bash
secretkube create -f <json-file> -n <secret-name> [-s <namespace>] [-c <kubeconfig>]
```

**Flags:**
- `-f, --file` (required): Path to JSON file
- `-n, --name` (required): Name of the secret
- `-s, --namespace`: Namespace to create the secret in (default: "default")
- `-c, --config`: Path to kubeconfig file (default: "$HOME/.kube/config")

### Output a Secret

Generate a Kubernetes secret YAML from a JSON file without applying it:

```bash
secretkube output -f <json-file> -n <secret-name> [-s <namespace>] [-o <output-file>]
```

**Flags:**
- `-f, --file` (required): Path to JSON file
- `-n, --name` (required): Name of the secret
- `-s, --namespace`: Namespace for the secret (default: "default")
- `-o, --output`: Output file to save the secret YAML

### Get a Secret

Retrieve and display a secret from your cluster:

```bash
secretkube get [-n <secret-name>]
```

**Flags:**
- `-n, --name`: Name of the secret to retrieve

### Watch Secrets

Watch for changes to secrets in a namespace:

```bash
secretkube watch [-n <namespace>]
```

**Flags:**
- `-n, --namespace`: Namespace to watch (default: all namespaces)

### Inform about Secrets

Get information about secrets using an informer pattern:

```bash
secretkube inform
```

## Examples

### Example JSON File

Create a file named `secrets.json`:

```json
{
  "database_url": "postgres://user:pass@localhost:5432/mydb",
  "api_key": "my-secret-api-key",
  "oauth_token": "abc123xyz789"
}
```

### Create and Apply Secret

```bash
secretkube create -f secrets.json -n my-app-secrets -s production
```

This will:
1. Read the JSON file
2. Base64 encode each value
3. Create a Kubernetes secret with keys `database_url`, `api_key`, and `oauth_token`
4. Apply the secret to the `production` namespace

### Generate Secret YAML

```bash
secretkube output -f secrets.json -n my-app-secrets -o secret.yaml
```

This will generate a `secret.yaml` file that you can review or apply manually with `kubectl apply -f secret.yaml`.

## Building from Source

### Requirements

- Go 1.25.0 or later
- Access to a Kubernetes cluster (for testing)

### Build

```bash
make secret
```

### Test

```bash
make test
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

Copyright (c) 2018-2026 Troy Connor

