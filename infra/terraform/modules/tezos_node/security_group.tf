resource "aws_security_group" "tezos_node" {
  name        = "tezos_node"
  description = "Tezos node"
  vpc_id      = data.aws_vpc.tzlink.id
}

resource "aws_security_group_rule" "rpc_ingress" {
  type            = "ingress"
  from_port       = 8000
  to_port         = 8000
  protocol        = "tcp"
  cidr_blocks     = ["0.0.0.0/0"]

  security_group_id = aws_security_group.tezos_node.id
}

resource "aws_security_group_rule" "ssh_ingress" {
  type            = "ingress"
  from_port       = 22
  to_port         = 22
  protocol        = "tcp"
  cidr_blocks     = ["0.0.0.0/0"]

  security_group_id = aws_security_group.tezos_node.id
}

resource "aws_security_group_rule" "all_egress" {
  type            = "egress"
  from_port       = 0
  to_port         = 0
  protocol        = "-1"
  cidr_blocks     = ["0.0.0.0/0"]
  # prefix_list_ids = ["pl-12c4e678"] # TODO: useful for VPC endpoint. To check later.

  security_group_id = aws_security_group.tezos_node.id
}