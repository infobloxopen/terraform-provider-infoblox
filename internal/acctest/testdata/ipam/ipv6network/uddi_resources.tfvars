# Auto-generated resource acceptance-test cases for Ipv6network.
case "basic" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address = "{{random_ipv6}}"
      cidr    = 128
      # space   = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
    }
    check = {
      "uddi.address"                       = "{{random_ipv6}}"
      "uddi.cidr"                          = "128"
      "uddi.ddns_client_update"            = "client"
      "uddi.ddns_conflict_resolution_mode" = "check_with_dhcid"
      "uddi.ddns_generate_name"            = "false"
      "uddi.ddns_generated_prefix"         = "myhost"
      "uddi.ddns_send_updates"             = "true"
      "uddi.ddns_update_on_renew"          = "false"
      "uddi.ddns_use_conflict_resolution"  = "true"
      "uddi.hostname_rewrite_char"         = "-"
      "uddi.hostname_rewrite_enabled"      = "false"
      "uddi.hostname_rewrite_regex"        = "[^a-zA-Z0-9_.]"
    }
  }

}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address = "{{random_ipv6}}"
      cidr    = 128
      # space   = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
    }
  }

}

case "address" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address = "{{random_ipv6}}"
      cidr    = 128
      # space   = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
    }
    check = {
      "uddi.address" = "{{random_ipv6}}"
      "uddi.cidr"    = "128"
    }
  }

  step {
    uddi {
      address = "{{random_ipv6_2}}"
      cidr    = 128
      # space   = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
    }
    check = {
      "uddi.address" = "{{random_ipv6_2}}"
      "uddi.cidr"    = "128"
    }
  }

}

case "cidr" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address = "{{random_ipv6}}"
      cidr    = 128
      # space   = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
    }
    check = {
      "uddi.address" = "{{random_ipv6}}"
      "uddi.cidr"    = "128"
    }
  }

  step {
    uddi {
      address = "{{random_ipv6}}"
      cidr    = 128
      # space   = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
    }
    check = {
      "uddi.address" = "{{random_ipv6}}"
      "uddi.cidr"    = "128"
    }
  }

}

case "space" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "one" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address = "{{random_ipv6}}"
      cidr    = 128
      # space   = infoblox_ip_space.one.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
    }
  }

  step {
    uddi {
      address = "{{random_ipv6}}"
      cidr    = 128
      # space   = infoblox_ip_space.two.id
      space = "ipam/ip_space/1fcd4065-8847-11f1-b283-5eecb1762ec1"
    }
  }

}

case "asm_config" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address    = "{{random_ipv6}}"
      cidr       = 128
      # space      = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      asm_config = { asm_threshold = 70, enable = true, enable_notification = true, forecast_period = 12, growth_factor = 40, growth_type = "count", history = 40, min_total = 30, min_unused = 30, reenable_date = "2020-01-10T10:11:22Z" }
    }
    check = {
      "uddi.asm_config.asm_threshold"       = "70"
      "uddi.asm_config.enable"              = "true"
      "uddi.asm_config.enable_notification" = "true"
      "uddi.asm_config.forecast_period"     = "12"
      "uddi.asm_config.growth_factor"       = "40"
      "uddi.asm_config.growth_type"         = "count"
      "uddi.asm_config.history"             = "40"
      "uddi.asm_config.min_total"           = "30"
      "uddi.asm_config.min_unused"          = "30"
      "uddi.asm_config.reenable_date"       = "2020-01-10T10:11:22Z"
    }
  }

  step {
    uddi {
      address    = "{{random_ipv6}}"
      cidr       = 128
      # space      = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      asm_config = { asm_threshold = 80, enable = false, enable_notification = false, forecast_period = 10, growth_factor = 50, growth_type = "percent", history = 50, min_total = 10, min_unused = 10, reenable_date = "2021-01-10T10:11:22Z" }
    }
    check = {
      "uddi.asm_config.asm_threshold"       = "80"
      "uddi.asm_config.enable"              = "false"
      "uddi.asm_config.enable_notification" = "false"
      "uddi.asm_config.forecast_period"     = "10"
      "uddi.asm_config.growth_factor"       = "50"
      "uddi.asm_config.growth_type"         = "percent"
      "uddi.asm_config.history"             = "50"
      "uddi.asm_config.min_total"           = "10"
      "uddi.asm_config.min_unused"          = "10"
      "uddi.asm_config.reenable_date"       = "2021-01-10T10:11:22Z"
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address = "{{random_ipv6}}"
      cidr    = 128
      # space   = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      comment = "some comment"
    }
    check = {
      "uddi.comment" = "some comment"
    }
  }

  step {
    uddi {
      address = "{{random_ipv6}}"
      cidr    = 128
      # space   = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      comment = "updated comment"
    }
    check = {
      "uddi.comment" = "updated comment"
    }
  }

}

