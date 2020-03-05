resource "aws_db_subnet_group" "database" {
  name       = format("tzlink-%s-database-test", var.ENV)
  subnet_ids = tolist(data.aws_subnet_ids.database.ids)
}

resource "aws_db_instance" "database" {
  engine         = "postgres"
  engine_version = "11.6"

  identifier = "db-test"

  username = "tmpadmin"
  password = local.vaulted_user_password

  instance_class = "db.t2.micro"

  db_subnet_group_name = format("tzlink-%s-database-test", var.ENV)
  publicly_accessible  = false
  port                 = 5432

  allocated_storage = 10
  storage_type      = "gp2"

  name = "tezoslink"

  parameter_group_name = "default.postgres11"

  depends_on = [aws_db_subnet_group.database]
}