resource "aws_security_group" "tezos_node" {
  name        = "tezos_node-${var.ENV}"
  description = "Security group for tezos-${var.ENV} nodes"
  vpc_id      = data.aws_vpc.tzlink.id
}

resource "aws_security_group_rule" "rpc_ingress_for_tezos_node" {
  type            = "ingress"
  from_port       = 8000
  to_port         = 8000
  protocol        = "tcp"
  cidr_blocks     = ["0.0.0.0/0"]

  security_group_id = aws_security_group.tezos_node.id
}

resource "aws_security_group_rule" "ssh_ingress_for_tezos_node" {
  type            = "ingress"
  from_port       = 22
  to_port         = 22
  protocol        = "tcp"
  cidr_blocks     = ["0.0.0.0/0"]

  security_group_id = aws_security_group.tezos_node.id
}

resource "aws_security_group_rule" "all_egress_for_tezos_node" {
  type            = "egress"
  from_port       = 0
  to_port         = 0
  protocol        = "-1"
  cidr_blocks     = ["0.0.0.0/0"]
  # prefix_list_ids = ["pl-12c4e678"] # TODO: useful for VPC endpoint. To check later.

  security_group_id = aws_security_group.tezos_node.id
}

resource "aws_security_group" "tezos_node_lb" {
  name        = "tezos_node_lb-${var.ENV}"
  description = "Security group for tezos-${var.ENV} loadbalancer"
  vpc_id      = data.aws_vpc.tzlink.id
}

resource "aws_security_group_rule" "http_ingress_for_loadbalancer" {
  type            = "ingress"
  from_port       = 80
  to_port         = 80
  protocol        = "tcp"
  cidr_blocks     = [ var.VPC_CIDR ]

  security_group_id = aws_security_group.tezos_node_lb.id
}


resource "aws_security_group_rule" "all_egress_for_loadbalancer" {
  type            = "egress"
  from_port       = 0
  to_port         = 0
  protocol        = "-1"
  cidr_blocks     = ["0.0.0.0/0"]

  security_group_id = aws_security_group.tezos_node_lb.id
}