case "ddns_client_update" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address            = "{{random_ipv6}}"
      cidr               = 128
      # space              = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      ddns_client_update = "server"
    }
    check = {
      "uddi.ddns_client_update" = "server"
    }
  }

  step {
    uddi {
      address            = "{{random_ipv6}}"
      cidr               = 128
      # space              = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      ddns_client_update = "over_client_update"
    }
    check = {
      "uddi.ddns_client_update" = "over_client_update"
    }
  }

}

case "ddns_conflict_resolution_mode" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address                       = "{{random_ipv6}}"
      cidr                          = 128
      # space                         = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      ddns_use_conflict_resolution  = false
      ddns_conflict_resolution_mode = "check_exists_with_dhcid"
    }
    check = {
      "uddi.ddns_use_conflict_resolution"  = "false"
      "uddi.ddns_conflict_resolution_mode" = "check_exists_with_dhcid"
    }
  }

  step {
    uddi {
      address                       = "{{random_ipv6}}"
      cidr                          = 128
      # space                         = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      ddns_use_conflict_resolution  = true
      ddns_conflict_resolution_mode = "check_with_dhcid"
    }
    check = {
      "uddi.ddns_use_conflict_resolution"  = "true"
      "uddi.ddns_conflict_resolution_mode" = "check_with_dhcid"
    }
  }

}

case "ddns_domain" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address     = "{{random_ipv6}}"
      cidr        = 128
      # space       = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      ddns_domain = "abc"
    }
    check = {
      "uddi.ddns_domain" = "abc"
    }
  }

  step {
    uddi {
      address     = "{{random_ipv6}}"
      cidr        = 128
      # space       = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      ddns_domain = "xyz"
    }
    check = {
      "uddi.ddns_domain" = "xyz"
    }
  }

}

case "ddns_generate_name" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address            = "{{random_ipv6}}"
      cidr               = 128
      # space              = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      ddns_generate_name = true
    }
    check = {
      "uddi.ddns_generate_name" = "true"
    }
  }

  step {
    uddi {
      address            = "{{random_ipv6}}"
      cidr               = 128
      # space              = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      ddns_generate_name = false
    }
    check = {
      "uddi.ddns_generate_name" = "false"
    }
  }

}

case "ddns_generated_prefix" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address               = "{{random_ipv6}}"
      cidr                  = 128
      # space                 = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      ddns_generated_prefix = "host-prefix"
    }
    check = {
      "uddi.ddns_generated_prefix" = "host-prefix"
    }
  }

  step {
    uddi {
      address               = "{{random_ipv6}}"
      cidr                  = 128
      # space                 = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      ddns_generated_prefix = "host-another-prefix"
    }
    check = {
      "uddi.ddns_generated_prefix" = "host-another-prefix"
    }
  }

}

case "dhcp_options" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # resource "infoblox_dhcp_option_code_unknown" "test" {
  #   uddi = {
  #     code = 234
  #     name = "test_dhcp_option_code"
  #     option_space = infoblox_dhcp_option_space_unknown.test.id
  #     type = "boolean"
  #   }
  # }
  # resource "infoblox_dhcp_option_group_unknown" "test" {
  #   uddi = {
  #     name = "\"og-\"+name"
  #     protocol = "ip4"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address      = "{{random_ipv6}}"
      cidr         = 128
      # space        = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      name         = "option_group_test"
      dhcp_options = [{ type = "option", option_code = "dhcp/option_code/88810e74-3909-4346-b0e9-de28d4a7c3ce", option_value = "http://example.com" }]
    }
    check = {
      "uddi.dhcp_options.#"              = "1"
      "uddi.dhcp_options.0.option_value" = "http://example.com"
    }
  }

  step {
    uddi {
      address      = "{{random_ipv6}}"
      cidr         = 128
      # space        = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      name         = "option_group_test"
      dhcp_options = [{ type = "option", option_code = "dhcp/option_code/8a67f062-edd7-4a5f-98d3-00fb4fad944b", option_value = "http://example.org/boot" }]
    }
    check = {
      "uddi.dhcp_options.#" = "1"
      "uddi.dhcp_options.0.option_value" = "http://example.org/boot"
    }
  }

}

