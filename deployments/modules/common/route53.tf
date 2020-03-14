data aws_route53_zone "main" {
  zone_id = var.route53_hosted_zone_id
}


resource "aws_route53_record" "main" {
  zone_id = data.aws_route53_zone.main.zone_id
  name    = var.sub_domain_name == "" ? var.root_domain_name : "${var.sub_domain_name}.${var.root_domain_name}"
  type    = "A"

  alias {
    name                   = aws_lb.alb.dns_name
    zone_id                = aws_lb.alb.zone_id
    evaluate_target_health = true
  }
}
