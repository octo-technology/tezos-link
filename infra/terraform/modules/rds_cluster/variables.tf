variable "REGION" {
  type    = string
  default = "eu-west-1"
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

variable "DATABASE_URL" {
  type    = string
  default = "database.example.nop"
}

variable "DATABASE_USERNAME" {
  type = string
}

variable "DATABASE_PASSWORD" {
  type = string
}

variable "DATABASE_TABLE" {
  type    = string
  default = "tezoslink"
}