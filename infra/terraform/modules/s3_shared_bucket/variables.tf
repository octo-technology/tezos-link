variable "REGION" {
  type        = string
  description = "The region where the module is deployed"
  default     = "eu-west-1"
}

variable "PROJECT_NAME" {
  type        = string
  description = "The name of the project associated to the lambda."
  default     = "tezos-link"
}

variable "BUILD_WITH" {
  type        = string
  description = "Permits to know on the AWS tags that objects are build with IaC."
  default     = "terraform"
}

variable "S3_BUCKET_NAME" {
  type        = string
  description = "The S3 bucket name that will be associated to the lambda"
}