# Auto-generated resource acceptance-test cases for Networkcontainer.
case "basic" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address = "192.168.0.0"
      cidr    = 16
      space   = infoblox_ip_space.test.id
    }
    check = {
      "uddi.address"                       = "192.168.0.0"
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
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address = "192.168.0.0"
      cidr    = 16
      space   = infoblox_ip_space.test.id
    }
  }

}

case "address" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address = "192.168.0.0"
      cidr    = 16
      space   = infoblox_ip_space.test.id
    }
    check = {
      "uddi.address" = "192.168.0.0"
    }
  }

  step {
    uddi {
      address = "10.0.0.0"
      cidr    = 16
      space   = infoblox_ip_space.test.id
    }
    check = {
      "uddi.address" = "10.0.0.0"
    }
  }

}

case "asm_config" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address    = "10.0.0.0"
      cidr       = 16
      space      = infoblox_ip_space.test.id
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
      address    = "10.0.0.0"
      cidr       = 16
      space      = infoblox_ip_space.test.id
      asm_config = { asm_threshold = 90, enable = false, enable_notification = false, forecast_period = 14, growth_factor = 60, growth_type = "count", history = 40, min_total = 60, min_unused = 50, reenable_date = "2020-01-10T10:11:22Z" }
    }
    check = {
      "uddi.asm_config.asm_threshold"       = "90"
      "uddi.asm_config.enable"              = "false"
      "uddi.asm_config.enable_notification" = "false"
      "uddi.asm_config.forecast_period"     = "14"
      "uddi.asm_config.growth_factor"       = "60"
      "uddi.asm_config.growth_type"         = "count"
      "uddi.asm_config.history"             = "40"
      "uddi.asm_config.min_total"           = "60"
      "uddi.asm_config.min_unused"          = "50"
      "uddi.asm_config.reenable_date"       = "2020-01-10T10:11:22Z"
    }
  }

}

case "cidr" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address = "192.168.0.0"
      cidr    = 16
      space   = infoblox_ip_space.test.id
    }
    check = {
      "uddi.cidr" = "16"
    }
  }

  step {
    uddi {
      address = "192.168.0.0"
      cidr    = 24
      space   = infoblox_ip_space.test.id
    }
    check = {
      "uddi.cidr" = "24"
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address = "192.168.0.0"
      cidr    = 16
      space   = infoblox_ip_space.test.id
      comment = "This address block is created through Terraform"
    }
    check = {
      "uddi.comment" = "This address block is created through Terraform"
    }
  }

  step {
    uddi {
      address = "192.168.0.0"
      cidr    = 16
      space   = infoblox_ip_space.test.id
      comment = "This address block was created through Terraform"
    }
    check = {
      "uddi.comment" = "This address block was created through Terraform"
    }
  }

}

case "compartment_id" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address        = "192.168.0.0"
      cidr           = 16
      space          = infoblox_ip_space.test.id
      compartment_id = ""
    }
    check = {
      "uddi.compartment_id" = ""
    }
  }

  step {
    uddi {
      address = "192.168.0.0"
      cidr    = 16
      space   = infoblox_ip_space.test.id
    }
    check = {
      "uddi.compartment_id" = ""
    }
  }

}

case "ddns_client_update" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address            = "192.168.0.0"
      cidr               = 16
      space              = infoblox_ip_space.test.id
      ddns_client_update = "client"
    }
    check = {
      "uddi.ddns_client_update" = "client"
    }
  }

  step {
    uddi {
      address            = "192.168.0.0"
      cidr               = 16
      space              = infoblox_ip_space.test.id
      ddns_client_update = "over_no_update"
    }
    check = {
      "uddi.ddns_client_update" = "over_no_update"
    }
  }

}

case "ddns_domain" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address     = "192.168.0.0"
      cidr        = 16
      space       = infoblox_ip_space.test.id
      ddns_domain = "test.com"
    }
    check = {
      "uddi.ddns_domain" = "test.com"
    }
  }

  step {
    uddi {
      address     = "192.168.0.0"
      cidr        = 16
      space       = infoblox_ip_space.test.id
      ddns_domain = "test123.com"
    }
    check = {
      "uddi.ddns_domain" = "test123.com"
    }
  }

}

