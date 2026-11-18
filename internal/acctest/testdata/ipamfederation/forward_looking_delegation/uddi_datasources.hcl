# ForwardLookingDelegation — uddi datasource cases
case "filters" {
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_federated_realm" "test" {
    uddi = {
      name = "{{random}}-realm"
    }
  }
  PREREQ

  filter {
    type = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.address", "uddi.cidr", "uddi.comment", "uddi.name", "uddi.protocol", "uddi.network_compliant"]

  step {
    uddi {
      address          = "{{random_ipv4_network}}"
      cidr             = 16
      federated_realms = [infoblox_federated_realm.test.id]
      name             = "{{random}}"
    }
  }
}

case "tag_filters" {
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_federated_realm" "test" {
    uddi = {
      name = "{{random}}-realm"
    }
  }
  PREREQ

  filter {
    type = "tag_filters"
    values = {
      tag1 = "uddi.tags.tag1"
    }
  }

  pair_checks = ["uddi.address", "uddi.cidr", "uddi.comment", "uddi.name"]

  step {
    uddi {
      address          = "{{random_ipv4_network}}"
      cidr             = 16
      federated_realms = [infoblox_federated_realm.test.id]
      tags             = { tag1 = "{{random2}}" }
    }
  }
}

case "ipv6" {
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_federated_realm" "test" {
    uddi = {
      name = "{{random}}-realm"
    }
  }
  PREREQ

  filter {
    type = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.address", "uddi.cidr", "uddi.name"]

  step {
    uddi {
      address          = "2001:db8:{{random_hextet}}::"
      cidr             = 48
      federated_realms = [infoblox_federated_realm.test.id]
      name             = "{{random}}"
    }
  }
}
