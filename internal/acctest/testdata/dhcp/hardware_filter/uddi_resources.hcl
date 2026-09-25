# HardwareFilter — uddi resource cases
# TODO: The following prerequisites MUST exist on the grid before running these tests:
#   - dhcp/option_code : dhcp/option_code/de50b0db-01cc-4da8-8213-aefd0880340f
#   - dhcp/option_space : dhcp/option_space/ae933dff-f5ff-415e-8e94-066e9c235295  (example_option_space_2 (custom option space))
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
      comment = "Hardware filter was created with Terraform"
    }
    check = {
      "uddi.comment" = "Hardware filter was created with Terraform"
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

}

case "dhcp_options" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
      dhcp_options = [{
        type         = "option"
        option_code  = "dhcp/option_code/de50b0db-01cc-4da8-8213-aefd0880340f"
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
        option_code  = "dhcp/option_code/de50b0db-01cc-4da8-8213-aefd0880340f"
        option_value = "value2"
      }]
    }
    check = {
      "uddi.dhcp_options.0.type"         = "option"
      "uddi.dhcp_options.0.option_value" = "value2"
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
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                                = "{{random}}"
      vendor_specific_option_option_space = "dhcp/option_space/ae933dff-f5ff-415e-8e94-066e9c235295"
    }
    check = {
      "uddi.vendor_specific_option_option_space" = "dhcp/option_space/ae933dff-f5ff-415e-8e94-066e9c235295"
    }
  }

}
