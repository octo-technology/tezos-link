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