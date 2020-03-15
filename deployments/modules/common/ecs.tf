resource "aws_ecs_cluster" "ecs_cluster" {
  name = "${var.name}-ecs-cluster"
}

# Services

# Frontend
resource "aws_ecs_service" "frontend_service" {
  name                               = "${var.name}-frontend-service"
  cluster                            = aws_ecs_cluster.ecs_cluster.id
  task_definition                    = aws_ecs_task_definition.frontend.arn
  launch_type                        = "FARGATE"
  desired_count                      = 1
  deployment_minimum_healthy_percent = 100
  deployment_maximum_percent         = 200
  health_check_grace_period_seconds  = 60

  network_configuration {
    subnets = [aws_subnet.public_a.id, aws_subnet.public_c.id]
    security_groups = [
    aws_security_group.frontend.id]
    assign_public_ip = "true"
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.frontend.arn
    container_name   = "${var.name}-frontend-container"
    container_port   = var.frontend_port
  }

  depends_on = [aws_ecs_service.api_service]
}

# api
resource "aws_ecs_service" "api_service" {
  name                               = "${var.name}-api-service"
  cluster                            = aws_ecs_cluster.ecs_cluster.id
  task_definition                    = aws_ecs_task_definition.api.arn
  launch_type                        = "FARGATE"
  desired_count                      = 1
  deployment_minimum_healthy_percent = 100
  deployment_maximum_percent         = 200
  health_check_grace_period_seconds  = 60

  network_configuration {
    subnets = [
    aws_subnet.public_a.id, aws_subnet.public_c.id]
    security_groups = [
    aws_security_group.api.id]
    assign_public_ip = "true"
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.api.arn
    container_name   = "${var.name}-api-container"
    container_port   = var.api_port
  }

  depends_on = [aws_lb.alb]
}