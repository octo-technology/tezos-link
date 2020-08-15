terraform {
  required_version = "0.12.20"
  backend "s3" {}
}

provider "random" {
  version = "2.3.0"
}

# Network configuration

data "aws_vpc" "tzlink" {
  cidr_block = var.vpc_cidr

  tags = {
    Name    = "tzlink"
    Project = var.project_name
  }
}

data "aws_subnet_ids" "database" {
  vpc_id = data.aws_vpc.tzlink.id

  tags = {
    Name    = "tzlink-private-database-*"
    Project = var.project_name
  }
}

resource "aws_db_subnet_group" "database" {
  name       = "tzlink-database"
  subnet_ids = tolist(data.aws_subnet_ids.database.ids)
}

# Network access control

resource "aws_security_group" "database" {
  name        = "tzlink-database"
  description = "Security group for database"
  vpc_id      = data.aws_vpc.tzlink.id

  tags = {
    Name    = "tzlink-database"
    Project = var.project_name
  }
}

resource "aws_security_group_rule" "all_ingress_from_vpc" {
  type        = "ingress"
  from_port   = 0
  to_port     = 0
  protocol    = "-1"
  cidr_blocks = [var.vpc_cidr]

  security_group_id = aws_security_group.database.id
}

resource "aws_security_group_rule" "all_egress_from_vpc" {
  type        = "egress"
  from_port   = 0
  to_port     = 0
  protocol    = "-1"
  cidr_blocks = [var.vpc_cidr]

  security_group_id = aws_security_group.database.id
}

# Serverless Aurora postgresql

resource "random_password" "database_master_password" {
  length  = 32
  special = false
}

resource "aws_secretsmanager_secret" "database_master_password" {
  name        = "tzlink-database-password"
  description = "The associated password to connect on the database"
}

resource "aws_secretsmanager_secret_version" "database_master_password" {
  secret_id     = aws_secretsmanager_secret.database_master_password.id
  secret_string = random_password.database_master_password.result
}

resource "aws_rds_cluster" "database" {
  cluster_identifier = "tzlink-database"
  engine             = "aurora-postgresql"
  engine_version     = "10.7"
  availability_zones = ["${var.region}a", "${var.region}b", "${var.region}c"]

  database_name = var.database_name

  master_username = var.database_master_username
  master_password = random_password.database_master_password.result

  backup_retention_period = 7
  preferred_backup_window = "01:00-02:00"
  skip_final_snapshot     = true

  apply_immediately = true

  db_subnet_group_name = aws_db_subnet_group.database.name
  port                 = 5432

  vpc_security_group_ids = [aws_security_group.database.id]

  engine_mode = "serverless"

  scaling_configuration {
    auto_pause               = true
    max_capacity             = 16
    min_capacity             = 2
    seconds_until_auto_pause = 300
    timeout_action           = "ForceApplyCapacityChange"
  }

  # parameters which permits idempotency
  backtrack_window                    = 0
  deletion_protection                 = false
  enable_http_endpoint                = false
  enabled_cloudwatch_logs_exports     = []
  iam_database_authentication_enabled = false
  iam_roles                           = []
  storage_encrypted                   = true

  tags = {
    Name    = "tzlink-database"
    Project = var.project_name
  }

  depends_on = [aws_db_subnet_group.database]
}
