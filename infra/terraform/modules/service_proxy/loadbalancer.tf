resource "aws_alb" "proxy" {
  name            = format("tzlink-%s-proxy", var.ENV)
  subnets         = tolist(data.aws_subnet_ids.tzlink_public_proxy.ids)
  security_groups = [aws_security_group.proxy_lb.id]

  tags = {
    Name        = format("tzlink-%s-proxy", var.ENV)
    Project     = var.PROJECT_NAME
    Environment = var.ENV
    BuildWith   = var.BUILD_WITH
  }
}

resource "aws_alb_target_group" "proxy" {
  name        = format("tzlink-%s-proxy", var.ENV)
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
    path     = "/v1/ad9df43d-a00a-4ded-b13b-dec1816402ad/chains/main/blocks/head"
    port     = var.PROXY_PORT
    protocol = "HTTP"
  }

  tags = {
    Name        = format("tzlink-%s-proxy", var.ENV)
    Project     = var.PROJECT_NAME
    Environment = var.ENV
    BuildWith   = var.BUILD_WITH
  }

  depends_on = [aws_alb.proxy]
}

resource "aws_alb_listener" "proxy" {
  load_balancer_arn = aws_alb.proxy.id
  port              = var.PROXY_PORT
  protocol          = "HTTP"

  default_action {
    target_group_arn = aws_alb_target_group.proxy.id
    type             = "forward"
  }

  depends_on = [aws_alb_target_group.proxy]
}