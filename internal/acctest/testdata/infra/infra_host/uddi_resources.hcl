case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      display_name = "{{random}}"
    }
    check = {
      "uddi.display_name"     = "{{random}}"
      "uddi.maintenance_mode" = "disabled"
    }
  }
}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true

  step {
    uddi {
      display_name = "{{random}}"
    }
  }
}

case "description" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      display_name = "{{random}}"
      description  = "some description"
    }
    check = {
      "uddi.description" = "some description"
    }
  }

  step {
    uddi {
      display_name = "{{random}}"
      description  = "some updated description"
    }
    check = {
      "uddi.description" = "some updated description"
    }
  }
}

case "ip_space" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }

  resource "infoblox_network_view" "test2" {
    uddi = {
      name = "{{random}}2"
    }
  }
  PREREQ

  step {
    uddi {
      display_name = "{{random}}"
      ip_space     = infoblox_network_view.test.id
    }
    check = {
      "uddi.ip_space" = infoblox_network_view.test.id
    }
  }

  step {
    uddi {
      display_name = "{{random}}"
      ip_space     = infoblox_network_view.test2.id
    }
    check = {
      "uddi.ip_space" = infoblox_network_view.test2.id
    }
  }
}

case "maintenance_mode" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      display_name     = "{{random}}"
      maintenance_mode = "enabled"
    }
    check = {
      "uddi.maintenance_mode" = "enabled"
    }
  }

  step {
    uddi {
      display_name     = "{{random}}"
      maintenance_mode = "disabled"
    }
    check = {
      "uddi.maintenance_mode" = "disabled"
    }
  }
}

case "serial_number" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      display_name  = "{{random}}"
      serial_number = "{{random_int}}"
      tags = {
          "host/serial_number" = "{{random_int}}"
      }
    }
    check = {
      "uddi.serial_number" = "{{random_int}}"
    }
  }

  step {
    uddi {
      display_name  = "{{random}}"
      serial_number = "{{random_int2}}"
      tags = {
                "host/serial_number" = "{{random_int2}}"
            }
    }
    check = {
      "uddi.serial_number" = "{{random_int2}}"
    }
  }
}

case "tags" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      display_name = "{{random}}"
      tags = {
        tag1 = "value1"
      }
    }
    check = {
      "uddi.tags.tag1" = "value1"
    }
  }

  step {
    uddi {
      display_name = "{{random}}"
      tags = {
        tag1 = "value2"
      }
    }
    check = {
      "uddi.tags.tag1" = "value2"
    }
  }
}
