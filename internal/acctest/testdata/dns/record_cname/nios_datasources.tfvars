# Auto-generated datasource acceptance-test cases for RecordCname.
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
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.canonical", "nios.comment", "nios.creator", "nios.ddns_principal", "nios.ddns_protected", "nios.disable", "nios.forbid_reclamation", "nios.name", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      canonical = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view      = infoblox_zone_auth.test.nios.view
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
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.canonical", "nios.comment", "nios.creator", "nios.ddns_principal", "nios.ddns_protected", "nios.disable", "nios.forbid_reclamation", "nios.name", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      canonical = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view      = infoblox_zone_auth.test.nios.view
      ext_attrs = { Site = "{{random4}}" }
    }
  }

}
