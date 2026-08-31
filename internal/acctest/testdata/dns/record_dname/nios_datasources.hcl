# Auto-generated datasource acceptance-test cases for RecordDname.
case "filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random3}}.com"
    }
  }
  PREREQ

  filter {
    type   = "filters"
    values = {
      name   = "nios.name"
      target = "nios.target"
      view   = "nios.view"
    }
  }

  pair_checks = ["nios.comment", "nios.creator", "nios.ddns_principal", "nios.ddns_protected", "nios.disable", "nios.forbid_reclamation", "nios.name", "nios.target", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      name   = infoblox_zone_auth.test.nios.fqdn
      target = "{{random}}.com"
      view   = infoblox_zone_auth.test.nios.view
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random3}}.com"
    }
  }
  PREREQ

  filter {
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.comment", "nios.creator", "nios.ddns_principal", "nios.ddns_protected", "nios.disable", "nios.forbid_reclamation", "nios.name", "nios.target", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      name      = infoblox_zone_auth.test.nios.fqdn
      target    = "{{random}}.com"
      view      = infoblox_zone_auth.test.nios.view
      ext_attrs = { Site = "{{random}}" }
    }
  }

}
