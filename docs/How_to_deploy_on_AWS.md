# How to deploy

< TODO : To improve>

## Requirements

- `Terraform` (version == 0.12.20)
- `Terragrunt` (version == 0.21.4)

> We recommend to install `tf-env` to manage easily your terraform environments.

## Deployment process

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