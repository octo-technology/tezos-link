resource "aws_db_subnet_group" "database" {
  name       = format("tzlink-%s-database", var.ENV)
  subnet_ids = tolist(data.aws_subnet_ids.database.ids)
}

resource "aws_db_instance" "database" {
  engine         = "postgres"
  engine_version = "11.6"

  identifier = format("tzlink-%s-database", var.ENV)

  username = var.DATABASE_USERNAME
  password = var.DATABASE_PASSWORD

  instance_class = "db.m4.large"

  db_subnet_group_name = aws_db_subnet_group.database.name
  publicly_accessible  = false
  port                 = 5432

  vpc_security_group_ids = [aws_security_group.database.id]

  allocated_storage = 20
  storage_type      = "gp2"

  name = var.DATABASE_TABLE
  
  backup_retention_period = 7
  backup_window = "01:00-02:00"
  skip_final_snapshot = true

  apply_immediately = true

  parameter_group_name = "default.postgres11"

  auto_minor_version_upgrade = true


  tags = {
    Name        = format("tzlink-%s-database", var.ENV)
    Project     = var.PROJECT_NAME
    Environment = var.ENV
    BuildWith   = var.BUILD_WITH
  }

  depends_on = [aws_db_subnet_group.database]
}