resource "aws_s3_bucket" "tfstate_storage" {
  bucket = "tzlink-tfstate"

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
    Name        = "tzlink-tfstate"
    Project     = "tezos-link"
    Environment = "all"
    BuildWith   = "terraform"
    Trigramme   = "adbo"
  }
}

resource "aws_dynamodb_table" "tfstate_lock" {
  name           = "tzlink-tfstate-lock"
  read_capacity  = 20
  write_capacity = 20
  hash_key       = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }

  tags = {
    Name        = "tzlink-tfstate-lock"
    Project     = "tezos-link"
    Environment = "all"
    BuildWith   = "terraform"
    Trigramme   = "adbo"
  }
}