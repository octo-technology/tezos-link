resource "aws_s3_bucket" "configuration_bucket" {
  bucket = "tezos-link-tfstate"

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
      Name = "tezos-link-tfstate"
      Environment = "all"
  }
}

resource "aws_dynamodb_table" "configuration_dynamodb_table" {
  name           = "tezos-link-tfstate-lock"
  read_capacity  = 20
  write_capacity = 20
  hash_key       = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }

  tags = {
    Name        = "tezos-link-tfstate-lock"
    Environment = "all"
  }
}