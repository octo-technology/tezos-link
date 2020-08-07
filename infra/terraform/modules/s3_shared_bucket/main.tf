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
