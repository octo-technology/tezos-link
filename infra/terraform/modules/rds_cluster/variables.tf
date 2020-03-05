variable "REGION" {
  type    = string
  default = "europe-west-1"
}

variable "ENV" {
  type    = string
  default = "dev"
}

variable "VPC_CIDR" {
  type    = string
  default = "10.1.0.0/16"
}