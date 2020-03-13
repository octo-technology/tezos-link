resource "aws_alb" "api" {
  name            = "tzlink-api"
  subnets         = tolist(data.aws_subnet_ids.tzlink_public_ecs.ids)
  security_groups = [aws_security_group.api_lb.id]

  tags = {
    Name      = "tzlink-api"
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }
}

resource "aws_alb_target_group" "api" {
  name        = "tzlink-api"
  port        = var.API_PORT
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
    port     = var.API_PORT
    protocol = "HTTP"
  }

  tags = {
    Name      = "tzlink-api"
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }

  depends_on = [aws_alb.api]
}

resource "aws_alb_listener" "api_https" {
  load_balancer_arn = aws_alb.api.arn
  port              = 443
  protocol          = "HTTPS"
  ssl_policy        = "ELBSecurityPolicy-TLS-1-1-2017-01"
  certificate_arn   = local.api_certificate_arn

  default_action {
    target_group_arn = aws_alb_target_group.api.arn
    type             = "forward"
  }

  depends_on = [aws_alb_target_group.api]
}

resource "aws_alb_listener" "api_http_redirect" {
  load_balancer_arn = aws_alb.api.arn
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

  depends_on = [aws_alb_target_group.api]
}