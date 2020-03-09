resource "aws_security_group" "database" {
  name        = "database"
  description = "Security group for database"
  vpc_id      = data.aws_vpc.tzlink.id
}

resource "aws_security_group_rule" "all_ingress_for_database" {
  type        = "ingress"
  from_port   = 0
  to_port     = 0
  protocol    = "-1"
  cidr_blocks = [var.VPC_CIDR]

  security_group_id = aws_security_group.database.id
}

resource "aws_security_group_rule" "all_egress_for_database" {
  type        = "egress"
  from_port   = 0
  to_port     = 0
  protocol    = "-1"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.database.id
}