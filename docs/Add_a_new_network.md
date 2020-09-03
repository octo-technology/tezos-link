# Add a new network

Currently, tezos-link is able to handle two networks :
- `Mainnet`
- `Carthagenet` (testnet)

With the evolution of the Tezos blockchain, new testnets will
be created and we want to be able to deploy them easily.

We organize this project to made this new deployment quick and 
easy. To do this, please follow those steps:

## 1 - Push the associated snapshot in the S3 bucket

### Deployment process

When a tezos-node new machine start, the first thing that it do is to retrieve an archive with the snapshot of the node. This archive is the mecanism which permits to avoid to wait to many time before starting a new node.

This archive needs to be named
- `<network name>_node_data.tar.gz` for an archive node.
- `<network name>_rolling-snapshot.tar.gz` for a rolling node.

> Due to security issues, the command can be done only by someone who can access to the `AWS` tenant.
>
> Please ask to a **maintainer** to push the archive in the S3 bucket.

## 2 - Tezos nodes creation

### Deployment process

The first things that we want to create is the node which will
receive request. To do this, you need to go in the `terragrunt` folder and make a copy of folders:
- `01_tezos_carthagenet_archive_node` if you need archive nodes.
- `01_tezos_carthagenet_rolling_node` if you need rolling nodes.

The `01` represents the order where stack will be deployed so name them `01_tezos_<network name>_<node type>_node`.

Inside this folder, you will find a file named `terragrunt.hcl`.
It is this file which permits to deploy tezos nodes based on the module `terraform/modules/tezos_node`.

