resource "aws_alb" "proxy" {
  name            = format("tzlink-proxy-%s", var.TZ_NETWORK)
  subnets         = tolist(data.aws_subnet_ids.tzlink_public_ecs.ids)
  security_groups = [aws_security_group.proxy_lb.id]

  tags = {
    Name      = format("tzlink-proxy-%s", var.TZ_NETWORK)
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }
}

resource "aws_alb_target_group" "proxy" {
  name        = format("tzlink-proxy-%s", var.TZ_NETWORK)
  port        = var.PROXY_PORT
  protocol    = "HTTP"
  vpc_id      = data.aws_vpc.tzlink.id
  target_type = "ip"

  stickiness {
    enabled         = true
    type            = "lb_cookie"
    cookie_duration = 600
  }

  health_check {
    enabled  = true
    path     = "/health"
    port     = var.PROXY_PORT
    protocol = "HTTP"
  }

  tags = {
    Name      = format("tzlink-proxy-%s", var.TZ_NETWORK)
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }

  depends_on = [aws_alb.proxy]
}

#resource "aws_alb_listener" "proxy" {
#  load_balancer_arn = aws_alb.proxy.id
#  port              = 80
#  protocol          = "HTTP"
#
#  default_action {
#    target_group_arn = aws_alb_target_group.proxy.id
#    type             = "forward"
#  }
#
#  depends_on = [aws_alb_target_group.proxy]
#}

############

resource "aws_alb_listener" "proxy_https" {
  load_balancer_arn = aws_alb.proxy.arn
  port              = 443
  protocol          = "HTTPS"
  ssl_policy        = "ELBSecurityPolicy-TLS-1-1-2017-01"
  certificate_arn   = local.network_certificate_arn

  default_action {
    target_group_arn = aws_alb_target_group.proxy.arn
    type             = "forward"
  }

  depends_on = [aws_alb_target_group.proxy]
}

resource "aws_alb_listener" "proxy_http_redirect" {
  load_balancer_arn = aws_alb.proxy.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    type = "redirect"

    redirect {
      port        = 443
      protocol    = "HTTPS"
      status_code = "HTTP_301"
    }
  }

  depends_on = [aws_alb_target_group.proxy]
}