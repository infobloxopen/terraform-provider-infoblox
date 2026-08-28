# Auto-generated datasource acceptance-test cases for Address.
case "filters" {
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_subnet" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_ip_space.test.id
    }
  }
  PREREQ

  filter {
    type   = "filters"
    values = {
      address = "uddi.address"
      space   = "uddi.space"
    }
  }

  pair_checks = ["uddi.address", "uddi.comment", "uddi.host", "uddi.hwaddr", "uddi.interface", "uddi.parent", "uddi.range", "uddi.space"]

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_ip_space.test.id
    }
  }

}

case "tag_filters" {
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_subnet" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_ip_space.test.id
    }
  }
  PREREQ

  filter {
    type   = "tag_filters"
    values = {
      tag1 = "uddi.tags.tag1"
    }
  }

  pair_checks = ["uddi.address", "uddi.comment", "uddi.host", "uddi.hwaddr", "uddi.interface", "uddi.parent", "uddi.range", "uddi.space"]

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_ip_space.test.id
      tags    = { tag1 = "{{random}}" }
    }
  }

}
