resource "aws_cloudwatch_log_group" "ecs" {
  name = "/ecs/${var.name}"
  tags = {
    Name    = "${var.name}-ecs-logs"
    Product = var.name
  }
}