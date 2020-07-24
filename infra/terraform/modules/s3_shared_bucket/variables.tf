variable "region" {
  type        = string
  description = "The region where the module is deployed"
  default     = "eu-west-1"
}

variable "project_name" {
  type        = string
  description = "The name of the project"
}

variable "s3_bucket_name" {
  type        = string
  description = "The S3 bucket name"
}

variable "vpc_endpoint_enabled" {
  type        = bool
  description = "Enable the endpoint associated with the VPC"
  default     = false
}

variable "vpc_cidr" {
  type        = string
  description = "The CIDR of the VPC where the S3 bucket endpoint will be placed. (required when vpc_endpoint_enabled = true)"
}