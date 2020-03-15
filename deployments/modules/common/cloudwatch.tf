# Frontend
resource "aws_cloudwatch_log_group" "frontend" {
  name = "/ecs/${var.name}-frontend"
  tags = {
    Name    = "${var.name}-frontend-logs"
    Product = var.name
  }
}

# API
resource "aws_cloudwatch_log_group" "api" {
  name = "/ecs/${var.name}-api"
  tags = {
    Name    = "${var.name}-api-logs"
    Product = var.name
  }
}
