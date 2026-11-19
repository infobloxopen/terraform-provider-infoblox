# OverlappingBlock — uddi resource cases
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
      cidr            = 24
      federated_realm = infoblox_federated_realm.test.id
    }
    check = {
      "uddi.cidr" = "24"
      "uddi.federated_realm" = infoblox_federated_realm.test.id
    }
  }

}

case "disappears" {
  backend           = "uddi"
  parallel          = true
  disappears            = true
  expect_non_empty_plan = true
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
      cidr            = 24
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
      cidr            = 24
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
      cidr            = 24
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
      cidr            = 24
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
      cidr            = 24
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
      cidr            = 24
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
      cidr            = 24
      federated_realm = infoblox_federated_realm.test.id
      tags            = { tag2 = "value2changed", tag3 = "value3" }
    }
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }

}

case "address" {
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
      cidr            = 24
      federated_realm = infoblox_federated_realm.test.id
    }
  }

  # address is immutable, so changing it replaces the resource
  step {
    uddi {
      address         = "{{random_ipv4_network2}}"
      cidr            = 24
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
      cidr               = 24
      federated_realm    = infoblox_federated_realm.test.id
      dynamic_allocation = { next_available_id = infoblox_federated_block.parent.id }
    }
    check = {
      "uddi.cidr" = "24"
    }
  }

}

case "dynamic_allocation_tags" {
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

  # The create-next-available endpoint rejects tags, so they are patched in after allocation.
  step {
    uddi {
      cidr               = 24
      federated_realm    = infoblox_federated_realm.test.id
      name               = "{{random}}"
      tags               = { tag1 = "value1" }
      dynamic_allocation = { next_available_id = infoblox_federated_block.parent.id }
    }
    check = {
      "uddi.name"      = "{{random}}"
      "uddi.tags.tag1" = "value1"
    }
  }

}
