terraform {
  required_version = "0.12.20"
  backend "s3" {}
}

# S3 Bucket

resource "aws_s3_bucket" "shared" {
  bucket = var.s3_bucket_name
  acl    = "private"

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm = "AES256"
      }
    }
  }

  versioning {
    enabled = true
  }

  tags = {
    Name    = var.s3_bucket_name
    Project = var.project_name
  }
}

# VPC Endpoint associated with the S3 Bucket
# (optional : will be created if vpc_endpoint_enabled = true)

data "aws_vpc" "tzlink" {
  count = var.vpc_endpoint_enabled ? 1 : 0

  cidr_block = var.vpc_cidr

  tags = {
    Name    = "tzlink"
    Project = var.project_name
  }
}

resource "aws_vpc_endpoint" "s3" {
  count = var.vpc_endpoint_enabled ? 1 : 0

  vpc_id       = data.aws_vpc.tzlink.0.id
  service_name = format("com.amazonaws.%s.s3", var.region)

  tags = {
    Name    = var.s3_bucket_name
    Project = var.project_name
  }
}