case "ddns_generate_name" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address            = "192.168.0.0"
      cidr               = 16
      space              = infoblox_ip_space.test.id
      ddns_generate_name = false
    }
    check = {
      "uddi.ddns_generate_name" = "false"
    }
  }

  step {
    uddi {
      address            = "192.168.0.0"
      cidr               = 16
      space              = infoblox_ip_space.test.id
      ddns_generate_name = true
    }
    check = {
      "uddi.ddns_generate_name" = "true"
    }
  }

}

case "ddns_generated_prefix" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address               = "192.168.0.0"
      cidr                  = 16
      space                 = infoblox_ip_space.test.id
      ddns_generated_prefix = "ut"
    }
    check = {
      "uddi.ddns_generated_prefix" = "ut"
    }
  }

  step {
    uddi {
      address               = "192.168.0.0"
      cidr                  = 16
      space                 = infoblox_ip_space.test.id
      ddns_generated_prefix = "ut-ut"
    }
    check = {
      "uddi.ddns_generated_prefix" = "ut-ut"
    }
  }

}

case "ddns_send_updates" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address           = "192.168.0.0"
      cidr              = 16
      space             = infoblox_ip_space.test.id
      ddns_send_updates = true
    }
    check = {
      "uddi.ddns_send_updates" = "true"
    }
  }

  step {
    uddi {
      address           = "192.168.0.0"
      cidr              = 16
      space             = infoblox_ip_space.test.id
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
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address          = "192.168.0.0"
      cidr             = 16
      space            = infoblox_ip_space.test.id
      ddns_ttl_percent = 25
    }
    check = {
      "uddi.ddns_ttl_percent" = "25"
    }
  }

  step {
    uddi {
      address          = "192.168.0.0"
      cidr             = 16
      space            = infoblox_ip_space.test.id
      ddns_ttl_percent = 75
    }
    check = {
      "uddi.ddns_ttl_percent" = "75"
    }
  }

}

case "ddns_update_on_renew" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address              = "192.168.0.0"
      cidr                 = 16
      space                = infoblox_ip_space.test.id
      ddns_update_on_renew = false
    }
    check = {
      "uddi.ddns_update_on_renew" = "false"
    }
  }

  step {
    uddi {
      address              = "192.168.0.0"
      cidr                 = 16
      space                = infoblox_ip_space.test.id
      ddns_update_on_renew = true
    }
    check = {
      "uddi.ddns_update_on_renew" = "true"
    }
  }

}

case "ddns_use_conflict_resolution" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address                      = "192.168.0.0"
      cidr                         = 16
      space                        = infoblox_ip_space.test.id
      ddns_use_conflict_resolution = true
    }
    check = {
      "uddi.ddns_use_conflict_resolution" = "true"
    }
  }

  step {
    uddi {
      address                      = "192.168.0.0"
      cidr                         = 16
      space                        = infoblox_ip_space.test.id
      ddns_use_conflict_resolution = false
    }
    check = {
      "uddi.ddns_use_conflict_resolution" = "false"
    }
  }

}

case "dhcp_config" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address     = "192.168.0.0"
      cidr        = 16
      space       = infoblox_ip_space.test.id
      dhcp_config = { allow_unknown = true, allow_unknown_v6 = true, ignore_client_uid = true, lease_time = 50, lease_time_v6 = 60 }
    }
    check = {
      "uddi.dhcp_config.allow_unknown"     = "true"
      "uddi.dhcp_config.allow_unknown_v6"  = "true"
      "uddi.dhcp_config.ignore_client_uid" = "true"
      "uddi.dhcp_config.lease_time"        = "50"
      "uddi.dhcp_config.lease_time_v6"     = "60"
    }
  }

  step {
    uddi {
      address     = "192.168.0.0"
      cidr        = 16
      space       = infoblox_ip_space.test.id
      dhcp_config = { allow_unknown = false, allow_unknown_v6 = true, ignore_client_uid = false, lease_time = 150, lease_time_v6 = 160 }
    }
    check = {
      "uddi.dhcp_config.allow_unknown"     = "false"
      "uddi.dhcp_config.allow_unknown_v6"  = "true"
      "uddi.dhcp_config.ignore_client_uid" = "false"
      "uddi.dhcp_config.lease_time"        = "150"
      "uddi.dhcp_config.lease_time_v6"     = "160"
    }
  }

}

