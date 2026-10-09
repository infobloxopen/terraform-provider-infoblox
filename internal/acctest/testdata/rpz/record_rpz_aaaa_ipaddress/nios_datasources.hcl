# Hand-authored datasource acceptance-test cases for RecordRpzAaaaIpaddress.
case "filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  filter {
    type = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.disable", "nios.ipv6addr", "nios.name", "nios.rp_zone", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      name     = "{{random_ipv6_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      ipv6addr = "2001:db8::10"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  filter {
    type = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.comment", "nios.disable", "nios.ipv6addr", "nios.name", "nios.rp_zone", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      name      = "{{random_ipv6_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      ipv6addr  = "2001:db8::10"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      ext_attrs = { Site = "value1" }
    }
  }

}
