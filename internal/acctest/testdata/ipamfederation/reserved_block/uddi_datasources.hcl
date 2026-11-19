# ReservedBlock — uddi datasource cases
case "filters" {
  backend           = "uddi"
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_federated_realm" "test" {
      uddi = {
        name = "{{random}}"
      }
    }
    PREREQ

  filter {
    type = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.comment", "uddi.name"]

  step {
    uddi {
      address         = "{{random_ipv4_network}}"
      cidr            = 24
      federated_realm = infoblox_federated_realm.test.id
      name            = "{{random}}"
    }
  }

}

case "tag_filters" {
  backend           = "uddi"
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_federated_realm" "test" {
      uddi = {
        name = "{{random}}"
      }
    }
    PREREQ

  filter {
    type = "tag_filters"
    values = {
      tag1 = "uddi.tags.tag1"
    }
  }

  pair_checks = ["uddi.comment", "uddi.name"]

  step {
    uddi {
      address         = "{{random_ipv4_network}}"
      cidr            = 24
      federated_realm = infoblox_federated_realm.test.id
      name            = "{{random}}"
      tags            = { tag1 = "{{random2}}" }
    }
  }

}
