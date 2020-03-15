resource "aws_subnet" "public_a" {
  availability_zone = "${var.region}a"
  cidr_block        = cidrsubnet(data.aws_vpc.vpc.cidr_block, 8, 1)
  vpc_id            = var.vpc_id

  tags = {
    "Name"    = "${var.name}-subnet-public-a"
    "Product" = var.name
  }
}

resource "aws_subnet" "public_c" {
  availability_zone = "${var.region}c"
  cidr_block        = cidrsubnet(data.aws_vpc.vpc.cidr_block, 8, 2)
  vpc_id            = var.vpc_id

  tags = {
    "Name"    = "${var.name}-subnet-public-c"
    "Product" = var.name
  }
}

resource "aws_subnet" "private_a" {
  availability_zone = "${var.region}a"
  cidr_block        = cidrsubnet(data.aws_vpc.vpc.cidr_block, 8, 65)
  vpc_id            = var.vpc_id

  tags = {
    "Name"    = "${var.name}-subnet-private-a"
    "Product" = var.name
  }
}

resource "aws_subnet" "private_c" {
  availability_zone = "${var.region}c"
  cidr_block        = cidrsubnet(data.aws_vpc.vpc.cidr_block, 8, 66)
  vpc_id            = var.vpc_id

  tags = {
    "Name"    = "${var.name}-subnet-private-c"
    "Product" = var.name
  }
}
