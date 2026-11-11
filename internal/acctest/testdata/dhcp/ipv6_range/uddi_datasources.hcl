# Auto-generated datasource acceptance-test cases for Ipv6range.
case "filters" {
  backend = "uddi"
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

  filter {
    type   = "filters"
    values = {
      start = "uddi.start"
      end   = "uddi.end"
      space = "uddi.space"
    }
  }

  pair_checks = ["uddi.comment", "uddi.dhcp_host", "uddi.disable_dhcp", "uddi.end", "uddi.name", "uddi.space", "uddi.start"]

  step {
    uddi {
      end   = "2001:db8:{{random_hextet}}:{{random_int}}::20"
      space = infoblox_network_view.test.id
      start = "2001:db8:{{random_hextet}}:{{random_int}}::8"
    }
    depends_on = [infoblox_ipv6_network.test]
  }

}

case "tag_filters" {
  backend = "uddi"
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

  filter {
    type   = "tag_filters"
    values = {
      tag1 = "uddi.tags.tag1"
    }
  }

  pair_checks = ["uddi.comment", "uddi.dhcp_host", "uddi.disable_dhcp", "uddi.end", "uddi.name", "uddi.space", "uddi.start"]

  step {
    uddi {
      end   = "2001:db8:{{random_hextet}}:{{random_int}}::20"
      space = infoblox_network_view.test.id
      start = "2001:db8:{{random_hextet}}:{{random_int}}::8"
      tags  = { tag1 = "{{random}}" }
    }
    depends_on = [infoblox_ipv6_network.test]
  }

}
