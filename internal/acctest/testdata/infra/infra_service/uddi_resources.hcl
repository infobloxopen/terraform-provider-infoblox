case "basic" {
  backend           = "uddi"
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_infra_host" "test" {
      uddi = {
        display_name  = "{{random}}"
        serial_number = "{{random_int}}"
        tags          = { "host/serial_number" = "{{random_int}}" }
      }
    }
    PREREQ

  step {
    uddi {
      name         = "{{random}}"
      pool_id      = infoblox_infra_host.test.uddi.pool_id
      service_type = "dns"
    }
    check = {
      "uddi.name"          = "{{random}}"
      "uddi.pool_id"       = infoblox_infra_host.test.uddi.pool_id
      "uddi.service_type"  = "dns"
      "uddi.desired_state" = "stop"
    }
  }

}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true
  prerequisites_hcl     = <<-PREREQ
    resource "infoblox_infra_host" "test" {
      uddi = {
        display_name  = "{{random}}"
        serial_number = "{{random_int}}"
        tags          = { "host/serial_number" = "{{random_int}}" }
      }
    }
    PREREQ

  step {
    uddi {
      name         = "{{random}}"
      pool_id      = infoblox_infra_host.test.uddi.pool_id
      service_type = "dns"
    }
  }

}

case "description" {
  backend           = "uddi"
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_infra_host" "test" {
      uddi = {
        display_name  = "{{random}}"
        serial_number = "{{random_int}}"
        tags          = { "host/serial_number" = "{{random_int}}" }
      }
    }
    PREREQ

  step {
    uddi {
      name         = "{{random}}"
      pool_id      = infoblox_infra_host.test.uddi.pool_id
      service_type = "dns"
      description  = "initial description"
    }
    check = {
      "uddi.description" = "initial description"
    }
  }

  step {
    uddi {
      name         = "{{random}}"
      pool_id      = infoblox_infra_host.test.uddi.pool_id
      service_type = "dns"
      description  = "updated description"
    }
    check = {
      "uddi.description" = "updated description"
    }
  }

  step {
    uddi {
      name         = "{{random}}"
      pool_id      = infoblox_infra_host.test.uddi.pool_id
      service_type = "dns"
    }
  }

}

case "desired_state" {
  backend           = "uddi"
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_infra_host" "test" {
      uddi = {
        display_name  = "{{random}}"
        serial_number = "{{random_int}}"
        tags          = { "host/serial_number" = "{{random_int}}" }
      }
    }
    PREREQ

  step {
    uddi {
      name          = "{{random}}"
      pool_id       = infoblox_infra_host.test.uddi.pool_id
      service_type  = "dns"
      desired_state = "stop"
    }
    check = {
      "uddi.desired_state" = "stop"
    }
  }

  step {
    uddi {
      name          = "{{random}}"
      pool_id       = infoblox_infra_host.test.uddi.pool_id
      service_type  = "dns"
      desired_state = "start"
    }
    check = {
      "uddi.desired_state" = "start"
    }
  }

}

case "desired_version" {
  backend           = "uddi"
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_infra_host" "test" {
      uddi = {
        display_name  = "{{random}}"
        serial_number = "{{random_int}}"
        tags          = { "host/serial_number" = "{{random_int}}" }
      }
    }
    PREREQ

  step {
    uddi {
      name            = "{{random}}"
      pool_id         = infoblox_infra_host.test.uddi.pool_id
      service_type    = "dns"
      desired_version = "3.5.0"
    }
    check = {
      "uddi.desired_version" = "3.5.0"
    }
  }

  step {
    uddi {
      name            = "{{random}}"
      pool_id         = infoblox_infra_host.test.uddi.pool_id
      service_type    = "dns"
      desired_version = "3.6.0"
    }
    check = {
      "uddi.desired_version" = "3.6.0"
    }
  }

}

case "interface_labels" {
  backend           = "uddi"
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_infra_host" "test" {
      uddi = {
        display_name  = "{{random}}"
        serial_number = "{{random_int}}"
        tags          = { "host/serial_number" = "{{random_int}}" }
      }
    }
    PREREQ

  step {
    uddi {
      name             = "{{random}}"
      pool_id          = infoblox_infra_host.test.uddi.pool_id
      service_type     = "dns"
      interface_labels = ["WAN", "LAN"]
    }
    check = {
      "uddi.interface_labels.0" = "WAN"
      "uddi.interface_labels.1" = "LAN"
    }
  }

  step {
    uddi {
      name             = "{{random}}"
      pool_id          = infoblox_infra_host.test.uddi.pool_id
      service_type     = "dns"
      interface_labels = ["WAN", "LAN", "MGMT"]
    }
    check = {
      "uddi.interface_labels.0" = "WAN"
      "uddi.interface_labels.1" = "LAN"
      "uddi.interface_labels.2" = "MGMT"
    }
  }

}

