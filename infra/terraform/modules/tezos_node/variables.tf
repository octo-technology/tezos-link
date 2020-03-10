variable "REGION" {
  type    = string
  default = "europe-west-1"
}

variable "ENV" {
  type    = string
  default = "dev"
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

variable "TZ_NODE_NUMBER" {
  type    = number
  default = 1
}

variable "INSTANCE_TYPE" {
  type    = string
  default = "i3.large"
}

variable "KEY_PAIR_NAME" {
  type    = string
  default = "AWS-TEZOS-KEY"
}