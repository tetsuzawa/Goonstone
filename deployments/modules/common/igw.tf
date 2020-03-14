data "aws_internet_gateway" "main" {
  internet_gateway_id = var.igw_id
}
