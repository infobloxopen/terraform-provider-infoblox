# Auto-generated datasource acceptance-test cases for RecordRpzPtr.
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
      ptrdname = "nios.ptrdname"
    }
  }

  pair_checks = ["nios.comment", "nios.disable", "nios.ipv4addr", "nios.ipv6addr", "nios.name", "nios.ptrdname", "nios.rp_zone", "nios.ttl", "nios.view"]

  step {
    nios {
      ptrdname = "{{random2}}.{{random}}.com"
      ipv4addr = "{{random_ip}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
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

  pair_checks = ["nios.comment", "nios.disable", "nios.ipv4addr", "nios.ipv6addr", "nios.name", "nios.ptrdname", "nios.rp_zone", "nios.ttl", "nios.view"]

  step {
    nios {
      ptrdname  = "{{random2}}.{{random}}.com"
      ipv4addr  = "{{random_ip}}"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      ext_attrs = { Site = "{{random3}}" }
    }
  }

}
