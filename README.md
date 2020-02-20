# Tezos Link

[![Go Report Card](https://goreportcard.com/badge/github.com/octo-technology/tezos-link)](https://goreportcard.com/report/github.com/octo-technology/tezos-link)

Scalable API access to the Tezos network

## Install

Install `go`, then

```shell
$> make deps
```

## Build 

```shell
$> make build
```

## Test

```shell
$> make test
```

## Run locally

> ### Prerequisites
> 1 - Install:
> - `docker-compose`
> - `docker`

Then run on first time (might takes minutes to complete):
```shell
$> make build-docker
```

Then, run the containers with:
```shell
$> make run
```

### Requirements

`backend` and `proxy` services requires :
- PostgreSQL v9.6

## Documentation

We use [Docsify](https://docsify.js.org/#/quickstart) to generate our documentation.

See docs [here](TODO).

### References

This repo took some ideas & code from:
- https://github.com/tezexInfo/TezProxy
- https://github.com/AymericBethencourt/serverless-mern-stack/
