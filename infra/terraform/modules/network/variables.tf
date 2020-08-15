variable "region" {
  type        = string
  default     = "eu-west-1"
  description = "The region where the module will be deployed"
}

variable "project_name" {
  type        = string
  description = "The name of the project"
}

variable "vpc_cidr" {
  type        = string
  description = "The CIDR of the VPC that will be deployed"
}

variable "subnet_tz_farm_cidr" {
  type        = string
  description = "The CIDR of the subnet associated to the tezos-node farm"
}

variable "subnet_tz_public_ecs_cidr" {
  type        = string
  description = "The CIDR of the subnet associated to the ecs's loadbalancers"
}

variable "subnet_tz_private_ecs_cidr" {
  type        = string
  description = "The CIDR of the subnet associated to the ecs containers"
}

variable "subnet_tz_private_database_cidr" {
  type        = string
  description = "The CIDR of the subnet associated to the databases"
}