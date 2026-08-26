# Auto-generated datasource acceptance-test cases for RecordAlias.
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

  pair_checks = ["nios.comment", "nios.creator", "nios.disable", "nios.name", "nios.target_name", "nios.target_type", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target_name = "server.example.com"
      target_type = "A"
      view        = infoblox_zone_auth.test.nios.view
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

  pair_checks = ["nios.comment", "nios.creator", "nios.disable", "nios.name", "nios.target_name", "nios.target_type", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target_name = "server.example.com"
      target_type = "A"
      view        = infoblox_zone_auth.test.nios.view
      ext_attrs   = { Site = "{{random3}}" }
    }
  }

}
