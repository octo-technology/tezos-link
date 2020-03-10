resource "aws_alb" "proxy" {
  name            = "tzlink-proxy"
  subnets         = tolist(data.aws_subnet_ids.tzlink_public_proxy.ids)
  security_groups = [aws_security_group.proxy_lb.id]

  tags = {
    Name        = "tzlink-proxy"
    Project     = var.PROJECT_NAME
    BuildWith   = var.BUILD_WITH
  }
}

resource "aws_alb_target_group" "proxy" {
  name        = "tzlink-proxy"
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
    Name        = "tzlink-proxy"
    Project     = var.PROJECT_NAME
    BuildWith   = var.BUILD_WITH
  }

  depends_on = [aws_alb.proxy]
}

resource "aws_alb_listener" "proxy" {
  load_balancer_arn = aws_alb.proxy.id
  port              = 80
  protocol          = "HTTP"

  default_action {
    target_group_arn = aws_alb_target_group.proxy.id
    type             = "forward"
  }

  depends_on = [aws_alb_target_group.proxy]
}