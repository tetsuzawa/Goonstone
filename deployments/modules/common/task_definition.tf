# Frontend
data "template_file" "front" {
  template = file("${path.root}/task-definitions/front.json")

  vars = {
    product   = var.name
    image-url = aws_ecr_repository.front.repository_url
    cpu       = var.front_cpu
    memory    = var.front_memory
  }
}

resource "aws_ecs_task_definition" "front" {
  family                   = "${var.name}-front-task"
  container_definitions    = data.template_file.front.rendered
  task_role_arn            = aws_iam_role.task-role.arn
  execution_role_arn       = aws_iam_role.task-role.arn
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = var.front_cpu
  memory                   = var.front_memory

  provisioner "local-exec" {
    command = "./push-image.sh ${var.name} front ${aws_ecr_repository.front.name} ${aws_ecr_repository.front.repository_url}"
  }
}

# API
data "template_file" "api" {
  template = file("${path.root}/task-definitions/api.json")

  vars =  {
    product   = var.name
    image-url = aws_ecr_repository.api.repository_url
    cpu       = var.api_cpu
    memory    = var.api_memory
  }
}

resource "aws_ecs_task_definition" "api" {
  family                   = "${var.name}-api-task"
  container_definitions    = data.template_file.api.rendered
  task_role_arn            = aws_iam_role.task-role.arn
  execution_role_arn       = aws_iam_role.task-role.arn
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = var.api_cpu
  memory                   = var.api_memory

  provisioner "local-exec" {
    command = "./push-image.sh ${var.name} api ${aws_ecr_repository.api.name} ${aws_ecr_repository.api.repository_url}"
  }
}