case "service_type_anycast" {
  backend           = "uddi"
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_infra_host" "test" {
      uddi = {
        display_name  = "{{random}}"
        serial_number = "{{random_int}}"
        tags          = { "host/serial_number" = "{{random_int}}" }
      }
    }
    PREREQ

  step {
    uddi {
      name         = "{{random}}"
      pool_id      = infoblox_infra_host.test.uddi.pool_id
      service_type = "anycast"
    }
    check = {
      "uddi.service_type" = "anycast"
    }
  }

}

case "service_type_dhcp" {
  backend           = "uddi"
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_infra_host" "test" {
      uddi = {
        display_name  = "{{random}}"
        serial_number = "{{random_int}}"
        tags          = { "host/serial_number" = "{{random_int}}" }
      }
    }
    PREREQ

  step {
    uddi {
      name         = "{{random}}"
      pool_id      = infoblox_infra_host.test.uddi.pool_id
      service_type = "dhcp"
    }
    check = {
      "uddi.service_type" = "dhcp"
    }
  }

}

case "service_type_dfp" {
  backend           = "uddi"
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_infra_host" "test" {
      uddi = {
        display_name  = "{{random}}"
        serial_number = "{{random_int}}"
        tags          = { "host/serial_number" = "{{random_int}}" }
      }
    }
    PREREQ

  step {
    uddi {
      name         = "{{random}}"
      pool_id      = infoblox_infra_host.test.uddi.pool_id
      service_type = "dfp"
    }
    check = {
      "uddi.service_type" = "dfp"
    }
  }

}

case "service_type_ntp" {
  backend           = "uddi"
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_infra_host" "test" {
      uddi = {
        display_name  = "{{random}}"
        serial_number = "{{random_int}}"
        tags          = { "host/serial_number" = "{{random_int}}" }
      }
    }
    PREREQ

  step {
    uddi {
      name         = "{{random}}"
      pool_id      = infoblox_infra_host.test.uddi.pool_id
      service_type = "ntp"
    }
    check = {
      "uddi.service_type" = "ntp"
    }
  }

}

case "service_type_cdc" {
  backend           = "uddi"
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_infra_host" "test" {
      uddi = {
        display_name  = "{{random}}"
        serial_number = "{{random_int}}"
        tags          = { "host/serial_number" = "{{random_int}}" }
      }
    }
    PREREQ

  step {
    uddi {
      name         = "{{random}}"
      pool_id      = infoblox_infra_host.test.uddi.pool_id
      service_type = "cdc"
    }
    check = {
      "uddi.service_type" = "cdc"
    }
  }

}

case "service_type_authn" {
  backend           = "uddi"
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_infra_host" "test" {
      uddi = {
        display_name  = "{{random}}"
        serial_number = "{{random_int}}"
        tags          = { "host/serial_number" = "{{random_int}}" }
      }
    }
    PREREQ

  step {
    uddi {
      name         = "{{random}}"
      pool_id      = infoblox_infra_host.test.uddi.pool_id
      service_type = "authn"
    }
    check = {
      "uddi.service_type" = "authn"
    }
  }

}

case "service_type_msad" {
  backend           = "uddi"
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_infra_host" "test" {
      uddi = {
        display_name  = "{{random}}"
        serial_number = "{{random_int}}"
        tags          = { "host/serial_number" = "{{random_int}}" }
      }
    }
    PREREQ

  step {
    uddi {
      name         = "{{random}}"
      pool_id      = infoblox_infra_host.test.uddi.pool_id
      service_type = "msad"
    }
    check = {
      "uddi.service_type" = "msad"
    }
  }

}

case "tags" {
  backend           = "uddi"
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_infra_host" "test" {
      uddi = {
        display_name  = "{{random}}"
        serial_number = "{{random_int}}"
        tags          = { "host/serial_number" = "{{random_int}}" }
      }
    }
    PREREQ

  step {
    uddi {
      name         = "{{random}}"
      pool_id      = infoblox_infra_host.test.uddi.pool_id
      service_type = "dns"
      tags         = { env = "{{random2}}" }
    }
    check = {
      "uddi.tags.env" = "{{random2}}"
    }
  }

  step {
    uddi {
      name         = "{{random}}"
      pool_id      = infoblox_infra_host.test.uddi.pool_id
      service_type = "dns"
      tags         = { env = "{{random3}}" }
    }
    check = {
      "uddi.tags.env" = "{{random3}}"
    }
  }

}
