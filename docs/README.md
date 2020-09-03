# Technical documentations

## Table of Contents

1 - [Project organization](./Project_organization.md)

2 - [Architecture](./Architecture.md)

3 - [Services](./services)
  - [API](./services/API.md)
  - [Proxy](./services/Proxy.md)
  - [Frontend](./services/Frontend.md)

4 - [Run the project on the machine](./How_to_run_locally.md)

5 - Deploy the project on:
  - [an AWS tenant](./How_to_deploy_on_AWS.md)

6 - [Lambdas](./lambdas)
  - [Snapshot exporter](./lambdas/Snapshot_Exporter.md)
  - [Metrics cleaner](./lambdas/Metrics_Cleaner.md)

A - [References](./References.md)


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