resource "aws_db_subnet_group" "database" {
  name       = "tzlink-database"
  subnet_ids = tolist(data.aws_subnet_ids.database.ids)
}

resource "aws_rds_cluster" "database" {
  cluster_identifier = "tzlink-database"
  engine             = "aurora-postgresql"
  engine_version     = "10.7"
  availability_zones = ["eu-west-1a", "eu-west-1b", "eu-west-1c"]

  database_name = var.DATABASE_TABLE

  master_username = var.DATABASE_USERNAME
  master_password = var.DATABASE_PASSWORD

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
    Name      = "tzlink-database"
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }

  depends_on = [aws_db_subnet_group.database]
}
