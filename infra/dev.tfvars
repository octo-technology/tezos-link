REGION = "eu-west-1"

ENV = "dev"

VPC_CIDR                        = "10.1.0.0/16"

SUBNET_TZ_FARM_CIDR             = "10.1.0.0/24"
SUBNET_TZ_PUBLIC_PROXY_CIDR     = "10.1.1.0/24"
SUBNET_TZ_PRIVATE_PROXY_CIDR    = "10.1.2.0/24"
SUBNET_TZ_PRIVATE_DATABASE_CIDR = "10.1.3.0/24"

INSTANCE_TYPE = "i3.large"
KEY_PAIR_NAME = "adbo"

PROXY_DOCKER_IMAGE_NAME = "louptheronlth/tezos-link"
PROXY_DOCKER_IMAGE_VERSION = "proxy-dev"

DATABASE_URL = "tzlink-dev-database.cmeu9dixowfa.eu-west-1.rds.amazonaws.com:5432"
DATABASE_USERNAME = "administrator"
DATABASE_TABLE = "tezoslink"

TEZOS_FARM_URL = "internal-tzlink-dev-farm-587823368.eu-west-1.elb.amazonaws.com"
TEZOS_FARM_PORT = 80

PROXY_CONFIGURATION_FILE = "prod"

PROXY_PORT = 8001
PROXY_CPU = 256
PROXY_MEMORY = 512

PROXY_DESIRED_COUNT = 1
