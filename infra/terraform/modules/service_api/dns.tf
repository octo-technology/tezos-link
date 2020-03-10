data "aws_route53_zone" "tezoslink" {
  name = "tezoslink.io."
}

resource "aws_route53_record" "api" {
  zone_id = data.aws_route53_zone.tezoslink.zone_id
  name    = "api.tezoslink.io"
  type    = "A"

  alias {
    name                   = aws_alb.api.dns_name
    zone_id                = aws_alb.api.zone_id
    evaluate_target_health = false
  }
}