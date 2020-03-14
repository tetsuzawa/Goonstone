# Task role
data "aws_arn" "task_role" {
  arn = var.iam_task_role_arn
}
