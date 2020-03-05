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

variable "SUBNET_TZ_PUBLIC_PROXY_CIDR" {
  type    = string
  default = "10.1.1.0/24"
}

variable "SUBNET_TZ_PRIVATE_PROXY_CIDR" {
  type    = string
  default = "10.1.2.0/24"
}