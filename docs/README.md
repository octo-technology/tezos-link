# Technical documentations

## Table of Contents

- [Project organization](#project-organization)
- [Run services locally on the machine](#run-services-locally-on-the-machine)
- [Build all services](#build-all-services)
- [Tests all services](#tests-all-services)
- [Frontend](#frontend)
- [Services](#services)
  - [API](#api)
  - [Proxy](#proxy)
  - [Snapshot exporter](#snapshot-exporter)
- [Infrastructure](#infrastructure)
  - [Architecture](#architecture)
  - [Requirements](#requirements)
  - [How To Deploy](#how-to-deploy)
- [Documentation](#documentation)
- [References](#references)

## Project Organization

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

## Run services locally on the machine with mockup blockchain node

> Blockchain nodes are mocked up for development environment the be as lightweight as possible. 

### Requirements

- `Docker`
- `docker-compose`
- `Yarn` (setup with 1.22.0)
- `Golang` (setup with 1.13)
- `GNU Make` (setup with 3.81)
- `Node.js` (setup with 11.14.0)

### How to

To run services locally on the machine, you will need to run those commands :

```bash
$> make deps
$> make build-docker
$> make run-dev
```

It will run:

- `tezos-link_proxy`
- `tezos-link_proxy-carthagenet`
- `tezos-link_api`
- `mockserver/mockserver:mockserver-5.9.0` (mocking a blockchain node)
- `postgres:9.6`

The only endpoint served by the blockchain mock is:

```bash
curl -X PUT localhost:8001/v1/<YOUR_PROJECT_ID>/mockserver/status
```

## Test all services

### Requirements

- `Golang` (setup with 1.13)
- `GNU Make` (setup with 3.81)

For integrations tests only:

- `Docker`
- `docker-compose`
- `yarn`

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

- `GNU Make` (setup with 3.81)
- `Golang` (setup with 1.13)
- `Yarn` (setup with 1.22.0)
- `Node.js` (setup with 11.14.0)

### How to

To build your project, you need first to `install dependencies`:

```bash
$> make deps
```

After, you can run the `build` with

```bash
$> make build
```

## Services

### Snapshot exporter lambda

Lambda function scheduled with a `Cloudwatch Rule` cronjob, connect to a node with SSH and trigger a snapshot export.

#### Deploy for testing and for development purpose

Individual deployment of the lambda is possible for testing and development purpose.

> You will need AWS credentials setup on your machine, see [AWS Credentials](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html)

It will build, send to the S3 bucket and update snapshot lambda code:

```bash
$> make build-unix
$> make deploy-snapshot-lambda
```

To execute the lambda, run:

```bash
aws lambda invoke --region=eu-west-1 --function-name=snapshot --log Tail output.txt | grep "LogResult"| awk -F'"' '{print $4}' | base64 --decode
```

### Metrics cleaner lambda

Lambda function scheduled with a `Cloudwatch Rule` cronjob, connect to a node with SSH and trigger a metrics clean.

#### Deploy for testing and development purpose

Individual deployment of the lambda is possible for testing and development purpose.

> You will need AWS credentials setup on your machine, see [AWS Credentials](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html)

It will build, send to the S3 bucket and update metrics-cleaner lambda code:

```bash
$> make build-unix
$> make deploy-metrics-cleaner-lambda
```

To execute the lambda, run:

```bash
aws lambda invoke --region=eu-west-1 --function-name=metrics --log Tail output.txt | grep "LogResult"| awk -F'"' '{print $4}' | base64 --decode
```

#### Environment variables

These environment variables are set in `infra/dev.tfvars`.

- `NODE_USER` (default: `ec2-user`)
- `S3_REGION` (default: `eu-west-1`)
- `S3_BUCKET` (default: `tzlink-snapshot-lambda-dev`)
- `S3_LAMBDA_KEY` (default: `snapshot_lambda_key`)

## Infrastructure

### Requirements

- `Terraform` (version == 0.12.20)
- `Terragrunt` (version == 0.21.4)

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