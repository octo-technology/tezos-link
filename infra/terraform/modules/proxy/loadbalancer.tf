resource "aws_alb" "proxy" {
  name            = "tf-ecs-chat"
  subnets         = tolist(data.aws_subnet_ids.tzlink_proxy.ids)
  security_groups = [aws_security_group.proxy_lb.id]
}

resource "aws_alb_target_group" "proxy" {
  name        = "tf-ecs-chat"
  port        = var.PROXY_PORT
  protocol    = "HTTP"
  vpc_id      = data.aws_vpc.tzlink.id
  target_type = "ip"
}

resource "aws_alb_listener" "proxy" {
  load_balancer_arn = aws_alb.proxy.id
  port              = var.PROXY_PORT
  protocol          = "HTTP"

  default_action {
    target_group_arn = aws_alb_target_group.proxy.id
    type             = "forward"
  }
}