case "dhcp_options" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_dhcp_option_code_unknown" "test" {
    uddi = {
      code = 234
      name = "test_dhcp_option_code"
      option_space = infoblox_dhcp_option_space_unknown.test.id
      type = "boolean"
    }
  }
  resource "infoblox_dhcp_option_group_unknown" "test" {
    uddi = {
      name = "\"og-\"+optionSpace"
      protocol = "ip4"
    }
  }
  PREREQ

  step {
    uddi {
      address      = "192.168.0.0"
      cidr         = 16
      space        = infoblox_ip_space.test.id
      dhcp_options = [{ type = "option", option_code = infoblox_dhcp_option_code_unknown.test.id, option_value = true }]
    }
    check = {
      "uddi.dhcp_options.#"              = "1"
      "uddi.dhcp_options.0.option_value" = "true"
    }
  }

  step {
    uddi {
      address      = "192.168.0.0"
      cidr         = 16
      space        = infoblox_ip_space.test.id
      dhcp_options = [{ type = "group", group = infoblox_dhcp_option_group_unknown.test.id }]
    }
    check = {
      "uddi.dhcp_options.#" = "1"
    }
  }

}

case "header_option_filename" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address                = "192.168.0.0"
      cidr                   = 16
      space                  = infoblox_ip_space.test.id
      header_option_filename = "testfile"
    }
    check = {
      "uddi.header_option_filename" = "testfile"
    }
  }

  step {
    uddi {
      address                = "192.168.0.0"
      cidr                   = 16
      space                  = infoblox_ip_space.test.id
      header_option_filename = "testfile1"
    }
    check = {
      "uddi.header_option_filename" = "testfile1"
    }
  }

}

case "header_option_server_address" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address                      = "192.168.0.0"
      cidr                         = 16
      space                        = infoblox_ip_space.test.id
      header_option_server_address = "1.1.1.1"
    }
    check = {
      "uddi.header_option_server_address" = "1.1.1.1"
    }
  }

  step {
    uddi {
      address                      = "192.168.0.0"
      cidr                         = 16
      space                        = infoblox_ip_space.test.id
      header_option_server_address = "2.2.2.2"
    }
    check = {
      "uddi.header_option_server_address" = "2.2.2.2"
    }
  }

}

case "header_option_server_name" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address                   = "192.168.0.0"
      cidr                      = 16
      space                     = infoblox_ip_space.test.id
      header_option_server_name = "test"
    }
    check = {
      "uddi.header_option_server_name" = "test"
    }
  }

  step {
    uddi {
      address                   = "192.168.0.0"
      cidr                      = 16
      space                     = infoblox_ip_space.test.id
      header_option_server_name = "test-1"
    }
    check = {
      "uddi.header_option_server_name" = "test-1"
    }
  }

}

case "hostname_rewrite_char" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address               = "192.168.0.0"
      cidr                  = 16
      space                 = infoblox_ip_space.test.id
      hostname_rewrite_char = "a"
    }
    check = {
      "uddi.hostname_rewrite_char" = "a"
    }
  }

  step {
    uddi {
      address               = "192.168.0.0"
      cidr                  = 16
      space                 = infoblox_ip_space.test.id
      hostname_rewrite_char = "c"
    }
    check = {
      "uddi.hostname_rewrite_char" = "c"
    }
  }

}

case "hostname_rewrite_enabled" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address                  = "192.168.0.0"
      cidr                     = 16
      space                    = infoblox_ip_space.test.id
      hostname_rewrite_enabled = true
    }
    check = {
      "uddi.hostname_rewrite_enabled" = "true"
    }
  }

  step {
    uddi {
      address                  = "192.168.0.0"
      cidr                     = 16
      space                    = infoblox_ip_space.test.id
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
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address                = "192.168.0.0"
      cidr                   = 16
      space                  = infoblox_ip_space.test.id
      hostname_rewrite_regex = "[^a-z]"
    }
    check = {
      "uddi.hostname_rewrite_regex" = "[^a-z]"
    }
  }

  step {
    uddi {
      address                = "192.168.0.0"
      cidr                   = 16
      space                  = infoblox_ip_space.test.id
      hostname_rewrite_regex = "[^g-hG-H0-9_.]"
    }
    check = {
      "uddi.hostname_rewrite_regex" = "[^g-hG-H0-9_.]"
    }
  }

}

