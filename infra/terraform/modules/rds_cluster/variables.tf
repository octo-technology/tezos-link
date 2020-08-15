variable "region" {
  type        = string
  default     = "eu-west-1"
  description = "The region where the module will be deployed."
}

variable "project_name" {
  type        = string
  description = "The name of the project."
}

variable "vpc_cidr" {
  type        = string
  description = "The CIDR of the VPC where the Aurora will be placed."
}

variable "database_master_username" {
  type        = string
  description = "The username used by the aurora database as the master username."
}

variable "database_name" {
  type        = string
  description = "The name of the database created during the creation of the aurora."
}