# Ipv6fixedaddress — nios list cases

case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        ipv6addr = "nios.ipv6addr"
      }
    }
  }

}
