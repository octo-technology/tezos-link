resource "aws_s3_bucket" "tz_front" {
  bucket = "tezoslink-front"
  acl    = "public-read"
  policy = <<EOF
{
  "Version":"2012-10-17",
  "Statement":[
    {
      "Sid":"PublicRead",
      "Effect":"Allow",
      "Principal": "*",
      "Action":["s3:GetObject"],
      "Resource":["arn:aws:s3:::tezoslink-front/*"]
    }
  ]
}
EOF

  tags = {
    Name      = "tezoslink-front"
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }
}

resource "aws_cloudfront_distribution" "tz_front" {
  enabled         = true
  is_ipv6_enabled = true

  default_root_object = "index.html"

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
    default_ttl            = 0
    max_ttl                = 0
  }

  aliases = ["tezoslink.io"]

  viewer_certificate {
    #cloudfront_default_certificate = true
    acm_certificate_arn      = "arn:aws:acm:us-east-1:609827314188:certificate/05558bc2-3703-49ab-ae81-e22ea2ea437e"
    ssl_support_method       = "sni-only"
    minimum_protocol_version = "TLSv1.2_2018"
  }

  tags = {}
}