resource "aws_internet_gateway" "main_igw" {
  vpc_id = aws_vpc.vpc.id

  tags {
    Name    = "${var.name}-main-igw"
    Product = var.name
  }
}