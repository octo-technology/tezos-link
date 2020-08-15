variable "region" {
  type        = string
  default     = "eu-west-1"
  description = "The region where the module will be deployed."
}

variable "tz_network" {
  type        = string
  description = "The tezos network that will be associated to the node. [mainnet, carthagenet]"
}

variable "tz_mode" {
  type        = string
  description = "The load mode of the tezos-node. [archive, rolling]"
}

variable "project_name" {
  type        = string
  description = "The name of the project"
}

variable "vpc_cidr" {
  type        = string
  description = "The CIDR of the VPC that will be deployed"
}

variable "instance_type" {
  type        = string
  description = "The size of the AWS instance that will be deployed"
}

variable "min_instance_number" {
  type        = number
  description = "The minimal number of instance inside the autoscaling_group"
}

variable "desired_instance_number" {
  type        = number
  description = "The desired number of instance wanted by default in the autoscaling_group"
}

variable "max_instance_number" {
  type        = number
  description = "The maximal number of instance inside the autoscaling_group"
}

variable "lambda_public_key" {
  type        = string
  description = "The ssh public key to permits to the lambda to connect on the instance"
}

variable "key_pair_name" {
  type        = string
  description = "The default key-pair placed on aws instance. Needed to connect with SSH on instances"
}

variable "cpu_out_scaling_cooldown" {
  type        = number
  description = "The cooldown (in second) between two trigger of the cpu_out autoscaling system"
}

variable "cpu_out_scaling_threshold" {
  type        = number
  description = "The percentage of CPU utilization which will trigger the cpu_out cloudwatch alarm"
}

variable "cpu_out_evaluation_periods" {
  type        = number
  description = "The number of consecutive minutes before triggering the cpu_out cloudwatch alarm"
}

variable "cpu_down_scaling_cooldown" {
  type        = number
  description = "The cooldown (in second) between two trigger of the cpu_down autoscaling system"
}

variable "cpu_down_scaling_threshold" {
  type        = number
  description = "The percentage of CPU utilization which will trigger the cpu_down cloudwatch alarm"
}

variable "cpu_down_evaluation_periods" {
  type        = number
  description = "The number of consecutive minutes before triggering the cpu_down cloudwatch alarm"
}

variable "responsetime_out_scaling_cooldown" {
  type        = number
  description = "The cooldown (in second) between two trigger of the responsetime_out autoscaling system"
}

variable "responsetime_out_scaling_threshold" {
  type        = string
  description = "The percentage of CPU utilization which will trigger the responsetime_out cloudwatch alarm"
}

variable "responsetime_out_evaluation_periods" {
  type        = number
  description = "The number of consecutive minutes before triggering the responsetime_out cloudwatch alarm"
}

variable "health_check_grace_period" {
  type        = number
  description = "The time given to a machine to reach the HEALTHY state. If not, the machine will be destroy and another will be created instead"
}