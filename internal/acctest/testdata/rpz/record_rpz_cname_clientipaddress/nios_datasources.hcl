# Auto-generated datasource acceptance-test cases for RecordRpzCnameClientipaddress.
case "filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "custom_view" {
    nios = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_zone_rp" "test_zone" {
    nios = {
      fqdn = "{{random}}.com"
      view = infoblox_view.custom_view.nios.name
    }
  }
  PREREQ

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.canonical", "nios.comment", "nios.disable", "nios.name", "nios.rp_zone", "nios.ttl", "nios.view"]

  step {
    nios {
      name      = "12.0.0.40.${infoblox_zone_rp.test_zone.nios.fqdn}"
      canonical = "rpz-passthru"
      view      = infoblox_view.custom_view.nios.name
      rp_zone   = infoblox_zone_rp.test_zone.nios.fqdn
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "custom_view" {
    nios = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_zone_rp" "test_zone" {
    nios = {
      fqdn = "{{random}}.com"
      view = infoblox_view.custom_view.nios.name
    }
  }
  PREREQ

  filter {
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.canonical", "nios.comment", "nios.disable", "nios.name", "nios.rp_zone", "nios.ttl", "nios.view"]

  step {
    nios {
      name      = "12.0.0.41.${infoblox_zone_rp.test_zone.nios.fqdn}"
      canonical = "rpz-passthru"
      rp_zone   = infoblox_zone_rp.test_zone.nios.fqdn
      view      = infoblox_view.custom_view.nios.name
      ext_attrs = { Site = "{{random2}}" }
    }
  }

}
