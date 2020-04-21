# Default
resource "aws_default_security_group" "security_group_default" {
  vpc_id = data.aws_vpc.vpc.id

  tags = {
    Name    = "${var.name}-security-group-default"
    Product = var.name
  }
}

# for ALB
resource "aws_security_group" "alb" {
  name   = "${var.name}-security-group-alb"
  vpc_id = data.aws_vpc.vpc.id

  tags = {
    Name    = "${var.name}-security-group-alb"
    Product = var.name
  }
}

# In:   All HTTP(port 80)
resource "aws_security_group_rule" "alb_in_rule_http80" {
  security_group_id = aws_security_group.alb.id
  type              = "ingress"
  from_port         = 80
  to_port           = 80
  protocol          = "tcp"
  cidr_blocks       = ["0.0.0.0/0"]
}

# In:   All HTTP(port 8080)
resource "aws_security_group_rule" "alb_in_rule_http8080" {
  security_group_id = aws_security_group.alb.id
  type              = "ingress"
  from_port         = 8080
  to_port           = 8080
  protocol          = "tcp"
  cidr_blocks       = ["0.0.0.0/0"]
}

# In:   All HTTPS
resource "aws_security_group_rule" "alb_in_rule_https443" {
  security_group_id = aws_security_group.alb.id
  type              = "ingress"
  from_port         = 443
  to_port           = 443
  protocol          = "tcp"
  cidr_blocks       = ["0.0.0.0/0"]
}

# Out:  ALL OK
resource "aws_security_group_rule" "alb_out_rule_all" {
  security_group_id = aws_security_group.alb.id
  type              = "egress"
  from_port         = 0
  to_port           = 0
  protocol          = "-1"
  cidr_blocks       = ["0.0.0.0/0"]
}

# for Frontend
resource "aws_security_group" "frontend" {
  name   = "${var.name}-security-group-frontend"
  vpc_id = data.aws_vpc.vpc.id

  tags = {
    Name    = "${var.name}-security-group-frontend"
    Product = var.name
  }
}

# In:   ALB
resource "aws_security_group_rule" "frontend_in_rule_alb" {
  security_group_id        = aws_security_group.frontend.id
  type                     = "ingress"
  from_port                = 0
  to_port                  = 0
  protocol                 = "-1"
  source_security_group_id = aws_security_group.alb.id
}

# Out:  ALL OK
resource "aws_security_group_rule" "frontend_out_rule_all" {
  security_group_id = aws_security_group.frontend.id
  type              = "egress"
  from_port         = 0
  to_port           = 0
  protocol          = "-1"
  cidr_blocks       = ["0.0.0.0/0"]
}

# for API
resource "aws_security_group" "api" {
  name   = "${var.name}-security-group-api"
  vpc_id = data.aws_vpc.vpc.id

  tags = {
    Name    = "${var.name}-security-group-api"
    Product = var.name
  }
}

# In:   ALB
resource "aws_security_group_rule" "security_group_api_in_rule_alb" {
  security_group_id        = aws_security_group.api.id
  type                     = "ingress"
  from_port                = 0
  to_port                  = 0
  protocol                 = "-1"
  source_security_group_id = aws_security_group.alb.id
}

# Out:  ALL OK
resource "aws_security_group_rule" "security_group_api_out_rule_all" {
  security_group_id = aws_security_group.api.id
  type              = "egress"
  from_port         = 0
  to_port           = 0
  protocol          = "-1"
  cidr_blocks       = ["0.0.0.0/0"]
}

# for RDS
resource "aws_security_group" "rds" {
  name   = "${var.name}-security-group-rds"
  vpc_id = data.aws_vpc.vpc.id


  tags = {
    Name    = "${var.name}-security-group-rds"
    Product = var.name
  }
}

# for ElastiCache
resource "aws_security_group" "elasticache" {
  name   = "${var.name}-security-group-elasticache"
  vpc_id = data.aws_vpc.vpc.id


  tags = {
    Name    = "${var.name}-security-group-elasticache"
    Product = var.name
  }
}

# In:   API / DB
resource "aws_security_group_rule" "rds_in_rule_api" {
  security_group_id        = aws_security_group.rds.id
  type                     = "ingress"
  from_port                = 3306
  to_port                  = 3306
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.api.id
}

# Out:  ALL OK
resource "aws_security_group_rule" "rds_out_rule_all" {
  security_group_id = aws_security_group.rds.id
  type              = "egress"
  from_port         = 0
  to_port           = 0
  protocol          = "-1"
  cidr_blocks       = ["0.0.0.0/0"]
}

# In:   API / ElastiCache
resource "aws_security_group_rule" "elasticache_in_rule_api" {
  security_group_id        = aws_security_group.elasticache.id
  type                     = "ingress"
  from_port                = 6379
  to_port                  = 6379
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.api.id
}

# Out:  ALL OK
resource "aws_security_group_rule" "elasticache_out_rule_all" {
  security_group_id = aws_security_group.elasticache.id
  type              = "egress"
  from_port         = 0
  to_port           = 0
  protocol          = "-1"
  cidr_blocks       = ["0.0.0.0/0"]
}
