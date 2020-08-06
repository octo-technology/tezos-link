variable "region" {
  type        = string
  default     = "eu-west-1"
  description = "The region where the module will be deployed."
}

variable "project_name" {
  type        = string
  description = "The name of the project"
}

variable "certificate_arn" {
  type        = string
  description = "The certificate associated to cloudwatch SSL"
}