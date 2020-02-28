# Subnet group

resource "aws_db_subnet_group" "rds_subnet_group" {
  name       = "${var.name}-rds-subnet-group"
  subnet_ids = [
    aws_subnet.private_a.id, aws_subnet.private_c.id]
}

# Parameter group

resource "aws_db_parameter_group" "rds_params" {
  name   = "${var.name}-rds-params"
  family = "mysql5.7"
}

# Write DB

resource "aws_db_instance" "rds" {
  identifier              = "${var.name}-rds"
  allocated_storage       = var.db_storage_size
  storage_type            = "gp2"
  engine                  = "mysql"
  engine_version          = "5.7"
  instance_class          = "db.t2.micro"
  name                    = var.db_name
  username                = var.db_user
  password                = var.db_pass
  parameter_group_name    = aws_db_parameter_group.rds_params.name
  db_subnet_group_name    = aws_db_subnet_group.rds_subnet_group.name
  vpc_security_group_ids  = [aws_security_group.rds.id]
  availability_zone       = "${var.region}a"
  multi_az                = false
  backup_retention_period = 7
  skip_final_snapshot     = true

  tags {
    Name    = "${var.name}-rds"
    Product = var.name
  }
}