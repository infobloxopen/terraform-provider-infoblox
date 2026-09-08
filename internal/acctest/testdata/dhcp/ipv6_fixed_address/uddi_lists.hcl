# Ipv6fixedaddress — uddi list cases
case "basic" {
  backend        = "uddi"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_ipv6_network" "test" {
    uddi = {
      address = "2001:db8:{{random_hextet}}:{{random_int}}::"
      cidr = 64
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = replace(infoblox_ipv6_network.test.uddi.address, "::", "::10")
      match_type  = "mac"
      match_value = "aa:aa:aa:aa:aa:aa"
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend        = "uddi"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_ipv6_network" "test" {
    uddi = {
      address = "2001:db8:{{random_hextet}}:{{random_int}}::"
      cidr = 64
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = replace(infoblox_ipv6_network.test.uddi.address, "::", "::10")
      match_type  = "mac"
      match_value = "aa:aa:aa:aa:aa:aa"
      name        = "{{random2}}"
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        name = "uddi.name"
      }
    }
  }

}
