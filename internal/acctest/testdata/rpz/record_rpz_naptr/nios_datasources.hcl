# Auto-generated datasource acceptance-test cases for RecordRpzNaptr.
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
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.disable", "nios.flags", "nios.name", "nios.order", "nios.preference", "nios.regexp", "nios.replacement", "nios.rp_zone", "nios.services", "nios.ttl", "nios.view"]

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
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
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.comment", "nios.disable", "nios.flags", "nios.name", "nios.order", "nios.preference", "nios.regexp", "nios.replacement", "nios.rp_zone", "nios.services", "nios.ttl", "nios.view"]

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
      ext_attrs   = { Site = "{{random3}}" }
    }
  }

}