Please check the section [Node module: Inputs](#node-module-inputs) to get more information about inputs specific fields.

> Due to security issues, the deployment command can be done only by someone who can access to the `AWS` tenant.
>
> Please ask to a **maintainer** to deploy it on the main TezosLink tenant.

Once this file is done, you will need to go inside the folder created and make the command:

```bash
terragrunt plan
```

This will print resources that will be created.
If the stack does not return an error, we can run the command:

```bash
terragrunt apply
```

Now, our nodes are created with a dedicated **internal** dns :
`http://<network name>-<node type>.internal.tezoslink.io`

### Node module: Inputs

#### Node specifications:

- **tz_network**: The name of the network to deploy.
- **tz_mode**: The mode of the node wanted (currently supported: `rolling` and `archive`).
- **min_instance_number**: The minimal number of node wanted. 
- **desired_instance_number**: The number of node wanted when the stack is executed. This will be adapted by the scaling process.
- **max_instance_number**: The maximal number of node wanted that can be reach by the deployment process.
- **instance_type**: The `AWS` node size associated to the instance.
- **key_pair_name**: The SSH key that will be associated to the instance.

#### Healthcheck system specifications:

- **health_check_grace_period**: The period where the healthcheck system will wait the node setup. If the node is not started in this period, it will be killed and another will start instead.

#### Autoscaling specifications:

- **cpu_out_scaling_cooldown**: The time between two nodes scale-up process.
- **cpu_out_scaling_threshold**: The CPU threshold to reach to trigger the autoscaling process. Work with the `cpu_out_evaluation_periods`.
- **cpu_out_evaluation_periods**: The number of consecutive minutes where the `cpu_out_scaling_threshold` is break to scale-up.
- **cpu_down_scaling_cooldown**: The time between two nodes scale-down process.
- **cpu_down_scaling_threshold**: The CPU threshold to be below to trigger the autoscaling process. Work with the `cpu_down_evaluation_periods`.
- **cpu_down_evaluation_periods**: The number of consecutive minutes where the `cpu_down_scaling_threshold` is break to scale-down.

> When the response time be too important, we engage a second scale-up process to permits to accelerate the loadbalancing.

- **responsetime_out_scaling_cooldown**: The time between two node scale-up based on response time.
- **responsetime_out_scaling_threshold**  = The response time to reach before triggering the autoscaling process (in second). Work with `responsetime_out_evaluation_periods`.
- **responsetime_out_evaluation_periods** = The number of consecutive minutes where the `responsetime_out_scaling_threshold` is break to scale-out.

> **Warning**: The scale-up process will not work until the snapshot process is setup.

## 3 - Deployment of the snapshot lambda

### Deployment process

Once Tezos nodes are deployed, we need to setup the snapshot process to avoid new nodes to retrieve all the chain history from the beginning. To do this, we use an AWS dedicated lambda which will connect in SSH in the machine to run a systemd process.

To deploy a dedicated for a network, please make a copy of the folder: `02_lambda_snapshot_carthagenet` and rename it `02_lambda_snapshot_<network name>` to deploy the lambda.

Inside this folder, you will find a file named `terragrunt.hcl`. Again, you will need to modify few parameters to deploy a new lambda named `snapshot-<network name>`.

You will need to made the first deployment of the lambda with terragrunt. To do this, you will need to
- Go inside the folder `02_lambda_snapshot_<network name>`.
- Run the command `terragrunt plan` to check if the stack works.
- Run the command `terragrunt apply` to deploy the code in the AWS tenant.

> Again, due to security issues, plan and apply command can be done only by someone who can access to the `AWS` tenant.
>
> Please ask to a **maintainer** to deploy it on the main TezosLink tenant.

### Update the lambda when a new code is pushed

When the lambda is updated, all snapshot lambda need to be updated to use the new version of the code.
This process is automatized in a Github process workflow (`.github/workflows/cicd_lambda_snapshot.yml`). You can add your new created lambda by adding the block:

```yaml
    - name: Update lambda for <network name>
      run: |
        aws lambda update-function-code --function-name snapshot-<network name> --s3-bucket tzlink-snapshot-lambda --s3-key v1.0.0/snapshot.zip
```

### Proxy module: Inputs

- **name**        = "snapshot-mainnet"
- **description** = "Snapshot exporter lambda for mainnet"
- **environment_variables**
    - **NODE_USER**: "ubuntu"
    - **S3_REGION**: "eu-west-1"
    - **S3_BUCKET**: "tzlink-snapshot-lambda"
    - **S3_LAMBDA_KEY**: "snapshot_lambda_key"
    - **NETWORK**: "mainnet"

- **run_every**: "cron(0 1/12 * * ? *)"

- **bucket_name**: "tzlink-snapshot-lambda"
- **code_path**: "v1.0.0/snapshot.zip"

## 4 - Deployment of the proxy

### Deployment process

In this step, we want to create a dedicated proxy for our network and connect it with:
- The `Tezos nodes` created in the previous steps.
- The `Postgresql database` (to get the authentification and push the associated metrics for the frontend).

To do it, you will need to make a copy of `02_service_proxy_carthagenet` and rename it `02_service_proxy_<network name>`.On this part, we can find the `terragrunt.hcl` with the parameter to deploy the stack in `AWS`.

Please check the section [Proxy module: Inputs](#proxy-module-inputs) to get more information about inputs specific fields.

> Before running the script, please ask your maintainers to generate a certificate and the associated record with AWS Certificate Manager in the AWS tenant.

To deploy the proxy in the AWS tenant, you will need to add the following filled blocks at the end of the file `.github/workflows/cicd_proxy.yml`:

```yaml
    - name: Deploy proxy <network name> ecs-tasks
      working-directory: infra/terragrunt/02_service_proxy_<network name>
      run: |
        TF_VAR_DOCKER_IMAGE_VERSION="proxy-${GITHUB_SHA::8}" terragrunt apply -auto-approve
      env:
        TF_VAR_DATABASE_PASSWORD: ${{ secrets.DATABASE_PASSWORD }}
        REGISTRY: ${{ secrets.DOCKER_REGISTRY }}
```

With it, once merged on `master`, the proxy will be deployed and can be accessed with the url `https://<network_name>.tezoslink.io`. You can access it directly with internet.

### Proxy module: Inputs

- **tz_network**: The name of the network to deploy.
- **database_master_username**: The username of the database. Please don't change this value if you want to access to the main tezos database.
- **database_name**: The name of the database to connect. Please don't change this value if you want to access to the main tezos database.

## 5 - Adapt the Backend API and the frontend code to handle the new network.

< TODO: LTH, FRAH, BETA -> Please, fill this part >