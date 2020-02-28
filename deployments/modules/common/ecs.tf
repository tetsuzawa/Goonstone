resource "aws_ecs_cluster" "ecs_cluster" {
  name = "${var.name}-ecs-cluster"
}

# Services

# Frontend
resource "aws_ecs_service" "front-service" {
  name                               = "${var.name}-front-service"
  cluster                            = aws_ecs_cluster.ecs_cluster.id
  task_definition                    = aws_ecs_task_definition.front.arn
  launch_type                        = "FARGATE"
  desired_count                      = 1
  deployment_minimum_healthy_percent = 100
  deployment_maximum_percent         = 200
  health_check_grace_period_seconds  = 60

  network_configuration {
    subnets          = [aws_subnet.public_a.id, aws_subnet.public_c.id]
    security_groups  = [
      aws_security_group.front.id]
    assign_public_ip = "true"
  }

  load_balancer {
    target_group_arn = aws_alb_target_group.front.arn
    container_name   = "${var.name}-front-container"
    container_port   = 80
  }

  depends_on = ["aws_ecs_service.api-service"]
}

# API
resource "aws_ecs_service" "api-service" {
  name                               = "${var.name}-api-service"
  cluster                            = aws_ecs_cluster.ecs_cluster.id
  task_definition                    = aws_ecs_task_definition.api.arn
  launch_type                        = "FARGATE"
  desired_count                      = 1
  deployment_minimum_healthy_percent = 100
  deployment_maximum_percent         = 200

  network_configuration {
    subnets          = [
      aws_subnet.public_a.id, aws_subnet.public_c.id]
    security_groups  = [
      aws_security_group.api.id]
    assign_public_ip = "true"
  }

  load_balancer {
    target_group_arn = aws_alb_target_group.api.arn
    container_name   = "${var.name}-api-container"
    container_port   = 8080
  }

  depends_on = ["aws_alb.alb"]
}