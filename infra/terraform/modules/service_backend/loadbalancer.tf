resource "aws_alb" "backend" {
  name            = format("tzlink-%s-backend", var.ENV)
  subnets         = tolist(data.aws_subnet_ids.tzlink_public_backend.ids)
  security_groups = [aws_security_group.backend_lb.id]

  tags = {
    Name        = format("tzlink-%s-backend", var.ENV)
    Project     = var.PROJECT_NAME
    Environment = var.ENV
    BuildWith   = var.BUILD_WITH
  }
}

resource "aws_alb_target_group" "backend" {
  name        = format("tzlink-%s-backend", var.ENV)
  port        = var.BACKEND_PORT
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
    port     = var.BACKEND_PORT
    protocol = "HTTP"
  }

  tags = {
    Name        = format("tzlink-%s-backend", var.ENV)
    Project     = var.PROJECT_NAME
    Environment = var.ENV
    BuildWith   = var.BUILD_WITH
  }

  depends_on = [aws_alb.backend]
}

resource "aws_alb_listener" "backend" {
  load_balancer_arn = aws_alb.backend.id
  port              = var.BACKEND_PORT
  protocol          = "HTTP"

  default_action {
    target_group_arn = aws_alb_target_group.backend.id
    type             = "forward"
  }

  depends_on = [aws_alb_target_group.backend]
}