# Auto-generated datasource acceptance-test cases for ZoneDelegated.
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

  pair_checks = ["nios.comment", "nios.delegated_ttl", "nios.disable", "nios.enable_rfc2317_exclusion", "nios.fqdn", "nios.locked", "nios.ms_ad_integrated", "nios.ms_ddns_mode", "nios.ns_group", "nios.prefix", "nios.use_delegated_ttl", "nios.view", "nios.zone_format"]

  step {
    nios {
      fqdn = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to = [{ name = "{{random}}.com", address = "10.0.0.1" }]
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

  pair_checks = ["nios.comment", "nios.delegated_ttl", "nios.disable", "nios.enable_rfc2317_exclusion", "nios.fqdn", "nios.locked", "nios.ms_ad_integrated", "nios.ms_ddns_mode", "nios.ns_group", "nios.prefix", "nios.use_delegated_ttl", "nios.view", "nios.zone_format"]

  step {
    nios {
      fqdn      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to = [{ name = "{{random}}.com", address = "10.0.0.1" }]
      ext_attrs = { Site = "{{random3}}" }
    }
  }

}
