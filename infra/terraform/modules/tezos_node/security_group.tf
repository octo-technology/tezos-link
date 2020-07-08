resource "aws_security_group" "tezos_node_lb" {
  name        = format("tezos_farm_lb_%s_%s", var.TZ_NETWORK, var.TZ_MODE)
  description = format("Security group for tezos loadbalancer targeting the %s network", var.TZ_NETWORK)
  vpc_id      = data.aws_vpc.tzlink.id
}

resource "aws_security_group_rule" "http_ingress_for_loadbalancer" {
  type        = "ingress"
  from_port   = 80
  to_port     = 80
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.tezos_node_lb.id
}


resource "aws_security_group_rule" "all_egress_for_loadbalancer" {
  type        = "egress"
  from_port   = 0
  to_port     = 0
  protocol    = "-1"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.tezos_node_lb.id
}

resource "aws_security_group" "tezos_node" {
  name        = format("tezos_farm_%s_%s", var.TZ_NETWORK, var.TZ_MODE)
  description = format("Security group for tezos nodes in network %s", var.TZ_NETWORK)
  vpc_id      = data.aws_vpc.tzlink.id
}

resource "aws_security_group_rule" "rpc_ingress_for_tezos_node" {
  type                     = "ingress"
  from_port                = 8000
  to_port                  = 8000
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.tezos_node_lb.id

  security_group_id = aws_security_group.tezos_node.id
}

resource "aws_security_group_rule" "p2p_ingress_for_tezos_node" {
  type        = "ingress"
  from_port   = 9732
  to_port     = 9732
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]


  security_group_id = aws_security_group.tezos_node.id
}

resource "aws_security_group_rule" "ssh_ingress_for_tezos_node" {
  type        = "ingress"
  from_port   = 22
  to_port     = 22
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.tezos_node.id
}

resource "aws_security_group_rule" "all_egress_for_tezos_node" {
  type        = "egress"
  from_port   = 0
  to_port     = 0
  protocol    = "-1"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.tezos_node.id
}
