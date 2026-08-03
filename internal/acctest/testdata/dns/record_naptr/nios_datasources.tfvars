# Auto-generated datasource acceptance-test cases for RecordNaptr.
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

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
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

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
      ext_attrs   = { Site = "{{random}}" }
    }
  }

}
