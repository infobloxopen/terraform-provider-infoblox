# Auto-generated datasource acceptance-test cases for Ipv6fixedaddress.
case "filters" {
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_ipv6_network" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_view.test.id
    }
  }
  PREREQ

  filter {
    type   = "filters"
    values = {
      address  = "uddi.address"
      ip_space = "uddi.ip_space"
    }
  }

  pair_checks = ["uddi.address", "uddi.comment", "uddi.disable_dhcp", "uddi.header_option_filename", "uddi.header_option_server_address", "uddi.header_option_server_name", "uddi.hostname", "uddi.inheritance_parent", "uddi.ip_space", "uddi.match_type", "uddi.match_value", "uddi.name", "uddi.parent"]

  step {
    uddi {
      ip_space    = infoblox_view.test.id
      address     = "10.0.0.10"
      match_type  = "mac"
      match_value = "aa:aa:aa:aa:aa:aa"
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
  resource "infoblox_ipv6_network" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_view.test.id
    }
  }
  PREREQ

  filter {
    type   = "tag_filters"
    values = {
      tag1 = "uddi.tags.tag1"
    }
  }

  pair_checks = ["uddi.address", "uddi.comment", "uddi.disable_dhcp", "uddi.header_option_filename", "uddi.header_option_server_address", "uddi.header_option_server_name", "uddi.hostname", "uddi.inheritance_parent", "uddi.ip_space", "uddi.match_type", "uddi.match_value", "uddi.name", "uddi.parent"]

  step {
    uddi {
      ip_space    = infoblox_view.test.id
      address     = "10.0.0.10"
      match_type  = "mac"
      match_value = "aa:aa:aa:aa:aa:aa"
      tags        = { tag1 = "{{random}}" }
    }
  }

}
