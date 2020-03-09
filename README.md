
# Tezos Link

[![Go Report Card](https://goreportcard.com/badge/github.com/octo-technology/tezos-link)](https://goreportcard.com/report/github.com/octo-technology/tezos-link) ![Build](https://github.com/octo-technology/tezos-link/workflows/Build/badge.svg?branch=master)

Tezos link is a gateway to access to the Tezos network aiming to improve developer experience when developing Tezos dApps.

# Table of Contents

* [Project organization](#project-organization)
* [Services](#services)
  * [API](#api)
  * [Proxy](#proxy)
* [Build all services](#build-all-services)
* [Tests all services](#tests-all-services)
* [Run services locally on the machine](#run-services-locally-on-the-machine)
* [Infrastructure](#infrastructure)
  * [Architecture](#architecture)
  * [Requirements](#requirements)
  * [How To Deploy](#how-to-deploy)
* [Documentation](#documentation)
* [References](#references)

# Project Organization

The repository is currently following this organization:

```
.
├── api         # api documentation
├── build       # packaging
├── cmd         # mains
├── config      # config parsers
├── data        # config and migrations
├── docs        # services documentation
├── infra       # infrastructure
├── internal    # services
├── test        # test-specific files
└── web         # frontend
```

# Services

## API

REST API to manage projects and get project's metrics.

### Dependencies

* `PostgreSQL` (setup with 9.6)

### Environment variables

- `DATABASE_URL` (i.e `postgres:5432`)
- `DATABASE_USERNAME` (i.e `user`)
- `DATABASE_PASSWORD` (i.e `pass`)
- `DATABASE_TABLE` (i.e `tezoslink`)
- `TEZOS_HOST` (i.e `localhost`)
- `TEZOS_PORT` (i.e `1090`)
- `SERVER_PORT` (i.e `8001`)

## Proxy

- HTTP proxy in front of the nodes
- In-memory (LRU) cache

### Dependencies

* `PostgreSQL` (setup with 9.6)

### Environment variables

- `DATABASE_URL` (i.e `postgres:5432`)
- `DATABASE_USERNAME` (i.e `user`)
- `DATABASE_PASSWORD` (i.e `pass`)
- `DATABASE_TABLE` (i.e `tezoslink`)
- `SERVER_HOST` (i.e `localhost`)
- `SERVER_PORT` (i.e `8000`)

## Build all services

### Requirements

* `GNU Make` (setup with 3.81)
* `Golang` (setup with 1.13)
* `yarn`

### How to

To build your project, you need first to `install dependencies`:
```bash
$> make deps
```
After, you can run the `build` with
```bash
$> make build
```

## Test all services

### Requirements

* `Golang` (setup with 1.13)
* `GNU Make` (setup with 3.81)

For integrations tests only:
* `Docker`
* `docker-compose`
* `yarn`

### How to

To run the `unit tests`, you can use the command
```bash
$> make unit-test
```

To run `integration tests` locally, you will need to run following commands :
```bash
# We build docker images and run them
$> make build-docker & make run

# We run integration tests...
$> make integration-test

# And we clean the environment when we are done
$> docker-compose down
``` 

## Run services locally on the machine

### Requirements

* `Docker`
* `docker-compose`
* `yarn`

### How to

To run services locally on the machine, you will need to run those commands :
```bash
$> make build-docker

$> make run
```

## Infrastructure

### Architecture

TODO : To be redacted

### Requirements

* `Terraform` (version == 0.12.20)
* `Terragrunt` (version == 0.21.4)

> We recommend to install `tf-env` to manage easily your terraform environments.

### How to deploy

All the files related to the infrastructure are based on the `infra` folder.

First, you will need to update the configuration (if needed). To do this, you will find `common.tfvars` and `<env>.tfvars` in the folder `infra/terragrunt`.

When they are updated, we will use Terragrunt to deploy our infrastructure by running:
```bash
# To check if all is OK
$> terragrunt plan-all

# To apply the change
$> terragrunt apply-all
```

If you want to apply a specific part of the infrastructure (ex: `00_network`), you can run
```bash
$> cd infra/terragrunt/00_network

# To check if all is OK
$> terragrunt plan

# To apply the change
$> terragrunt apply
```

## Documentation

We use [Docsify](https://docsify.js.org/#/quickstart) to generate our documentation.

See docs [here](TODO).

## References

This repo took some ideas & code from:
- https://github.com/tezexInfo/TezProxy
- https://github.com/AymericBethencourt/serverless-mern-stack/
