# Default(Private)
resource "aws_default_route_table" "default_route_table" {
  default_route_table_id = aws_vpc.vpc.default_route_table_id

  tags {
    Name    = "${var.name}-default-route-table"
    Product = var.name
  }
}

resource "aws_route_table_association" "route_private_a" {
  route_table_id = aws_default_route_table.default_route_table.id
  subnet_id      = aws_subnet.private_a.id
}

resource "aws_route_table_association" "route_private_c" {
  route_table_id = aws_default_route_table.default_route_table.id
  subnet_id      = aws_subnet.private_c.id
}

# Public
resource "aws_route_table" "public_route_table" {
  vpc_id = aws_vpc.vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.main_igw.id
  }

  tags {
    Name    = "${var.name}-public-route-table"
    Product = var.name
    Env     = terraform.workspace
  }
}

resource "aws_route_table_association" "route-public-a" {
  route_table_id = aws_route_table.public_route_table.id
  subnet_id      = aws_subnet.public_a.id
}

resource "aws_route_table_association" "route-public-c" {
  route_table_id = aws_route_table.public_route_table.id
  subnet_id      = aws_subnet.public_c.id
}