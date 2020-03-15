# Frontend
resource "aws_ecr_repository" "frontend" {
  name = "${var.name}-api"

  provisioner "local-exec" {
    command = "./push-image.sh ${self.repository_url}-frontend"
  }
}

# api
resource "aws_ecr_repository" "api" {
  name = "${var.name}-api"

  provisioner "local-exec" {
    command = "./push-image.sh ${self.repository_url}-api"
  }
}