# Auto-generated datasource acceptance-test cases for RecordPtr.
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
      ipv4addr = "nios.ipv4addr"
      ptrdname = "nios.ptrdname"
      view     = "nios.view"
    }
  }

  pair_checks = ["nios.comment", "nios.creator", "nios.ddns_principal", "nios.ddns_protected", "nios.disable", "nios.forbid_reclamation", "nios.ipv4addr", "nios.ipv6addr", "nios.name", "nios.ptrdname", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      ipv4addr = "192.168.104.40"
      ptrdname = "host.{{random}}.com"
      view     = "default"
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

  pair_checks = ["nios.comment", "nios.creator", "nios.ddns_principal", "nios.ddns_protected", "nios.disable", "nios.forbid_reclamation", "nios.ipv4addr", "nios.ipv6addr", "nios.name", "nios.ptrdname", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      ipv4addr  = "192.168.104.41"
      ptrdname  = "host.{{random}}.com"
      view      = "default"
      ext_attrs = { Site = "{{random2}}" }
    }
  }

}
