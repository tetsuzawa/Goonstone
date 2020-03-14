data "aws_vpc" "vpc" {
  id         = var.vpc_id
  cidr_block = var.vpc_cidr
}

