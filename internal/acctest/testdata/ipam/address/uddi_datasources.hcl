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

  pair_checks = ["id", "uddi.address", "uddi.comment", "uddi.external_keys", "uddi.hwaddr", "uddi.interface", "uddi.names", "uddi.space", "uddi.tags"]

  step {
    uddi {
      address       = "10.0.0.1"
      space         = infoblox_network_view.test.id
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

  pair_checks = ["id", "uddi.address", "uddi.comment", "uddi.external_keys", "uddi.hwaddr", "uddi.interface", "uddi.names", "uddi.space", "uddi.tags"]

  step {
    uddi {
      address       = "10.0.0.1"
      space         = infoblox_network_view.test.id
      tags          = { tag1 = "{{random}}" }
    }
    depends_on = [infoblox_network.test]
  }

}
