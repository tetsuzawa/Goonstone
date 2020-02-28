# Front
resource "aws_ecr_repository" "front" {
  name = "${var.name}-front"

  provisioner "local-exec" {
    command = "build-image.sh front ${self.repository_url} ${self.name}"
  }
}

# API
resource "aws_ecr_repository" "api" {
  name = "${var.name}-api"

  provisioner "local-exec" {
    command = "build-image.sh api ${self.repository_url} ${self.name}"
  }
}