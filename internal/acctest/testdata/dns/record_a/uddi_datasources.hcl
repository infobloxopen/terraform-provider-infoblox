# Auto-generated datasource acceptance-test cases for RecordA.
case "filters" {
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  filter {
    type   = "filters"
    values = {
      name_in_zone = "uddi.name_in_zone"
      zone         = "uddi.zone"
    }
  }

  pair_checks = ["uddi.absolute_name_spec", "uddi.comment", "uddi.disabled", "uddi.name_in_zone", "uddi.ttl", "uddi.type", "uddi.view", "uddi.zone"]

  step {
    uddi {
      name_in_zone = "{{random2}}"
      zone         = infoblox_zone_auth.test.id
      rdata        = { address = "{{random_ip}}" }
    }
  }

}

case "tag_filters" {
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  filter {
    type   = "tag_filters"
    values = {
      tag1 = "uddi.tags.tag1"
    }
  }

  pair_checks = ["uddi.absolute_name_spec", "uddi.comment", "uddi.disabled", "uddi.name_in_zone", "uddi.ttl", "uddi.type", "uddi.view", "uddi.zone"]

  step {
    uddi {
      zone  = infoblox_zone_auth.test.id
      rdata = { address = "{{random_ip}}" }
      tags  = { tag1 = "{{random}}" }
    }
  }

}
