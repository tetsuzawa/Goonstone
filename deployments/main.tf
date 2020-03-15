variable "access_key" {}
variable "secret_key" {}
variable "profile" {}

variable "name" {}
variable "region" {
  default = "ap-northeast-1"
}
variable "root_domain_name" {}
variable "sub_domain_name" {}
variable "route53_hosted_zone_id" {}
variable "s3_bucket" {}
variable "alb_certificate_arn" {}
variable "cf_certificate_arn" {}
variable "vpc_id" {}
variable "vpc_cidr" {}
variable "vpc_default_route_table_id" {}
variable "igw_id" {}
variable "iam_task_role_arn" {}


variable "mysql_protocol" {}
variable "mysql_db_name" {}
variable "mysql_charset" {}
variable "mysql_loc" {}
variable "mysql_parse_time" {}
variable "mysql_user" {}
variable "mysql_password" {}
variable "frontend_host" {}
variable "frontend_port" {}
variable "api_host" {}
variable "api_port" {}

terraform {
  backend "s3" {
    bucket  = "goonstone"
    key     = "terraform.tfstate"
    region  = "ap-northeast-1"
    profile = "admin"
  }

  required_providers {
    aws = "~> 2.49"
  }
}

provider "aws" {
  access_key = var.access_key
  secret_key = var.secret_key
  region     = var.region
}

module "common" {
  source                     = "./modules/common"
  name                       = var.name
  region                     = var.region
  root_domain_name           = var.root_domain_name
  sub_domain_name            = var.sub_domain_name
  route53_hosted_zone_id     = var.route53_hosted_zone_id
  s3_bucket                  = var.s3_bucket
  iam_task_role_arn          = var.iam_task_role_arn
  alb_certificate_arn        = var.alb_certificate_arn
  cf_certificate_arn         = var.cf_certificate_arn
  vpc_id                     = var.vpc_id
  vpc_cidr                   = var.vpc_cidr
  vpc_default_route_table_id = var.vpc_default_route_table_id
  igw_id                     = var.igw_id
  mysql_user                 = var.mysql_user
  mysql_password             = var.mysql_password
  mysql_protocol             = var.mysql_protocol
  mysql_db_name              = var.mysql_db_name
  mysql_charset              = var.mysql_charset
  mysql_loc                  = var.mysql_loc
  mysql_parse_time           = var.mysql_parse_time
  frontend_host              = var.frontend_host
  frontend_port              = var.frontend_port
  api_host                   = var.api_host
  api_port                   = var.api_port
}