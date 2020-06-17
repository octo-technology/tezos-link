variable "REGION" {
  type    = string
  default = "europe-west-1"
}

variable "TZ_NETWORK" {
  type    = string
  default = "mainnet"

  description = "The current network to deploy in the tezos_node. (possible choice: mainnet, carthagenet)"
}

variable "TZ_MODE" {
  type    = string
  default = "archive"

  description = "The current mode of the node wanted (archive, rolling)"
}

variable "PROJECT_NAME" {
  type    = string
  default = "tezos-link"
}

variable "BUILD_WITH" {
  type    = string
  default = "terraform"
}

variable "VPC_CIDR" {
  type    = string
  default = "10.1.0.0/16"
}

variable "SUBNET_TZ_FARM_CIDR" {
  type    = string
  default = "10.1.0.0/24"
}

variable "INSTANCE_TYPE" {
  type    = string
  default = "i3.large"
}

variable "MIN_INSTANCE_NUMBER" {
  type    = number
  default = 1
}

variable "MAX_INSTANCE_NUMBER" {
  type    = number
  default = 3
}

variable "DESIRED_INSTANCE_NUMBER" {
  type    = number
  default = 1
}

variable "KEY_PAIR_NAME" {
  type    = string
  default = "AWS-TEZOS-KEY"
}