case "inheritance_sources" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address             = "192.168.0.0"
      cidr                = 16
      space               = infoblox_ip_space.test.id
      inheritance_sources = { asm_config = { action = "inherit", asm_enable_block = { action = "inherit" }, asm_growth_block = { action = "inherit" }, asm_threshold = { action = "inherit" }, forecast_period = { action = "inherit" }, history = { action = "inherit" }, min_total = { action = "inherit" }, min_unused = { action = "inherit" } }, ddns_client_update = { action = "inherit" }, ddns_conflict_resolution_mode = { action = "inherit" }, ddns_enabled = { action = "inherit" }, ddns_hostname_block = { action = "inherit" }, ddns_ttl_percent = { action = "inherit" }, ddns_update_block = { action = "inherit" }, ddns_update_on_renew = { action = "inherit" }, ddns_use_conflict_resolution = { action = "inherit" }, header_option_filename = { action = "inherit" }, header_option_server_address = { action = "inherit" }, header_option_server_name = { action = "inherit" }, hostname_rewrite_block = { action = "inherit" } }
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
      "uddi.inheritance_sources.header_option_filename.action"        = "inherit"
      "uddi.inheritance_sources.header_option_server_address.action"  = "inherit"
      "uddi.inheritance_sources.header_option_server_name.action"     = "inherit"
      "uddi.inheritance_sources.hostname_rewrite_block.action"        = "inherit"
    }
  }

  step {
    uddi {
      address             = "192.168.0.0"
      cidr                = 16
      space               = infoblox_ip_space.test.id
      inheritance_sources = { asm_config = { action = "override", asm_enable_block = { action = "override" }, asm_growth_block = { action = "override" }, asm_threshold = { action = "override" }, forecast_period = { action = "override" }, history = { action = "override" }, min_total = { action = "override" }, min_unused = { action = "override" } }, ddns_client_update = { action = "override" }, ddns_conflict_resolution_mode = { action = "override" }, ddns_enabled = { action = "inherit" }, ddns_hostname_block = { action = "override" }, ddns_ttl_percent = { action = "override" }, ddns_update_block = { action = "override" }, ddns_update_on_renew = { action = "override" }, ddns_use_conflict_resolution = { action = "override" }, header_option_filename = { action = "override" }, header_option_server_address = { action = "override" }, header_option_server_name = { action = "override" }, hostname_rewrite_block = { action = "override" } }
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
      "uddi.inheritance_sources.header_option_filename.action"        = "override"
      "uddi.inheritance_sources.header_option_server_address.action"  = "override"
      "uddi.inheritance_sources.header_option_server_name.action"     = "override"
      "uddi.inheritance_sources.hostname_rewrite_block.action"        = "override"
    }
  }

}

case "multiple_federated_realms" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_federated_realm_unknown" "%s" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address = "192.168.0.0"
      cidr    = 16
      space   = infoblox_ip_space.test.id
    }
    check = {
      "uddi.federated_realms.#" = "5"
    }
  }

  step {
    uddi {
      address = "192.168.0.0"
      cidr    = 16
      space   = infoblox_ip_space.test.id
    }
    check = {
      "uddi.federated_realms.#" = "5"
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address = "192.168.0.0"
      cidr    = 16
      space   = infoblox_ip_space.test.id
      name    = "test_name"
    }
    check = {
      "uddi.name" = "test_name"
    }
  }

  step {
    uddi {
      address = "192.168.0.0"
      cidr    = 16
      space   = infoblox_ip_space.test.id
      name    = "test_name_1"
    }
    check = {
      "uddi.name" = "test_name_1"
    }
  }

}

case "space" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "one" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address = "192.168.0.0"
      cidr    = 16
      space   = infoblox_ip_space.one.id
    }
  }

  step {
    uddi {
      address = "192.168.0.0"
      cidr    = 16
      space   = infoblox_ip_space.two.id
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address = "192.168.0.0"
      cidr    = 16
      space   = infoblox_ip_space.test.id
      tags    = { tag1 = "value1", tag2 = "value2" }
    }
    check = {
      "uddi.tags.tag1" = "value1"
      "uddi.tags.tag2" = "value2"
    }
  }

  step {
    uddi {
      address = "192.168.0.0"
      cidr    = 16
      space   = infoblox_ip_space.test.id
      tags    = { tag2 = "value2changed", tag3 = "value3" }
    }
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }

}

case "next_available_id" {
  backend     = "uddi"
  skip        = true
  skip_reason = "helper declares 2 'bloxone_ipam_address_block' resource blocks with no single func_call target (ambiguous which is the resource under test)"
}

case "next_available_id_count" {
  backend     = "uddi"
  skip        = true
  skip_reason = "helper declares 2 'bloxone_ipam_address_block' resource blocks with no single func_call target (ambiguous which is the resource under test)"
}
