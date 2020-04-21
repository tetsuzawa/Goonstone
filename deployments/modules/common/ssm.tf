resource "aws_ssm_parameter" "access_key" {
  name  = "${upper(var.name)}_AWS_ACCESS_KEY"
  type  = "SecureString"
  value = var.access_key
}

resource "aws_ssm_parameter" "secret_access_key" {
  name  = "${upper(var.name)}_AWS_SECRET_ACCESS_KEY"
  type  = "SecureString"
  value = var.secret_key
}

resource "aws_ssm_parameter" "s3_bucket" {
  name  = "${upper(var.name)}_AWS_S3_BUCKET"
  type  = "SecureString"
  value = var.s3_bucket
}

resource "aws_ssm_parameter" "mysql_host" {
  name  = "${upper(var.name)}_MYSQL_HOST"
  type  = "SecureString"
  value = aws_db_instance.rds.address
}

resource "aws_ssm_parameter" "mysql_port" {
  name  = "${upper(var.name)}_MYSQL_PORT"
  type  = "SecureString"
  value = aws_db_instance.rds.port
}

resource "aws_ssm_parameter" "mysql_protocol" {
  name  = "${upper(var.name)}_MYSQL_PROTOCOL"
  type  = "SecureString"
  value = var.mysql_protocol
}

resource "aws_ssm_parameter" "mysql_user" {
  name  = "${upper(var.name)}_MYSQL_USER"
  type  = "SecureString"
  value = var.mysql_user
}

resource "aws_ssm_parameter" "mysql_password" {
  name  = "${upper(var.name)}_MYSQL_PASSWORD"
  type  = "SecureString"
  value = var.mysql_password
}

resource "aws_ssm_parameter" "mysql_db_name" {
  name  = "${upper(var.name)}_MYSQL_DB_NAME"
  type  = "SecureString"
  value = var.mysql_db_name
}

resource "aws_ssm_parameter" "mysql_charset" {
  name  = "${upper(var.name)}_MYSQL_CHARSET"
  type  = "SecureString"
  value = var.mysql_charset
}


resource "aws_ssm_parameter" "mysql_loc" {
  name  = "${upper(var.name)}_MYSQL_LOC"
  type  = "SecureString"
  value = var.mysql_loc
}


resource "aws_ssm_parameter" "mysql_parse_time" {
  name  = "${upper(var.name)}_MYSQL_PARSE_TIME"
  type  = "SecureString"
  value = var.mysql_parse_time
}

resource "aws_ssm_parameter" "redis_host" {
  name  = "${upper(var.name)}_REDIS_HOST"
  type  = "SecureString"
  value = aws_elasticache_cluster.elasticache.cache_nodes.0.address
}

resource "aws_ssm_parameter" "redis_port" {
  name  = "${upper(var.name)}_REDIS_PORT"
  type  = "SecureString"
  value = aws_elasticache_cluster.elasticache.cache_nodes.0.port
}

resource "aws_ssm_parameter" "redis_protocol" {
  name  = "${upper(var.name)}_REDIS_PROTOCOL"
  type  = "SecureString"
  value = var.redis_protocol
}

resource "aws_ssm_parameter" "frontend_host" {
  name  = "${upper(var.name)}_FRONTEND_HOST"
  type  = "SecureString"
  value = var.frontend_host
}

resource "aws_ssm_parameter" "frontend_port" {
  name  = "${upper(var.name)}_FRONTEND_PORT"
  type  = "SecureString"
  value = var.frontend_port
}

resource "aws_ssm_parameter" "api_protocol" {
  name  = "${upper(var.name)}_API_PROTOCOL"
  type  = "SecureString"
  value = var.api_protocol
}

resource "aws_ssm_parameter" "api_host" {
  name  = "${upper(var.name)}_API_HOST"
  type  = "SecureString"
  value = var.api_host
}

resource "aws_ssm_parameter" "api_port" {
  name  = "${upper(var.name)}_API_PORT"
  type  = "SecureString"
  value = var.api_port
}

resource "aws_ssm_parameter" "api_base_root" {
  name  = "${upper(var.name)}_API_BASE_ROOT"
  type  = "SecureString"
  value = var.api_base_root
}
