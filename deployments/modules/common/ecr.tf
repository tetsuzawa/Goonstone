# Frontend
resource "aws_ecr_repository" "frontend" {
  name = "${var.name}-frontend"
}

# API
resource "aws_ecr_repository" "api" {
  name = "${var.name}-api"
}