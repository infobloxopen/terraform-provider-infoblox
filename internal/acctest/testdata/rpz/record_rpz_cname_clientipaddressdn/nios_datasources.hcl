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

  pair_checks = ["nios.canonical", "nios.comment", "nios.disable", "nios.name", "nios.rp_zone", "nios.view"]

  step {
    nios {
      name      = "{{random_cidr_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
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

  pair_checks = ["nios.canonical", "nios.comment", "nios.disable", "nios.name", "nios.rp_zone", "nios.view"]

  step {
    nios {
      name      = "{{random_cidr_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      ext_attrs = { Site = "{{random3}}" }
    }
  }

}
