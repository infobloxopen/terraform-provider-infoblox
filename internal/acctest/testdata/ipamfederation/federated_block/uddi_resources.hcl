# FederatedBlock — uddi resource cases
case "basic" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_federated_realm" "test" {
      uddi = {
        name = "{{random}}"
      }
    }
    PREREQ

  step {
    uddi {
      address         = "{{random_ipv4_network}}"
      cidr            = 16
      federated_realm = infoblox_federated_realm.test.id
    }
    check = {
      "uddi.cidr"            = "16"
      "uddi.federated_realm" = infoblox_federated_realm.test.id
    }
  }

}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl     = <<-PREREQ
    resource "infoblox_federated_realm" "test" {
      uddi = {
        name = "{{random}}"
      }
    }
    PREREQ

  step {
    uddi {
      address         = "{{random_ipv4_network}}"
      cidr            = 16
      federated_realm = infoblox_federated_realm.test.id
    }
  }

}

case "comment" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_federated_realm" "test" {
      uddi = {
        name = "{{random}}"
      }
    }
    PREREQ

  step {
    uddi {
      address         = "{{random_ipv4_network}}"
      cidr            = 16
      federated_realm = infoblox_federated_realm.test.id
      comment         = "test comment"
    }
    check = {
      "uddi.comment" = "test comment"
    }
  }

  step {
    uddi {
      address         = "{{random_ipv4_network}}"
      cidr            = 16
      federated_realm = infoblox_federated_realm.test.id
      comment         = "test comment update"
    }
    check = {
      "uddi.comment" = "test comment update"
    }
  }

}

case "name" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_federated_realm" "test" {
      uddi = {
        name = "{{random}}"
      }
    }
    PREREQ

  step {
    uddi {
      address         = "{{random_ipv4_network}}"
      cidr            = 16
      federated_realm = infoblox_federated_realm.test.id
      name            = "{{random}}"
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

  step {
    uddi {
      address         = "{{random_ipv4_network}}"
      cidr            = 16
      federated_realm = infoblox_federated_realm.test.id
      name            = "{{random2}}"
    }
    check = {
      "uddi.name" = "{{random2}}"
    }
  }

}

case "tags" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_federated_realm" "test" {
      uddi = {
        name = "{{random}}"
      }
    }
    PREREQ

  step {
    uddi {
      address         = "{{random_ipv4_network}}"
      cidr            = 16
      federated_realm = infoblox_federated_realm.test.id
      tags            = { tag1 = "value1", tag2 = "value2" }
    }
    check = {
      "uddi.tags.tag1" = "value1"
      "uddi.tags.tag2" = "value2"
    }
  }

  step {
    uddi {
      address         = "{{random_ipv4_network}}"
      cidr            = 16
      federated_realm = infoblox_federated_realm.test.id
      tags            = { tag2 = "value2changed", tag3 = "value3" }
    }
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }

}

case "network_compliance" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_federated_realm" "test" {
      uddi = {
        name = "{{random}}"
      }
    }
    PREREQ

  step {
    uddi {
      address         = "{{random_ipv4_network}}"
      cidr            = 16
      federated_realm = infoblox_federated_realm.test.id
      network_compliance = {
        minimum_netmask_length = 20
        maximum_netmask_length = 28
        default_netmask_length = 24
      }
    }
    check = {
      "uddi.network_compliance.minimum_netmask_length" = "20"
      "uddi.network_compliance.maximum_netmask_length" = "28"
      "uddi.network_compliance.default_netmask_length" = "24"
    }
  }

  step {
    uddi {
      address         = "{{random_ipv4_network}}"
      cidr            = 16
      federated_realm = infoblox_federated_realm.test.id
      network_compliance = {
        minimum_netmask_length = 20
        maximum_netmask_length = 28
      }
    }
    check = {
      "uddi.network_compliance.minimum_netmask_length" = "20"
      "uddi.network_compliance.maximum_netmask_length" = "28"
    }
  }

  step {
    uddi {
      address         = "{{random_ipv4_network}}"
      cidr            = 16
      federated_realm = infoblox_federated_realm.test.id
    }
  }

}

case "dynamic_allocation" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_federated_realm" "test" {
      uddi = {
        name = "{{random}}"
      }
    }
    resource "infoblox_federated_block" "parent" {
      uddi = {
        address         = "{{random_ipv4_network}}"
        cidr            = 16
        federated_realm = infoblox_federated_realm.test.id
      }
    }
    PREREQ

  step {
    uddi {
      cidr               = 20
      federated_realm    = infoblox_federated_realm.test.id
      dynamic_allocation = { next_available_id = infoblox_federated_block.parent.id }
    }
    check = {
      "uddi.cidr" = "20"
    }
  }

}
