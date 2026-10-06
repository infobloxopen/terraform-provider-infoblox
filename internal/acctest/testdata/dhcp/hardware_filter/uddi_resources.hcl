# HardwareFilter — uddi resource cases
case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.name" = "{{random}}"
      "uddi.role" = "values"
    }
  }

}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

  step {
    uddi {
      name = "{{random}}"
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "{{random}}"
      comment = "Hardware filter created with Terraform"
    }
    check = {
      "uddi.comment" = "Hardware filter created with Terraform"
    }
  }

  step {
    uddi {
      name    = "{{random}}"
      comment = "Hardware filter was updated with Terraform"
    }
    check = {
      "uddi.comment" = "Hardware filter was updated with Terraform"
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

  step {
    uddi {
      name = "{{random2}}"
    }
    check = {
      "uddi.name" = "{{random2}}"
    }
  }

}

case "addresses" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name      = "{{random}}"
      addresses = ["12:34:56:78:9a:bc"]
    }
    check = {
      "uddi.addresses.#" = "1"
      "uddi.addresses.0" = "12:34:56:78:9a:bc"
    }
  }

  step {
    uddi {
      name      = "{{random}}"
      addresses = ["12:34:56:78:9a:bc", "ab:cd:ef:12:34:56"]
    }
    check = {
      "uddi.addresses.#" = "2"
    }
  }

  step {
    uddi {
      name      = "{{random}}"
      addresses = ["ab:cd:ef:12:34:56", "12:34:56:78:9a:bc"]
    }
    check = {
      "uddi.addresses.#" = "2"
      "uddi.addresses.0" = "ab:cd:ef:12:34:56"
      "uddi.addresses.1" = "12:34:56:78:9a:bc"
    }
  }

}

case "dhcp_options" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_dhcp_optionspace" "test" {
      uddi = {
        name = "{{random}}"
      }
    }
    resource "infoblox_dhcp_optiondefinition" "test" {
      uddi = {
        code         = 150
        name         = "{{random}}"
        option_space = infoblox_dhcp_optionspace.test.id
        type         = "text"
      }
    }
    resource "infoblox_option_group" "test" {
      uddi = {
        name     = "{{random}}"
        protocol = "ip4"
      }
    }
    PREREQ

  step {
    uddi {
      name = "{{random}}"
      dhcp_options = [{
        type         = "option"
        option_code  = infoblox_dhcp_optiondefinition.test.id
        option_value = "value1"
      }]
    }
    check = {
      "uddi.dhcp_options.0.type"         = "option"
      "uddi.dhcp_options.0.option_value" = "value1"
    }
  }

  step {
    uddi {
      name = "{{random}}"
      dhcp_options = [{
        type         = "option"
        option_code  = infoblox_dhcp_optiondefinition.test.id
        option_value = "value2"
      }]
    }
    check = {
      "uddi.dhcp_options.0.type"         = "option"
      "uddi.dhcp_options.0.option_value" = "value2"
    }
  }

  step {
    uddi {
      name = "{{random}}"
      dhcp_options = [{
        type  = "group"
        group = infoblox_option_group.test.id
      }]
    }
    check = {
      "uddi.dhcp_options.0.type" = "group"
    }
  }

}

case "header_option_filename" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                   = "{{random}}"
      header_option_filename = "pxeboot.img"
    }
    check = {
      "uddi.header_option_filename" = "pxeboot.img"
    }
  }

  step {
    uddi {
      name                   = "{{random}}"
      header_option_filename = "pxeboot-update.img"
    }
    check = {
      "uddi.header_option_filename" = "pxeboot-update.img"
    }
  }

}

case "header_option_server_address" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                         = "{{random}}"
      header_option_server_address = "192.168.10.10"
    }
    check = {
      "uddi.header_option_server_address" = "192.168.10.10"
    }
  }

  step {
    uddi {
      name                         = "{{random}}"
      header_option_server_address = "192.168.11.11"
    }
    check = {
      "uddi.header_option_server_address" = "192.168.11.11"
    }
  }

}

case "header_option_server_name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                      = "{{random}}"
      header_option_server_name = "tf-infoblox-test.com."
    }
    check = {
      "uddi.header_option_server_name" = "tf-infoblox-test.com."
    }
  }

  step {
    uddi {
      name                      = "{{random}}"
      header_option_server_name = "tf-infoblox.com."
    }
    check = {
      "uddi.header_option_server_name" = "tf-infoblox.com."
    }
  }

}

case "lease_time" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "{{random}}"
      lease_time = 600
    }
    check = {
      "uddi.lease_time" = "600"
    }
  }

  step {
    uddi {
      name       = "{{random}}"
      lease_time = 1200
    }
    check = {
      "uddi.lease_time" = "1200"
    }
  }

}

case "role" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
      role = "values"
    }
    check = {
      "uddi.role" = "values"
    }
  }

  step {
    uddi {
      name = "{{random}}"
      role = "selection"
    }
    check = {
      "uddi.role" = "selection"
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
      tags = { tag1 = "value1", tag2 = "value2" }
    }
    check = {
      "uddi.tags.tag1" = "value1"
      "uddi.tags.tag2" = "value2"
    }
  }

  step {
    uddi {
      name = "{{random}}"
      tags = { tag2 = "value2changed", tag3 = "value3" }
    }
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }

}

case "vendor_specific_option_option_space" {
  backend           = "uddi"
  parallel          = true
  skip_if_env_empty = "INFOBLOX_ACC_DHCP_OPTION_SPACE"

  step {
    uddi {
      name                                = "{{random}}"
      vendor_specific_option_option_space = "{{dhcp_option_space}}"
    }
    check = {
      "uddi.vendor_specific_option_option_space" = "{{dhcp_option_space}}"
    }
  }

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.vendor_specific_option_option_space" = "{{dhcp_option_space}}"
    }
  }

}

