# Auto-generated datasource acceptance-test cases for RecordCaa.
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

  pair_checks = ["nios.ca_flag", "nios.ca_tag", "nios.ca_value", "nios.comment", "nios.creator", "nios.ddns_principal", "nios.ddns_protected", "nios.disable", "nios.forbid_reclamation", "nios.name", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag  = 0
      ca_tag   = "issue"
      ca_value = "digicert.com"
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

  pair_checks = ["nios.ca_flag", "nios.ca_tag", "nios.ca_value", "nios.comment", "nios.creator", "nios.ddns_principal", "nios.ddns_protected", "nios.disable", "nios.forbid_reclamation", "nios.name", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag   = 0
      ca_tag    = "issue"
      ca_value  = "digicert.com"
      ext_attrs = { Site = "{{random}}" }
    }
  }

}
