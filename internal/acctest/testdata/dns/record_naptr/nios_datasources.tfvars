# Auto-generated datasource acceptance-test cases for RecordNaptr.
case "filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
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

  pair_checks = ["nios.comment", "nios.creator", "nios.ddns_principal", "nios.ddns_protected", "nios.disable", "nios.flags", "nios.forbid_reclamation", "nios.name", "nios.order", "nios.preference", "nios.regexp", "nios.replacement", "nios.services", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      name        = "{{random2}}.$${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
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

  pair_checks = ["nios.comment", "nios.creator", "nios.ddns_principal", "nios.ddns_protected", "nios.disable", "nios.flags", "nios.forbid_reclamation", "nios.name", "nios.order", "nios.preference", "nios.regexp", "nios.replacement", "nios.services", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
      ext_attrs   = { Site = "{{random}}" }
    }
  }

}