case "ddns_send_updates" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address           = "{{random_ipv6}}"
      cidr              = 128
      # space             = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      ddns_send_updates = true
    }
    check = {
      "uddi.ddns_send_updates" = "true"
    }
  }

  step {
    uddi {
      address           = "{{random_ipv6}}"
      cidr              = 128
      # space             = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      ddns_send_updates = false
    }
    check = {
      "uddi.ddns_send_updates" = "false"
    }
  }

}

case "ddns_ttl_percent" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address          = "{{random_ipv6}}"
      cidr             = 128
      # space            = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      ddns_ttl_percent = 20
    }
    check = {
      "uddi.ddns_ttl_percent" = "20"
    }
  }

  step {
    uddi {
      address          = "{{random_ipv6}}"
      cidr             = 128
      # space            = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      ddns_ttl_percent = 40
    }
    check = {
      "uddi.ddns_ttl_percent" = "40"
    }
  }

}

case "ddns_update_on_renew" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address              = "{{random_ipv6}}"
      cidr                 = 128
      # space                = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      ddns_update_on_renew = true
    }
    check = {
      "uddi.ddns_update_on_renew" = "true"
    }
  }

  step {
    uddi {
      address              = "{{random_ipv6}}"
      cidr                 = 128
      # space                = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      ddns_update_on_renew = false
    }
    check = {
      "uddi.ddns_update_on_renew" = "false"
    }
  }

}

case "dhcp_config" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address     = "{{random_ipv6}}"
      cidr        = 128
      # space       = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      dhcp_config = { allow_unknown_v6 = true, ignore_client_uid = true, lease_time_v6 = 60 }
    }
    check = {
      "uddi.dhcp_config.allow_unknown_v6"  = "true"
      "uddi.dhcp_config.ignore_client_uid" = "true"
      "uddi.dhcp_config.lease_time_v6"     = "60"
    }
  }

  step {
    uddi {
      address     = "{{random_ipv6}}"
      cidr        = 128
      # space       = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      dhcp_config = { allow_unknown_v6 = false, ignore_client_uid = false, lease_time_v6 = 65 }
    }
    check = {
      "uddi.dhcp_config.allow_unknown_v6"  = "false"
      "uddi.dhcp_config.ignore_client_uid" = "false"
      "uddi.dhcp_config.lease_time_v6"     = "65"
    }
  }

}

case "disable_dhcp" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address      = "{{random_ipv6}}"
      cidr         = 128
      # space        = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      disable_dhcp = true
    }
    check = {
      "uddi.disable_dhcp" = "true"
    }
  }

  step {
    uddi {
      address      = "{{random_ipv6}}"
      cidr         = 128
      # space        = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      disable_dhcp = false
    }
    check = {
      "uddi.disable_dhcp" = "false"
    }
  }

}

case "hostname_rewrite_char" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address               = "{{random_ipv6}}"
      cidr                  = 128
      # space                 = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      hostname_rewrite_char = "+"
    }
    check = {
      "uddi.hostname_rewrite_char" = "+"
    }
  }

  step {
    uddi {
      address               = "{{random_ipv6}}"
      cidr                  = 128
      # space                 = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      hostname_rewrite_char = "/"
    }
    check = {
      "uddi.hostname_rewrite_char" = "/"
    }
  }

}

case "hostname_rewrite_enabled" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address                  = "{{random_ipv6}}"
      cidr                     = 128
      # space                    = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      hostname_rewrite_enabled = true
    }
    check = {
      "uddi.hostname_rewrite_enabled" = "true"
    }
  }

  step {
    uddi {
      address                  = "{{random_ipv6}}"
      cidr                     = 128
      # space                    = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      hostname_rewrite_enabled = false
    }
    check = {
      "uddi.hostname_rewrite_enabled" = "false"
    }
  }

}

case "hostname_rewrite_regex" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address                = "{{random_ipv6}}"
      cidr                   = 128
      # space                  = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      hostname_rewrite_regex = "[^a-z]"
    }
    check = {
      "uddi.hostname_rewrite_regex" = "[^a-z]"
    }
  }

  step {
    uddi {
      address                = "{{random_ipv6}}"
      cidr                   = 128
      # space                  = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      hostname_rewrite_regex = "[^0-9]"
    }
    check = {
      "uddi.hostname_rewrite_regex" = "[^0-9]"
    }
  }

}

