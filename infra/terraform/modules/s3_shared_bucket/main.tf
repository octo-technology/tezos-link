resource "aws_s3_bucket" "shared" {
  bucket = var.S3_BUCKET_NAME
  acl    = "private"

  tags = {
    Name      = var.S3_BUCKET_NAME
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }
}