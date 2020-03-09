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

variable "PROXY_DOCKER_IMAGE_NAME" {
  type    = string
  default = "louptheronlth/tezos-link"
}

variable "PROXY_DOCKER_IMAGE_VERSION" {
  type    = string
  default = "proxy-dev"
}

variable "PROXY_PORT" {
  type    = number
  default = 8001
}

variable "PROXY_CPU" {
  type    = number
  default = 1024 # 1 vCPU
}

variable "PROXY_MEMORY" {
  type    = number
  default = 250
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

variable "TEZOS_FARM_URL" {
  type    = string
  default = "farm.example.nop"
}

variable "TEZOS_FARM_PORT" {
  type    = string
  default = 80
}

variable "PROXY_CONFIGURATION_FILE" {
  type    = string
  default = "dev"
}

variable "PROXY_DESIRED_COUNT" {
  type    = number
  default = 0
}