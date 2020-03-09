resource "aws_security_group" "backend_lb" {
  name        = "backend_loadbalancer"
  description = "Security group for loadbalancer"
  vpc_id      = data.aws_vpc.tzlink.id
}

resource "aws_security_group_rule" "http_ingress_for_backend_loadbalancer" {
  type        = "ingress"
  from_port   = var.BACKEND_PORT
  to_port     = var.BACKEND_PORT
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.backend_lb.id
}

resource "aws_security_group_rule" "all_egress_for_backend_loadbalancer" {
  type        = "egress"
  from_port   = 0
  to_port     = 0
  protocol    = "-1"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.backend_lb.id
}

resource "aws_security_group" "backend_ecs_task" {
  name        = "backend_ecs_task"
  description = "Security group for backend (access only by loadbalancer)"
  vpc_id      = data.aws_vpc.tzlink.id
}

resource "aws_security_group_rule" "main_ingress_for_backend" {

  type                     = "ingress"
  from_port                = var.BACKEND_PORT
  to_port                  = var.BACKEND_PORT
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.backend_lb.id

  security_group_id = aws_security_group.backend_ecs_task.id
}

resource "aws_security_group_rule" "all_egress_for_backend" {
  type        = "egress"
  from_port   = 0
  to_port     = 0
  protocol    = "-1"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.backend_ecs_task.id
}