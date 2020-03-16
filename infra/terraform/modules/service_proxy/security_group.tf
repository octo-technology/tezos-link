resource "aws_security_group" "proxy_lb" {
  name        = format("proxy_%s_loadbalancer", var.TZ_NETWORK)
  description = format("Security group for loadbalancer (network %s)", var.TZ_NETWORK)
  vpc_id      = data.aws_vpc.tzlink.id
}

resource "aws_security_group_rule" "http_ingress_for_loadbalancer" {
  type        = "ingress"
  from_port   = 80
  to_port     = 80
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.proxy_lb.id
}

resource "aws_security_group_rule" "https_ingress_for_loadbalancer" {
  type        = "ingress"
  from_port   = 443
  to_port     = 443
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.proxy_lb.id
}

resource "aws_security_group_rule" "all_egress_for_loadbalancer" {
  type        = "egress"
  from_port   = 0
  to_port     = 0
  protocol    = "-1"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.proxy_lb.id
}

resource "aws_security_group" "proxy_ecs_task" {
  name        = format("proxy_%s_ecs_task", var.TZ_NETWORK)
  description = format("Security group for proxy %s (access only by loadbalancer)", var.TZ_NETWORK)
  vpc_id      = data.aws_vpc.tzlink.id
}

resource "aws_security_group_rule" "main_ingress_for_proxy" {

  type                     = "ingress"
  from_port                = var.PROXY_PORT
  to_port                  = var.PROXY_PORT
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.proxy_lb.id

  security_group_id = aws_security_group.proxy_ecs_task.id
}

resource "aws_security_group_rule" "all_egress_for_proxy" {
  type        = "egress"
  from_port   = 0
  to_port     = 0
  protocol    = "-1"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.proxy_ecs_task.id
}