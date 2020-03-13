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

variable "SNAPSHOT_S3_KEY" {
  type    = string
  default = "v1.0.0/snapshot.zip"
}

variable "NODE_IP" {
  type    = string
  default = "0.0.0.0"
}

variable "NODE_USER" {
  type    = string
  default = "ec2-user"
}

variable "S3_LAMBDA_KEY" {
  type    = string
  default = "snapshot_lambda_key"
}
