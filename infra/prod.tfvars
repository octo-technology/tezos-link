REGION = "eu-west-1"

ENV = "prod"

VPC_CIDR                        = "10.1.0.0/16"
vpc_cidr                        = "10.1.0.0/16"

SUBNET_TZ_FARM_CIDR             = "10.1.0.0/24"
SUBNET_TZ_PUBLIC_PROXY_CIDR     = "10.1.1.0/24"
SUBNET_TZ_PRIVATE_PROXY_CIDR    = "10.1.2.0/24"
SUBNET_TZ_PRIVATE_DATABASE_CIDR = "10.1.3.0/24"
subnet_tz_farm_cidr             = "10.1.0.0/24"
subnet_tz_public_ecs_cidr       = "10.1.1.0/24"
subnet_tz_private_ecs_cidr      = "10.1.2.0/24"
subnet_tz_private_database_cidr = "10.1.3.0/24"

DATABASE_USERNAME = "administrator"
DATABASE_TABLE = "tezoslink"