terraform {
  required_version = "0.12.20"
  backend "s3" {}
}

data "aws_route53_zone" "tezoslink" {
  name = "tezoslink.io."
}

resource "aws_route53_record" "front" {
  zone_id = data.aws_route53_zone.tezoslink.zone_id
  name    = "tezoslink.io"
  type    = "A"

  alias {
    name                   = aws_cloudfront_distribution.tz_front.domain_name
    zone_id                = aws_cloudfront_distribution.tz_front.hosted_zone_id
    evaluate_target_health = false
  }
}

resource "aws_s3_bucket" "tz_front" {
  bucket = "tz-front"

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm = "AES256"
      }
    }
  }

  tags = {
    Name    = "tz-front"
    Project = var.project_name
  }
}

resource "aws_cloudfront_distribution" "tz_front" {
  enabled         = true
  is_ipv6_enabled = true

  default_root_object = "index.html"

  origin {
    domain_name = aws_s3_bucket.tz_front.bucket_regional_domain_name
    origin_id   = "TZLinkFrontS3"

    s3_origin_config {
      origin_access_identity = "origin-access-identity/cloudfront/E1H8RH9S03SPXS"
    }
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  custom_error_response {
    error_caching_min_ttl = "0"
    error_code            = "400"
    response_code         = "200"
    response_page_path    = "/index.html"
  }

  custom_error_response {
    error_caching_min_ttl = "0"
    error_code            = "404"
    response_code         = "200"
    response_page_path    = "/index.html"
  }

  custom_error_response {
    error_caching_min_ttl = "0"
    error_code            = "403"
    response_code         = "200"
    response_page_path    = "/index.html"
  }

  default_cache_behavior {
    target_origin_id = "TZLinkFrontS3"

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
    default_ttl            = 0
    max_ttl                = 0
  }

  aliases = ["tezoslink.io"]

  viewer_certificate {
    acm_certificate_arn      = var.certificate_arn
    ssl_support_method       = "sni-only"
    minimum_protocol_version = "TLSv1.2_2018"
  }

  tags = {
    Name    = "tz-front"
    Project = var.project_name
  }
}