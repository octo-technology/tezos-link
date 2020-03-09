resource "aws_s3_bucket" "tz_front" {
  bucket = "tezoslink-front"
  acl    = "public-read"

  tags = {}
}

resource "aws_cloudfront_distribution" "tz_front" {
  enabled             = true
  is_ipv6_enabled     = true

  origin {
    domain_name = aws_s3_bucket.tz_front.bucket_regional_domain_name
    origin_id   = "TezosLinkFrontS3"
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  default_cache_behavior {
    target_origin_id = "TezosLinkFrontS3"

    allowed_methods = ["GET", "HEAD"]
    cached_methods  = ["GET", "HEAD"]

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "redirect-to-https"
    min_ttl                = 0
    default_ttl            = 7200
    max_ttl                = 86400
  }

  aliases = ["tezoslink.io"]

  viewer_certificate {
    cloudfront_default_certificate = true
  }

  tags = {}
}