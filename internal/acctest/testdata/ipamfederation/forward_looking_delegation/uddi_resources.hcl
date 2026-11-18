# ForwardLookingDelegation — uddi resource cases
case "basic" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_federated_realm" "test" {
    uddi = {
      name = "{{random}}-realm"
    }
  }
  PREREQ

  step {
    uddi {
      address          = "{{random_ipv4_network}}"
      cidr             = 16
      federated_realms = [infoblox_federated_realm.test.id]
    }
    check = {
      "uddi.cidr"               = "16"
      "uddi.federated_realms.#" = "1"
      "uddi.comment"            = ""
      "uddi.protocol"           = "ip4"
    }
    check_pair = {
      "uddi.federated_realms.0" = infoblox_federated_realm.test.id
    }
  }
}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_federated_realm" "test" {
    uddi = {
      name = "{{random}}-realm"
    }
  }
  PREREQ

  step {
    uddi {
      address          = "{{random_ipv4_network}}"
      cidr             = 16
      federated_realms = [infoblox_federated_realm.test.id]
    }
  }
}

case "ipv6" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_federated_realm" "test" {
    uddi = {
      name = "{{random}}-realm"
    }
  }
  PREREQ

  step {
    uddi {
      address          = "2001:db8:{{random_hextet}}::"
      cidr             = 48
      federated_realms = [infoblox_federated_realm.test.id]
    }
    check = {
      "uddi.cidr" = "48"
    }
  }
}

# Changing address forces a replacement of the delegation.
case "address" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_federated_realm" "test" {
    uddi = {
      name = "{{random}}-realm"
    }
  }
  PREREQ

  step {
    uddi {
      address          = "10.{{random_octet}}.0.0"
      cidr             = 16
      federated_realms = [infoblox_federated_realm.test.id]
    }
    check = {
      "uddi.cidr" = "16"
    }
  }

  step {
    uddi {
      address          = "11.{{random_octet}}.0.0"
      cidr             = 16
      federated_realms = [infoblox_federated_realm.test.id]
    }
    check = {
      "uddi.cidr" = "16"
    }
  }
}

case "cidr" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_federated_realm" "test" {
    uddi = {
      name = "{{random}}-realm"
    }
  }
  PREREQ

  step {
    uddi {
      address          = "10.{{random_octet}}.0.0"
      cidr             = 16
      federated_realms = [infoblox_federated_realm.test.id]
    }
    check = {
      "uddi.cidr" = "16"
    }
  }

  step {
    uddi {
      address          = "10.{{random_octet}}.0.0"
      cidr             = 17
      federated_realms = [infoblox_federated_realm.test.id]
    }
    check = {
      "uddi.cidr" = "17"
    }
  }
}

case "comment" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_federated_realm" "test" {
    uddi = {
      name = "{{random}}-realm"
    }
  }
  PREREQ

  step {
    uddi {
      address          = "{{random_ipv4_network}}"
      cidr             = 16
      federated_realms = [infoblox_federated_realm.test.id]
      comment          = "test comment"
    }
    check = {
      "uddi.comment" = "test comment"
    }
  }

  step {
    uddi {
      address          = "{{random_ipv4_network}}"
      cidr             = 16
      federated_realms = [infoblox_federated_realm.test.id]
      comment          = "test comment updated"
    }
    check = {
      "uddi.comment" = "test comment updated"
    }
  }
}

case "name" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_federated_realm" "test" {
    uddi = {
      name = "{{random}}-realm"
    }
  }
  PREREQ

  step {
    uddi {
      address          = "{{random_ipv4_network}}"
      cidr             = 16
      federated_realms = [infoblox_federated_realm.test.id]
      name             = "{{random}}"
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

  step {
    uddi {
      address          = "{{random_ipv4_network}}"
      cidr             = 16
      federated_realms = [infoblox_federated_realm.test.id]
      name             = "{{random2}}"
    }
    check = {
      "uddi.name" = "{{random2}}"
    }
  }
}

case "federated_realms" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_federated_realm" "one" {
    uddi = {
      name = "{{random}}-one"
    }
  }

  resource "infoblox_federated_realm" "two" {
    uddi = {
      name = "{{random}}-two"
    }
  }
  PREREQ

  step {
    uddi {
      address          = "10.{{random_octet}}.0.0"
      cidr             = 16
      federated_realms = [infoblox_federated_realm.one.id]
    }
    check = {
      "uddi.federated_realms.#" = "1"
    }
  }

  step {
    uddi {
      address          = "10.{{random_octet}}.0.0"
      cidr             = 16
      federated_realms = [infoblox_federated_realm.one.id, infoblox_federated_realm.two.id]
    }
    check = {
      "uddi.federated_realms.#" = "2"
    }
  }

  # Order of realms is not significant (unordered list).
  step {
    uddi {
      address          = "10.{{random_octet}}.0.0"
      cidr             = 16
      federated_realms = [infoblox_federated_realm.two.id, infoblox_federated_realm.one.id]
    }
    check = {
      "uddi.federated_realms.#" = "2"
    }
  }
}

case "tags" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_federated_realm" "test" {
    uddi = {
      name = "{{random}}-realm"
    }
  }
  PREREQ

  step {
    uddi {
      address          = "{{random_ipv4_network}}"
      cidr             = 16
      federated_realms = [infoblox_federated_realm.test.id]
      tags             = { tag1 = "value1", tag2 = "value2" }
    }
    check = {
      "uddi.tags.tag1" = "value1"
      "uddi.tags.tag2" = "value2"
    }
  }

  step {
    uddi {
      address          = "{{random_ipv4_network}}"
      cidr             = 16
      federated_realms = [infoblox_federated_realm.test.id]
      tags             = { tag2 = "value2changed", tag3 = "value3" }
    }
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }
}
