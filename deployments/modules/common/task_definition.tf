# Frontend
data "template_file" "frontend" {
  template = file("${path.root}/task-definitions/frontend-task-definition.json")

  vars = {
    product              = var.name
    image-url            = aws_ecr_repository.frontend.repository_url
    cpu                  = var.frontend_cpu
    memory               = var.frontend_memory
    cloudwatch_log_group = aws_cloudwatch_log_group.frontend.name
  }
}

resource "aws_ecs_task_definition" "frontend" {
  family                   = "${var.name}-frontend-task"
  container_definitions    = data.template_file.frontend.rendered
  task_role_arn            = data.aws_arn.task_role.arn
  execution_role_arn       = data.aws_arn.task_role.arn
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = var.frontend_cpu
  memory                   = var.frontend_memory

  provisioner "local-exec" {
    command = "./push-image.sh ${aws_ecr_repository.frontend.repository_url} ${aws_ecr_repository.api.repository_url}"
  }
}

# API
data "template_file" "api" {
  template = file("${path.root}/task-definitions/api-task-definition.json")

  vars = {
    product              = var.name
    image-url            = aws_ecr_repository.api.repository_url
    cpu                  = var.api_cpu
    memory               = var.api_memory
    cloudwatch_log_group = aws_cloudwatch_log_group.api.name
  }
}

resource "aws_ecs_task_definition" "api" {
  family                   = "${var.name}-api-task"
  container_definitions    = data.template_file.api.rendered
  task_role_arn            = data.aws_arn.task_role.arn
  execution_role_arn       = data.aws_arn.task_role.arn
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = var.api_cpu
  memory                   = var.api_memory

  provisioner "local-exec" {
    command = "./push-image.sh ${aws_ecr_repository.frontend.repository_url} ${aws_ecr_repository.api.repository_url}"
  }
}