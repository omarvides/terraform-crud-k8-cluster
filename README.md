# Terraform crud k8 cluster

This repository will contain a terraform module to create a GCP kubernetes
cluster to deploy a crud web app with its backend service

## Testing

To run the tests of this repository locally 

1. Place this repository under your GOPATH
2. Move into tests directory
3. Install the dependencies with dep ensure
4. Run ```go test -v ./...```

## Usage

To see how this module should be applied and its variables, please see the [Module usage document](gcp/README.md)