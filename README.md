# Product API Service

## Description

A sample product API service that uses an embedded sqlite database implemented in Go and containerized.

## Prerequisites

- Docker
- K8s cli (For K8s deployment)

## Building

The project can be built as a docker container with the following command:

`docker build -t go-product-api:latest .`

## Running locally

The built container can be run locally as follows:

`docker run -p 8080:8080 go-product-api:latest`

## Deploying to k8s

First step is to build and push the container to a container registry that can be accessed from a K8s cluster.

Under the K8s directory there is a sample deployment and service files. The deployment definition needs the path to the container image. It can be installed using the following sample commands. Note that the example puts all the data in an **emptyDir** mount which means all the data is reset if the application is restarted.

- `kubectl create -f k8s/deployment.yaml -n <namespace>`
- `kubectl create -f k8s/service.yaml -n <namespace>`

## Open API Spec

The open api spec for the API is available [here](./product-api.yaml)