# Subnet group
resource "aws_elasticache_subnet_group" "elasticache_subnet_group" {
  name = "${var.name}-elasticache-subnet-group"
  subnet_ids = [
    aws_subnet.private_a.id,
  aws_subnet.private_c.id]
}

# Parameter group
resource "aws_elasticache_parameter_group" "elasticache_params" {
  name   = "${var.name}-elasticache-params"
  family = "redis5.0"
}

# Write Cache
resource "aws_elasticache_cluster" "elasticache" {
  cluster_id           = "${var.name}-elasticache"
  engine               = "redis"
  node_type            = "cache.t2.micro"
  num_cache_nodes      = 1
  engine_version       = "5.0.6"
  port                 = 6379
  parameter_group_name = aws_elasticache_parameter_group.elasticache_params.name
  subnet_group_name    = aws_elasticache_subnet_group.elasticache_subnet_group.name
  security_group_ids   = [aws_security_group.elasticache.id]
  availability_zone    = "${var.region}a"

  tags = {
    Name    = "${var.name}-elasticache"
    Product = var.name
  }
}
