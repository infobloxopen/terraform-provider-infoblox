terraform {
  # Required providers block for Terraform v0.14.7
  required_providers {
    infoblox = {
      source  = "terraform-providers/infoblox"
      version = ">= 1.0"
    }
  }
}

resource "infoblox_cname_record" "foo"{
  alias="test.a.com"
  canonical="test-name.a.com"	
  comment = "CNAME example rec"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}

data "infoblox_cname_record" "test" {
	dns_view="default"
	alias=infoblox_cname_record.foo.alias
	canonical=infoblox_cname_record.foo.canonical
}

output "id" {
  value = data.infoblox_cname_record.test
}

output "zone" {
  value = data.infoblox_cname_record.test.zone
}

output "ttl" {
  value = data.infoblox_cname_record.test.ttl
}

output "comment" {
  value = data.infoblox_cname_record.test.comment
}

output "ext_attrs" {
  value = data.infoblox_cname_record.test.ext_attrs
}

