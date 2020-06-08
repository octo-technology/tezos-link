resource "aws_db_subnet_group" "database" {
  name       = "tzlink-database"
  subnet_ids = tolist(data.aws_subnet_ids.database.ids)
}

resource "aws_rds_cluster" "database" {
  cluster_identifier = "tzlink-database"
  engine             = "aurora-postgresql"
  engine_version     = "10.7"
  availability_zones = ["eu-west-1a", "eu-west-1b"]

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
    max_capacity             = 192
    min_capacity             = 2
    seconds_until_auto_pause = 300
    timeout_action           = "ForceApplyCapacityChange"
  }

  tags = {
    Name      = "tzlink-database"
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }

  depends_on = [aws_db_subnet_group.database]
}

# resource "aws_db_instance" "database" {
#   engine         = "postgres" #
#   engine_version = "11.6" #

#   identifier = "tzlink-database" #

#   username = var.DATABASE_USERNAME #
#   password = var.DATABASE_PASSWORD #

#   instance_class = "db.m4.large" #

#   db_subnet_group_name = aws_db_subnet_group.database.name #
#   publicly_accessible  = false #
#   port                 = 5432

#   vpc_security_group_ids = [aws_security_group.database.id] #

#   allocated_storage = 20
#   storage_type      = "gp2"

#   name = var.DATABASE_TABLE #

#   backup_retention_period = 7 #
#   backup_window           = "01:00-02:00" #
#   skip_final_snapshot     = true #

#   apply_immediately = true #

#   parameter_group_name = "default.postgres11"

#   auto_minor_version_upgrade = true


#   tags = {
#     Name      = "tzlink-database"
#     Project   = var.PROJECT_NAME
#     BuildWith = var.BUILD_WITH
#   }

#   depends_on = [aws_db_subnet_group.database]
# }