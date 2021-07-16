terraform {
  # Required providers block for Terraform v0.14.7
  required_providers {
    infoblox = {
      source  = "terraform-providers/infoblox"
      version = ">= 1.0"
    }
  }
}

resource "infoblox_a_record" "a_record"{
  fqdn = "static.test.com" # the zone 'test.com' MUST exist in the DNS view
  ip_addr = "192.168.31.31"
  ttl = 10
  comment = "static A-record"
  ext_attrs = jsonencode({
    "Location" = "New York"
    "Site" = "HQ"
  })
}

data "infoblox_a_record" "test" {
  dns_view = "default"
  ip_addr = infoblox_a_record.a_record.ip_addr
  fqdn = infoblox_a_record.a_record.fqdn
}

output "id" {
  value = data.infoblox_a_record.test
}

output "zone" {
  value = data.infoblox_a_record.test.zone
}

output "ttl" {
  value = data.infoblox_a_record.test.ttl
}

output "comment" {
  value = data.infoblox_a_record.test.comment
}

output "ext_attrs" {
  value = data.infoblox_a_record.test.ext_attrs
}