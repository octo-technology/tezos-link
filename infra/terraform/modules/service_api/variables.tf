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
  description = "The CIDR of the VPC where the ECS service will be deployed"
}

variable "docker_image_name" {
  type        = string
  description = "The name of the deployed image"
}

variable "docker_image_version" {
  type        = string
  description = "The version of the deployed image"
}

variable "desired_container_number" {
  type        = number
  description = "The desired of container deployed by the service"
}

variable "port" {
  type        = number
  description = "The port open by the container that will be targeted by the loadbalancer"
}

variable "cpu" {
  type        = number
  description = "The CPU used by the service to run service's containers. (AWS CPU unit: 1vCPU = 1024)"
}

variable "memory" {
  type        = number
  description = "The RAM used by the service to run service's containers"
}

variable "configuration_file" {
  type        = string
  description = "the name of the configuration file inside the container"
}

variable "database_master_username" {
  type        = string
  description = "The username used by the service to connect on the RDS database."
}

variable "database_name" {
  type        = string
  description = "The name of the database used by the service"
}