case "inheritance_sources" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address             = "{{random_ipv6}}"
      cidr                = 128
      # space               = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      inheritance_sources = { asm_config = { action = "inherit", asm_enable_block = { action = "inherit" }, asm_growth_block = { action = "inherit" }, asm_threshold = { action = "inherit" }, forecast_period = { action = "inherit" }, history = { action = "inherit" }, min_total = { action = "inherit" }, min_unused = { action = "inherit" } }, dhcp_config = { allow_unknown = { action = "inherit" }, allow_unknown_v6 = { action = "inherit" }, filters = { action = "inherit" }, filters_v6 = { action = "inherit" }, ignore_client_uid = { action = "inherit" }, ignore_list = { action = "inherit" }, lease_time = { action = "inherit" }, lease_time_v6 = { action = "inherit" } }, ddns_client_update = { action = "inherit" }, ddns_conflict_resolution_mode = { action = "inherit" }, ddns_enabled = { action = "inherit" }, ddns_hostname_block = { action = "inherit" }, ddns_ttl_percent = { action = "inherit" }, ddns_update_block = { action = "inherit" }, ddns_update_on_renew = { action = "inherit" }, ddns_use_conflict_resolution = { action = "inherit" }, header_option_filename = { action = "inherit" }, header_option_server_address = { action = "inherit" }, header_option_server_name = { action = "inherit" }, hostname_rewrite_block = { action = "inherit" } }
    }
    check = {
      "uddi.inheritance_sources.asm_config.asm_enable_block.action"   = "inherit"
      "uddi.inheritance_sources.asm_config.asm_growth_block.action"   = "inherit"
      "uddi.inheritance_sources.asm_config.asm_threshold.action"      = "inherit"
      "uddi.inheritance_sources.asm_config.forecast_period.action"    = "inherit"
      "uddi.inheritance_sources.asm_config.history.action"            = "inherit"
      "uddi.inheritance_sources.asm_config.min_total.action"          = "inherit"
      "uddi.inheritance_sources.asm_config.min_unused.action"         = "inherit"
      "uddi.inheritance_sources.ddns_client_update.action"            = "inherit"
      "uddi.inheritance_sources.ddns_conflict_resolution_mode.action" = "inherit"
      "uddi.inheritance_sources.ddns_enabled.action"                  = "inherit"
      "uddi.inheritance_sources.ddns_hostname_block.action"           = "inherit"
      "uddi.inheritance_sources.ddns_ttl_percent.action"              = "inherit"
      "uddi.inheritance_sources.ddns_update_block.action"             = "inherit"
      "uddi.inheritance_sources.ddns_update_on_renew.action"          = "inherit"
      "uddi.inheritance_sources.ddns_use_conflict_resolution.action"  = "inherit"
      "uddi.inheritance_sources.dhcp_config.allow_unknown.action"     = "inherit"
      "uddi.inheritance_sources.dhcp_config.allow_unknown_v6.action"  = "inherit"
      "uddi.inheritance_sources.dhcp_config.filters.action"           = "inherit"
      "uddi.inheritance_sources.dhcp_config.filters_v6.action"        = "inherit"
      "uddi.inheritance_sources.dhcp_config.ignore_client_uid.action" = "inherit"
      "uddi.inheritance_sources.dhcp_config.lease_time.action"        = "inherit"
      "uddi.inheritance_sources.dhcp_config.lease_time_v6.action"     = "inherit"
      "uddi.inheritance_sources.header_option_filename.action"        = "inherit"
      "uddi.inheritance_sources.header_option_server_address.action"  = "inherit"
      "uddi.inheritance_sources.header_option_server_name.action"     = "inherit"
      "uddi.inheritance_sources.hostname_rewrite_block.action"        = "inherit"
    }
  }

  step {
    uddi {
      address             = "{{random_ipv6}}"
      cidr                = 128
      # space               = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      inheritance_sources = { asm_config = { action = "override", asm_enable_block = { action = "override" }, asm_growth_block = { action = "override" }, asm_threshold = { action = "override" }, forecast_period = { action = "override" }, history = { action = "override" }, min_total = { action = "override" }, min_unused = { action = "override" } }, dhcp_config = { allow_unknown = { action = "override" }, allow_unknown_v6 = { action = "override" }, filters = { action = "override" }, filters_v6 = { action = "override" }, ignore_client_uid = { action = "override" }, ignore_list = { action = "override" }, lease_time = { action = "override" }, lease_time_v6 = { action = "override" } }, ddns_client_update = { action = "override" }, ddns_conflict_resolution_mode = { action = "override" }, ddns_enabled = { action = "inherit" }, ddns_hostname_block = { action = "override" }, ddns_ttl_percent = { action = "override" }, ddns_update_block = { action = "override" }, ddns_update_on_renew = { action = "override" }, ddns_use_conflict_resolution = { action = "override" }, header_option_filename = { action = "override" }, header_option_server_address = { action = "override" }, header_option_server_name = { action = "override" }, hostname_rewrite_block = { action = "override" } }
    }
    check = {
      "uddi.inheritance_sources.asm_config.asm_enable_block.action"   = "override"
      "uddi.inheritance_sources.asm_config.asm_growth_block.action"   = "override"
      "uddi.inheritance_sources.asm_config.asm_threshold.action"      = "override"
      "uddi.inheritance_sources.asm_config.forecast_period.action"    = "override"
      "uddi.inheritance_sources.asm_config.history.action"            = "override"
      "uddi.inheritance_sources.asm_config.min_total.action"          = "override"
      "uddi.inheritance_sources.asm_config.min_unused.action"         = "override"
      "uddi.inheritance_sources.ddns_client_update.action"            = "override"
      "uddi.inheritance_sources.ddns_conflict_resolution_mode.action" = "override"
      "uddi.inheritance_sources.ddns_hostname_block.action"           = "override"
      "uddi.inheritance_sources.ddns_ttl_percent.action"              = "override"
      "uddi.inheritance_sources.ddns_update_block.action"             = "override"
      "uddi.inheritance_sources.ddns_update_on_renew.action"          = "override"
      "uddi.inheritance_sources.ddns_use_conflict_resolution.action"  = "override"
      "uddi.inheritance_sources.dhcp_config.allow_unknown.action"     = "override"
      "uddi.inheritance_sources.dhcp_config.allow_unknown_v6.action"  = "override"
      "uddi.inheritance_sources.dhcp_config.filters.action"           = "override"
      "uddi.inheritance_sources.dhcp_config.filters_v6.action"        = "override"
      "uddi.inheritance_sources.dhcp_config.ignore_client_uid.action" = "override"
      "uddi.inheritance_sources.dhcp_config.lease_time.action"        = "override"
      "uddi.inheritance_sources.dhcp_config.lease_time_v6.action"     = "override"
      "uddi.inheritance_sources.header_option_filename.action"        = "override"
      "uddi.inheritance_sources.header_option_server_address.action"  = "override"
      "uddi.inheritance_sources.header_option_server_name.action"     = "override"
      "uddi.inheritance_sources.hostname_rewrite_block.action"        = "override"
    }
  }

}

