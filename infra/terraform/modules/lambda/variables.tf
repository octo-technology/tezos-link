variable "region" {
  type        = string
  default     = "eu-west-1"
  description = "The region where the module will be deployed."
}

variable "project_name" {
  type        = string
  description = "The name of the project"
}

variable "bucket_name" {
  type        = string
  description = "The S3 bucket name where the code used is placed"
}

variable "code_path" {
  type        = string
  description = "The path to the code used by the lambda function inside the S3 bucket"
}

variable "name" {
  type        = string
  description = "The name of the lambda function"
}

variable "description" {
  type        = string
  description = "The description of the lambda function"
}

variable "environment_variables" {
  type        = map(string)
  description = "The map of the environment variables used by the lambda"
  default     = {}
}

variable "run_every" {
  type        = string
  description = "The cron expression which schedule the lambda trigger alarm. (default: every 12 hours)"
  default     = "rate(12 hours)"
}

variable "vpc_config_enable" {
  type        = bool
  description = "Enable the placement of the lambda inside the VPC"
}

variable "vpc_cidr" {
  type        = string
  description = "[OPTIONAL] The VPC CIDR where the lambda will be deployed. (needed when vpc_config_enable=true)"
  default     = ""
}

variable "subnet_name" {
  type        = string
  description = "[OPTIONAL] The subnet where the lambda will be deployed. (needed when vpc_config_enable=true)"
  default     = ""
}

variable "security_group_name" {
  type        = string
  description = "[OPTIONAL] The security_group used by the lambda. (needed when vpc_config_enable=true)"
  default     = ""
}