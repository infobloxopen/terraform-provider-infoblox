case "filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test_zone" {
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

  pair_checks = ["nios.comment", "nios.disable", "nios.ipv4addr", "nios.name", "nios.rp_zone", "nios.ttl", "nios.view"]

  step {
    nios {
      name     = "10.10.0.0/16.${infoblox_zone_rp.test_zone.nios.fqdn}"
      ipv4addr = "10.10.0.1"
      rp_zone  = infoblox_zone_rp.test_zone.nios.fqdn
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test_zone" {
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

  pair_checks = ["nios.comment", "nios.disable", "nios.ipv4addr", "nios.name", "nios.rp_zone", "nios.ttl", "nios.view"]

  step {
    nios {
      name      = "10.10.0.0/16.${infoblox_zone_rp.test_zone.nios.fqdn}"
      ipv4addr  = "10.10.0.1"
      rp_zone   = infoblox_zone_rp.test_zone.nios.fqdn
      ext_attrs = { Site = "{{random2}}" }
    }
  }

}
