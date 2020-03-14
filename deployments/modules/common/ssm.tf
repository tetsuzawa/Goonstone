resource "aws_ssm_parameter" "api_port" {
  name  = "API_PORT"
  type  = "SecureString"
  value = var.api_port
}

resource "aws_ssm_parameter" "api_host" {
  name  = "API_HOST"
  type  = "SecureString"
  value = var.api_host
}

resource "aws_ssm_parameter" "mysql_user" {
  name  = "MYSQL_USER"
  type  = "SecureString"
  value = var.mysql_user
}

resource "aws_ssm_parameter" "mysql_password" {
  name  = "MYSQL_PASSWORD"
  type  = "SecureString"
  value = var.mysql_password
}

resource "aws_ssm_parameter" "mysql_db_name" {
  name  = "MYSQL_DB_NAME"
  type  = "SecureString"
  value = var.mysql_db_name
}

resource "aws_ssm_parameter" "mysql_charset" {
  name  = "MYSQL_CHARSET"
  type  = "SecureString"
  value = var.mysql_charset
}

resource "aws_ssm_parameter" "mysql_parse_time" {
  name  = "MYSQL_PARSE_TIME"
  type  = "SecureString"
  value = var.mysql_parse_time
}

resource "aws_ssm_parameter" "mysql_port" {
  name  = "MYSQL_PORT"
  type  = "SecureString"
  value = aws_db_instance.rds.port
}

resource "aws_ssm_parameter" "mysql_protocol" {
  name  = "MYSQL_PROTOCOL"
  type  = "SecureString"
  value = var.mysql_protocol
}

resource "aws_ssm_parameter" "mysql_loc" {
  name  = "MYSQL_LOC"
  type  = "SecureString"
  value = var.mysql_loc
}

resource "aws_ssm_parameter" "mysql_host" {
  name  = "MYSQL_HOST"
  type  = "SecureString"
  value = aws_db_instance.rds.address
}

