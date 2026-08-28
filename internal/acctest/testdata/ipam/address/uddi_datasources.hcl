# Auto-generated datasource acceptance-test cases for Address.
case "filters" {
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_network" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr    = 24
      space   = infoblox_network_view.test.id
    }
  }
  PREREQ

  filter {
    type = "filters"
    values = {
      address = "uddi.address"
      space   = "uddi.space"
    }
  }

  pair_checks = ["uddi.address", "uddi.comment", "uddi.hwaddr", "uddi.interface", "uddi.space"]

  step {
    uddi {
      address   = "10.0.0.1"
      space     = infoblox_network_view.test.id
      comment   = "{{random}}"
      hwaddr    = "00:11:22:33:44:55"
      interface = "eth0"
    }
    depends_on = [infoblox_network.test]
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
  resource "infoblox_network" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr    = 24
      space   = infoblox_network_view.test.id
    }
  }
  PREREQ

  filter {
    type = "tag_filters"
    values = {
      tag1 = "uddi.tags.tag1"
    }
  }

  pair_checks = ["uddi.address", "uddi.comment", "uddi.hwaddr", "uddi.interface", "uddi.space"]

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_network_view.test.id
      tags    = { tag1 = "{{random}}" }
    }
    depends_on = [infoblox_network.test]
  }

}