case "multiple_federated_realms" {
  backend     = "uddi"
  skip        = true
  skip_reason = "t.Skip: Skipping this test as Multiple federated realms are not supported"
  parallel    = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_federated_realm_unknown" "%s" {
    uddi = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    uddi {
      address = "192.168.0.0"
      cidr    = 16
      # space   = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
    }
    check = {
      "uddi.federated_realms.#" = "5"
    }
  }

  step {
    uddi {
      address = "192.168.0.0"
      cidr    = 16
      # space   = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
    }
    check = {
      "uddi.federated_realms.#" = "5"
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address = "{{random_ipv6}}"
      cidr    = 128
      # space   = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      name    = "subnet_name"
    }
    check = {
      "uddi.name" = "subnet_name"
    }
  }

  step {
    uddi {
      address = "{{random_ipv6}}"
      cidr    = 128
      # space   = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      name    = "subnet_name_updated"
    }
    check = {
      "uddi.name" = "subnet_name_updated"
    }
  }

}

case "renew_time_and_rebind_time" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address     = "{{random_ipv6}}"
      cidr        = 128
      # space       = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      rebind_time = 60
      renew_time  = 50
    }
    check = {
      "uddi.rebind_time" = "60"
      "uddi.renew_time"  = "50"
    }
  }

  step {
    uddi {
      address     = "{{random_ipv6}}"
      cidr        = 128
      # space       = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      rebind_time = 90
      renew_time  = 80
    }
    check = {
      "uddi.rebind_time" = "90"
      "uddi.renew_time"  = "80"
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address = "{{random_ipv6}}"
      cidr    = 128
      # space   = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      tags    = { tag1 = "value1", tag2 = "value2" }
    }
    check = {
      "uddi.tags.tag1" = "value1"
      "uddi.tags.tag2" = "value2"
    }
  }

  step {
    uddi {
      address = "{{random_ipv6}}"
      cidr    = 128
      # space   = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      tags    = { tag2 = "value2changed", tag3 = "value3" }
    }
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }

}
