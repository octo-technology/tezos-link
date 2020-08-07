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