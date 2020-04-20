variable "access_key" {
  type        = string
  description = "AWS Access Key"
}

variable "secret_key" {
  type        = string
  description = "AWS Secret Access Key"
}

variable "name" {
  type        = string
  default     = "sample"
  description = "Product name"
}

variable "region" {
  type        = string
  default     = "ap-northeast-1"
  description = "Main region"
}

variable "root_domain_name" {
  type        = string
  default     = "example.com"
  description = "Root domain name"
}

variable "sub_domain_name" {
  type        = string
  default     = "sub"
  description = "Sub domain name"
}

variable "route53_hosted_zone_id" {
  type        = string
  default     = "Z2S***********"
  description = "ID of hosted zone which your domain will be hosted "
}

variable "s3_bucket" {
  type        = string
  default     = "sub.example.com"
  description = "Backend bucket name"
}

variable "vpc_id" {
  type        = string
  default     = "vpc-xxxxxxxx"
  description = "ID of vpc"
}

variable "vpc_cidr" {
  type        = string
  default     = "10.0.0.0/16"
  description = "cidr of vpc"
}

variable "vpc_default_route_table_id" {
  type        = string
  default     = "rtb-xxxxxxxxxxxxxxxxxxx"
  description = "default route table ID"
}

variable "igw_id" {
  type        = string
  default     = "igw-xxxxxxxxxxxxxxxxxxx"
  description = "internet gateway ID"
}
variable "iam_task_role_arn" {
  type        = string
  default     = "arn:aws:iam::*************:role/*****************"
  description = "ARN of IAM task role"
}

variable "alb_certificate_arn" {
  type        = string
  default     = "arn:aws:acm:ap-northeast-1:xxxxxxxxxxxx:certificate/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxx"
  description = "ARN certification of ALB"
}

variable "cf_certificate_arn" {
  default = "arn:aws:acm:us-east-1:xxxxxxxxxxxx:certificate/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxx"
}

variable "frontend_cpu" {
  type        = string
  default     = "256"
  description = "CPU size of frontend task"
}

variable "frontend_memory" {
  type        = string
  default     = "512"
  description = "Memory size of frontend task"
}

variable "api_cpu" {
  type        = string
  default     = "256"
  description = "CPU size of api task"
}

variable "api_memory" {
  type        = string
  default     = "512"
  description = "Memory size of api task"
}

variable "mysql_db_storage_size" {
  type        = string
  default     = "20"
  description = "Storage size of DB"
}

variable "mysql_db_name" {
  type        = string
  default     = "mydb"
  description = "Name of DB"
}

variable "mysql_protocol" {
  type = string
}

variable "mysql_user" {
  type        = string
  default     = "user"
  description = "User name of DB"
}

variable "mysql_password" {
  type        = string
  default     = "pass"
  description = "Password of DB"
}

variable "mysql_charset" {
  type = string
}

variable "mysql_loc" {
  type = string
}

variable "mysql_parse_time" {
  type = string
}

variable "redis_protocol" {
  type = string
}

variable "frontend_host" {
  type = string
}

variable "frontend_port" {
  type = string
}
variable "api_protocol" {
  type = string
}

variable "api_host" {
  type = string
}

variable "api_port" {
  type = string
}

variable "api_base_root" {
  type = string
}
