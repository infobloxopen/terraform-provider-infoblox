# Auto-generated datasource acceptance-test cases for RecordAaaa.
case "filters" {
  backend           = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
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

  pair_checks = ["nios.comment", "nios.creator", "nios.ddns_principal", "nios.ddns_protected", "nios.disable", "nios.forbid_reclamation", "nios.ipv6addr", "nios.name", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      name     = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      view     = infoblox_zone_auth.test.nios.view
    }
  }

}

case "ext_attr_filters" {
  backend           = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
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

  pair_checks = ["nios.comment", "nios.creator", "nios.ddns_principal", "nios.ddns_protected", "nios.disable", "nios.forbid_reclamation", "nios.ipv6addr", "nios.name", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      name      = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr  = "{{random_ipv6}}"
      view      = infoblox_zone_auth.test.nios.view
      ext_attrs = { Site = "{{random2}}" }
    }
  }

}
