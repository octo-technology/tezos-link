resource "aws_s3_bucket" "remote_state_storage" {
  bucket = "tzlink-remote-state"

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
    Name        = "tzlink-remote-state"
    Project     = "tezos-link"
    Environment = "global"
  }
}

resource "aws_dynamodb_table" "remote_state_lock" {
  name           = "tzlink-remote-state-lock"
  read_capacity  = 20
  write_capacity = 20
  hash_key       = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }

  tags = {
    Name        = "tzlink-remote-state-lock"
    Project     = "tezos-link"
    Environment = "global"
  }
}