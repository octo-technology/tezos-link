data "aws_route53_zone" "tezoslink_private" {
  name         = "internal.tezoslink.io."
  private_zone = true
}

resource "aws_route53_record" "internal_lb_farm" {
  zone_id = data.aws_route53_zone.tezoslink_private.zone_id
  name    = format("%s-%s.internal.tezoslink.io", var.TZ_NETWORK, var.TZ_MODE)
  type    = "A"

  alias {
    name                   = aws_alb.tz_farm.dns_name
    zone_id                = aws_alb.tz_farm.zone_id
    evaluate_target_health = false
  }
}