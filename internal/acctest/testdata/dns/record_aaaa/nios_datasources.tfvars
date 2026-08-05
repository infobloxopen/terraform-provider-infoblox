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
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

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
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  step {
    nios {
      name      = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr  = "{{random_ipv6}}"
      view      = infoblox_zone_auth.test.nios.view
      ext_attrs = { Site = "{{random2}}" }
    }
  }

}
