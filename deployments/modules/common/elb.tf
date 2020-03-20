resource "aws_lb" "alb" {
  name               = "${var.name}-alb"
  internal           = false
  load_balancer_type = "application"
  security_groups = [
  aws_security_group.alb.id]
  subnets = [
  aws_subnet.public_a.id, aws_subnet.public_c.id]

  tags = {
    Name    = "${var.name}-alb"
    Product = var.name
  }
}

# ALB target group

resource "aws_lb_target_group" "frontend" {
  name        = "${var.name}-alb-tg-frontend"
  port        = var.frontend_port
  protocol    = "HTTP"
  vpc_id      = data.aws_vpc.vpc.id
  target_type = "ip"

  health_check {
    interval            = 60
    path                = "/"
    protocol            = "HTTP"
    timeout             = 20
    unhealthy_threshold = 4
    matcher             = 200
  }

  tags = {
    Name    = "${var.name}-alb-tg-frontend"
    Product = var.name
  }
}

resource "aws_lb_target_group" "api" {
  name        = "${var.name}-alb-tg-api"
  port        = var.api_port
  protocol    = "HTTP"
  vpc_id      = data.aws_vpc.vpc.id
  target_type = "ip"

  health_check {
    interval            = 60
    path                = "/api/ping"
    protocol            = "HTTP"
    timeout             = 20
    unhealthy_threshold = 4
    matcher             = 200
  }

  tags = {
    Name    = "${var.name}-alb-tg-api"
    Product = var.name
  }
}

# ALB Listener
resource "aws_lb_listener" "frontend" {
  load_balancer_arn = aws_lb.alb.arn
  port              = var.frontend_port
  protocol          = "HTTP"

  default_action {
    target_group_arn = aws_lb_target_group.frontend.arn
    type             = "forward"
  }
}

resource "aws_lb_listener" "api" {
  load_balancer_arn = aws_lb.alb.arn
  port              = var.api_port
  protocol          = "HTTP"

  default_action {
    target_group_arn = aws_lb_target_group.api.arn
    type             = "forward"
  }
}