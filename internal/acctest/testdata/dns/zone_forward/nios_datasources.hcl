# Auto-generated datasource acceptance-test cases for ZoneForward.
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
      fqdn = "nios.fqdn"
    }
  }

  pair_checks = ["nios.comment", "nios.disable", "nios.disable_ns_generation", "nios.external_ns_group", "nios.forwarders_only", "nios.fqdn", "nios.locked", "nios.ms_ad_integrated", "nios.ms_ddns_mode", "nios.ns_group", "nios.prefix", "nios.view", "nios.zone_format"]

  step {
    nios {
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
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

  pair_checks = ["nios.comment", "nios.disable", "nios.disable_ns_generation", "nios.external_ns_group", "nios.forwarders_only", "nios.fqdn", "nios.locked", "nios.ms_ad_integrated", "nios.ms_ddns_mode", "nios.ns_group", "nios.prefix", "nios.view", "nios.zone_format"]

  step {
    nios {
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
      ext_attrs         = { Site = "{{random2}}" }
    }
  }

}
