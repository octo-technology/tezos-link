variable "REGION" {
  type        = string
  description = "The region where the module is deployed"
  default     = "eu-west-1"
}

variable "ENV" {
  type        = string
  description = "The environment where the lambda will be executed."
  default     = "dev"
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

variable "S3_CODE_PATH" {
  type        = string
  description = "The path where the code is stored."
  default     = "v1.0.0/snapshot.zip"
}

variable "LAMBDA_PURPOSE" {
  type        = string
  description = "The purpose of the lambda (ex: snapshot)."
}

variable "LAMBDA_DESCRIPTION" {
  type        = string
  description = "The description of the lambda."
}

variable "LAMBDA_ENVIRONMENT_VARIABLES" {
  type        = map(string)
  description = "The environment variables to give to the lambda."
  default     = {}
}

variable "RUN_EVERY" {
  type        = string
  description = "The cron expression which will will trigger the lambda."
  default     = "rate(12 hours)"
}