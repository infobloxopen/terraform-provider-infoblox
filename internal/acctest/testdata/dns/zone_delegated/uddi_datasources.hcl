# Auto-generated datasource acceptance-test cases for ZoneDelegated.
case "filters" {
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn         = "{{random2}}.com."
      primary_type = "cloud"
      view         = infoblox_view.test.id
    }
  }
  PREREQ

  filter {
    type   = "filters"
    values = {
      fqdn = "uddi.fqdn"
    }
  }

  pair_checks = ["uddi.comment", "uddi.compartment_id", "uddi.disabled", "uddi.fqdn", "uddi.parent", "uddi.view"]

  step {
    uddi {
      fqdn = "{{random}}.${infoblox_zone_auth.test.uddi.fqdn}"
      delegation_servers = [{ address = "12.0.0.0", fqdn = "ns1.com." }]
      view = infoblox_view.test.id
    }
  }

}

case "tag_filters" {
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn         = "{{random2}}.com."
      primary_type = "cloud"
      view         = infoblox_view.test.id
    }
  }
  PREREQ

  filter {
    type   = "tag_filters"
    values = {
      tag1 = "uddi.tags.tag1"
    }
  }

  pair_checks = ["uddi.comment", "uddi.compartment_id", "uddi.disabled", "uddi.forward_only", "uddi.fqdn", "uddi.parent", "uddi.view"]

  step {
    uddi {
      fqdn = "{{random}}.${infoblox_zone_auth.test.uddi.fqdn}"
      delegation_servers = [{ address = "12.0.0.0", fqdn = "ns1.com." }]
      tags = { tag1 = "{{random}}" }
      view = infoblox_view.test.id
    }
  }

}
