variable "REGION" {
  type    = string
  default = "europe-west1"
}

variable "ENV" {
  type    = "string"
  default = "dev"
}

variable "PROJECT_NAME" {
  type = "string"
  default = "tezos-link"
}

variable "BUILD_WITH" {
  type = "string"
  default = "terraform"
}

variable "VPC_CIDR" {
  type    = "string"
  default = "10.1.0.0/16"
}

variable "SUBNET_TZ_FARM_CIDR" {
  type    = "string"
  default = "10.1.0.0/24"
}