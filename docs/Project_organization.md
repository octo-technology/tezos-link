# Project organization

The repository is currently following this organization:

`.github`: This folder contains the code associated to Github CI/CD pipeline. Inside, you will find Github workflow files for:
  - The api service
  - The frontend service
  - The proxy service
  - Lambda functions (Snapshot exporter and Metrics cleaner)
  - The terraform/terragrunt part

`api-docs`: This folder contains the swagger code for the api.

`build`: This folder contains files for Docker image generation.

`cmd`: This folder contains the entrypoint of every applications:
  - proxy
  - api
  - lambda snapshot exporter
  - lambda metrics cleaner

`config`: This folder contains the configuration parser and the associated model for
  - the api
  - the proxy

`data`: In this folder, you will find:
  - the configuration file for every applications and for specific environment (local, production...)
  - database models used by the golang migration library.

`docs`: The folder with the technical documentation about how to use this repository.

`infra`: In this folder, you can find files associated with the infrastructure deployment:
  - `terraform` which contains terraform modules, AWS IAM policy generator...
  - `terragrunt` which call the terraform module to apply specifics variables and avoid code repetition.

`internal`: In this folder, you will find application codes using the Domain Driven Design.

`perf`: The folder where was placed loadtest codebase.

`pkg`: <TODO>

`test`: Specific configuration files used to run tests.

`web`: The folder which contains frontend:
  - code
  - graphical assets
  - documentation about application usage