
# Tezos Link

[![Go Report Card](https://goreportcard.com/badge/github.com/octo-technology/tezos-link)](https://goreportcard.com/report/github.com/octo-technology/tezos-link) ![Build](https://github.com/octo-technology/tezos-link/workflows/Build/badge.svg?branch=master)

Tezos link is a gateway to access to the Tezos network aiming to improve developer experience when developing Tezos dApps.

# Table of Contents

* [Project organization](#project-organization)
* [Run services locally on the machine](#run-services-locally-on-the-machine)
* [Build all services](#build-all-services)
* [Tests all services](#tests-all-services)
* [Frontend](#frontend)
* [Services](#services)
  * [API](#api)
  * [Proxy](#proxy)
  * [Snapshot exporter](#snapshot-exporter)
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
├── api          # api documentation
├── build        # packaging
├── cmd          # mains
├── config       # config parsers
├── data         # config and migrations
├── infra        # infrastructure
├── internal     # services
├── test         # test-specific files
└── web          # frontend
    └── public
        └── docs # usage documentation
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

It will run:
- `tezos-link_proxy`
- `tezos-link_api`
- `mockserver/mockserver:mockserver-5.9.0`
- `postgres:9.6`

Mockserver is mocking the blockchain node, the only endpoint served by this mock is:
```bash
curl -X PUT localhost:8001/v1/<YOUR_PROJECT_ID/mockserver/status
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

## Build all services and frontend

### Requirements

* `GNU Make` (setup with 3.81)
* `Golang` (setup with 1.13)
* `yarn` (setup with 1.22.0)

### How to

To build your project, you need first to `install dependencies`:
```bash
$> make deps
```
After, you can run the `build` with
```bash
$> make build
```

## Frontend

### Run

To run the frontend, execute:
```bash
$> cd web && yarn start
```

### Build

To run the frontend, execute:
```bash
$> make build-frontend
```

### Deploy

> You will need AWS credentials setup on your machine, see [AWS Credentials](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html)

To run the frontend, execute:
```bash
$> make deploy-frontend
```

## Services

### API

REST API to manage projects and get project's metrics.

#### Dependencies

* `PostgreSQL` (setup with 9.6)

#### Environment variables

- `DATABASE_URL` (default: `postgres:5432`)
- `DATABASE_USERNAME` (default: `user`)
- `DATABASE_PASSWORD` (default: `pass`)
- `DATABASE_TABLE` (default: `tezoslink`)
- `DATABASE_ADDITIONAL_PARAMETER` (default: `sslmode=disable`)
- `SERVER_HOST` (default: `localhost`)
- `SERVER_PORT` (default: `8000`)

### Proxy

- HTTP proxy in front of the nodes
- In-memory (LRU) cache

#### Dependencies

* `PostgreSQL` (setup with 9.6)

#### Environment variables

- `DATABASE_URL` (default: `postgres:5432`)
- `DATABASE_USERNAME` (default: `user`)
- `DATABASE_PASSWORD` (default: `pass`)
- `DATABASE_TABLE` (default: `tezoslink`)
- `DATABASE_ADDITIONAL_PARAMETER` (default: `sslmode=disable`)
- `TEZOS_HOST` (default: `node`)
- `TEZOS_PORT` (default: `1090`)
- `SERVER_PORT` (default: `8001`)

### Snapshot exporter

Lambda function scheduled with a `Cloudwatch Rule` cronjob, connect to a node with SSH and trigger a snapshot export.

#### Deploy

> You will need AWS credentials setup on your machine, see [AWS Credentials](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html)

It will build, send to the S3 bucket and update lambda code:
```bash
$> make build-unix
$> make deploy-lambda
```

To execute the lambda, run:
```bash
aws lambda invoke --region=eu-west-1 --function-name=snapshot --log Tail output.txt | grep "LogResult"| awk -F'"' '{print $4}' | base64 --decode
```

#### Environment variables

These environment variables are set in `infra/dev.tfvars`.

- `NODE_USER` (default: `ec2-user`)
- `S3_REGION` (default: `eu-west-1`)
- `S3_BUCKET` (default: `tzlink-snapshot-lambda-dev`)
- `S3_LAMBDA_KEY` (default: `snapshot_lambda_key`)


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

> Currently, database password is encrypted inside the file `vaulted.tfvars`. To see it content, you will need ansible-vault and a passphrase to decrypt it with the command `ansible-vault decrypt vaulted.tfvars`.
>
> This will be changed soon with AWS Secret Manager.

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

The documentation usage is located at `web/public/docs`.

It contains the various Markdown files served by the application at `/documentation`.

## References

This repo took some ideas & code from:
- https://github.com/tezexInfo/TezProxy
- https://github.com/AymericBethencourt/serverless-mern-stack/
