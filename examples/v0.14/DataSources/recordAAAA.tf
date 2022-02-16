terraform {
  # Required providers block for Terraform v0.14.7
  required_providers {
    infoblox = {
      source  = "infobloxopen/infoblox"
      version = ">=2.0"
    }
  }
}

resource "infoblox_aaaa_record" "aaaa_record"{
  fqdn = "static.test.com" # the zone 'test.com' MUST exist in the DNS view
  ipv6_addr = "2000::1"
  ttl = 10
  comment = "static AAAA-record"
  ext_attrs = jsonencode({
    "Location" = "New York"
    "Site" = "HQ"
  })
}

data "infoblox_aaaa_record" "test" {
  dns_view = "default"
  ipv6_addr = infoblox_aaaa_record.aaaa_record.ipv6_addr
  fqdn = infoblox_aaaa_record.aaaa_record.fqdn
}

output "id" {
  value = data.infoblox_aaaa_record.test
}

output "zone" {
  value = data.infoblox_aaaa_record.test.zone
}

output "ttl" {
  value = data.infoblox_aaaa_record.test.ttl
}

output "comment" {
  value = data.infoblox_aaaa_record.test.comment
}

output "ext_attrs" {
  value = data.infoblox_aaaa_record.test.ext_attrs
}