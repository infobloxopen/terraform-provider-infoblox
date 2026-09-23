# Auto-generated resource acceptance-test cases for Networkcontainer.
case "basic" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address = "{{random_ipv4_network}}"
      cidr    = 16
      space   = infoblox_network_view.test.id
    }
    check = {
      "uddi.address"                       = "{{random_ipv4_network}}"
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
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }

  step {
    uddi {
      address = "{{random_ipv4_network}}"
      cidr    = 16
      space   = infoblox_network_view.test.id
    }
  }

}

case "address" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address = "{{random_ipv4_network}}"
      cidr    = 16
      space   = infoblox_network_view.test.id
    }
    check = {
      "uddi.address" = "{{random_ipv4_network}}"
    }
  }

  step {
    uddi {
      address = "{{random_ipv4_network2}}"
      cidr    = 16
      space   = infoblox_network_view.test.id
    }
    check = {
      "uddi.address" = "{{random_ipv4_network2}}"
    }
  }

}

case "asm_config" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address    = "{{random_ipv4_network}}"
      cidr       = 16
      space      = infoblox_network_view.test.id
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
      address    = "{{random_ipv4_network}}"
      cidr       = 16
      space      = infoblox_network_view.test.id
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
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address = "{{random_ipv4_network}}"
      cidr    = 16
      space   = infoblox_network_view.test.id
    }
    check = {
      "uddi.cidr" = "16"
    }
  }

  step {
    uddi {
      address = "{{random_ipv4_network}}"
      cidr    = 24
      space   = infoblox_network_view.test.id
    }
    check = {
      "uddi.cidr" = "24"
    }
  }

}

case "comment" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address = "{{random_ipv4_network}}"
      cidr    = 16
      space   = infoblox_network_view.test.id
      comment = "This address block is created through Terraform"
    }
    check = {
      "uddi.comment" = "This address block is created through Terraform"
    }
  }

  step {
    uddi {
      address = "{{random_ipv4_network}}"
      cidr    = 16
      space   = infoblox_network_view.test.id
      comment = "This address block was created through Terraform"
    }
    check = {
      "uddi.comment" = "This address block was created through Terraform"
    }
  }

}

case "compartment_id" {
  backend           = "uddi"
  parallel          = true
  skip_if_env_empty = ["UDDI_COMPARTMENT_ID_1"]
  skip_reason       = "UDDI_COMPARTMENT_ID_1 environment variable must be set for this test to run"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address        = "{{random_ipv4_network}}"
      cidr           = 16
      space          = infoblox_network_view.test.id
      compartment_id = "{{uddi_compartment_id_1}}"
    }
    check = {
      "uddi.compartment_id" = "{{uddi_compartment_id_1}}"
    }
  }

  step {
    uddi {
      address        = "{{random_ipv4_network}}"
      cidr           = 16
      space          = infoblox_network_view.test.id
      compartment_id = ""
    }
    check = {
      "uddi.compartment_id" = ""
    }
  }

}

case "ddns_client_update" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address            = "{{random_ipv4_network}}"
      cidr               = 16
      space              = infoblox_network_view.test.id
      ddns_client_update = "client"
    }
    check = {
      "uddi.ddns_client_update" = "client"
    }
  }

  step {
    uddi {
      address            = "{{random_ipv4_network}}"
      cidr               = 16
      space              = infoblox_network_view.test.id
      ddns_client_update = "over_no_update"
    }
    check = {
      "uddi.ddns_client_update" = "over_no_update"
    }
  }

}

case "ddns_domain" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address     = "{{random_ipv4_network}}"
      cidr        = 16
      space       = infoblox_network_view.test.id
      ddns_domain = "test.com"
    }
    check = {
      "uddi.ddns_domain" = "test.com"
    }
  }

  step {
    uddi {
      address     = "{{random_ipv4_network}}"
      cidr        = 16
      space       = infoblox_network_view.test.id
      ddns_domain = "test123.com"
    }
    check = {
      "uddi.ddns_domain" = "test123.com"
    }
  }

}

case "ddns_generate_name" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address            = "{{random_ipv4_network}}"
      cidr               = 16
      space              = infoblox_network_view.test.id
      ddns_generate_name = false
    }
    check = {
      "uddi.ddns_generate_name" = "false"
    }
  }

  step {
    uddi {
      address            = "{{random_ipv4_network}}"
      cidr               = 16
      space              = infoblox_network_view.test.id
      ddns_generate_name = true
    }
    check = {
      "uddi.ddns_generate_name" = "true"
    }
  }

}

case "ddns_generated_prefix" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address               = "{{random_ipv4_network}}"
      cidr                  = 16
      space                 = infoblox_network_view.test.id
      ddns_generated_prefix = "ut"
    }
    check = {
      "uddi.ddns_generated_prefix" = "ut"
    }
  }

  step {
    uddi {
      address               = "{{random_ipv4_network}}"
      cidr                  = 16
      space                 = infoblox_network_view.test.id
      ddns_generated_prefix = "ut-ut"
    }
    check = {
      "uddi.ddns_generated_prefix" = "ut-ut"
    }
  }

}

case "ddns_send_updates" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address           = "{{random_ipv4_network}}"
      cidr              = 16
      space             = infoblox_network_view.test.id
      ddns_send_updates = true
    }
    check = {
      "uddi.ddns_send_updates" = "true"
    }
  }

  step {
    uddi {
      address           = "{{random_ipv4_network}}"
      cidr              = 16
      space             = infoblox_network_view.test.id
      ddns_send_updates = false
    }
    check = {
      "uddi.ddns_send_updates" = "false"
    }
  }

}

case "ddns_ttl_percent" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address          = "{{random_ipv4_network}}"
      cidr             = 16
      space            = infoblox_network_view.test.id
      ddns_ttl_percent = 25
    }
    check = {
      "uddi.ddns_ttl_percent" = "25"
    }
  }

  step {
    uddi {
      address          = "{{random_ipv4_network}}"
      cidr             = 16
      space            = infoblox_network_view.test.id
      ddns_ttl_percent = 75
    }
    check = {
      "uddi.ddns_ttl_percent" = "75"
    }
  }

}

case "ddns_update_on_renew" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address              = "{{random_ipv4_network}}"
      cidr                 = 16
      sspace               = infoblox_network_view.test.id
      ddns_update_on_renew = false
    }
    check = {
      "uddi.ddns_update_on_renew" = "false"
    }
  }

  step {
    uddi {
      address              = "{{random_ipv4_network}}"
      cidr                 = 16
      sspace               = infoblox_network_view.test.id
      ddns_update_on_renew = true
    }
    check = {
      "uddi.ddns_update_on_renew" = "true"
    }
  }

}

case "ddns_use_conflict_resolution" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address                      = "{{random_ipv4_network}}"
      cidr                         = 16
      space                        = infoblox_network_view.test.id
      ddns_use_conflict_resolution = true
    }
    check = {
      "uddi.ddns_use_conflict_resolution" = "true"
    }
  }

  step {
    uddi {
      address                      = "{{random_ipv4_network}}"
      cidr                         = 16
      space                        = infoblox_network_view.test.id
      ddns_use_conflict_resolution = false
    }
    check = {
      "uddi.ddns_use_conflict_resolution" = "false"
    }
  }

}

case "dhcp_config" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address     = "{{random_ipv4_network}}"
      cidr        = 16
      space       = infoblox_network_view.test.id
      dhcp_config = { allow_unknown = true, ignore_client_uid = true, lease_time = 50 }
    }
    check = {
      "uddi.dhcp_config.allow_unknown"     = "true"
      "uddi.dhcp_config.ignore_client_uid" = "true"
      "uddi.dhcp_config.lease_time"        = "50"
    }
  }

  step {
    uddi {
      address     = "{{random_ipv4_network}}"
      cidr        = 16
      space       = infoblox_network_view.test.id
      dhcp_config = { allow_unknown = false, ignore_client_uid = false, lease_time = 150 }
    }
    check = {
      "uddi.dhcp_config.allow_unknown"     = "false"
      "uddi.dhcp_config.ignore_client_uid" = "false"
      "uddi.dhcp_config.lease_time"        = "150"
    }
  }

}

case "dhcp_options" {
  backend           = "uddi"
  parallel          = true
  skip_if_env_empty = ["UDDI_OPTION_GROUP_1_ID"]
  skip_reason       = "UDDI_OPTION_GROUP_1_ID environment variable must be set for this test to run"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random3}}"
    }
  }
  resource "infoblox_dhcp_optiondefinition" "test" {
    uddi = {
      code = 234
      name = "{{random4}}"
      option_space = infoblox_dhcp_optionspace.test.id
      type = "boolean"
    }
  }

//   resource "infoblox_dhcp_option_group_unknown" "test" {
//       uddi = {
//         name = "\"og-\"+optionSpace"
//         protocol = "ip4"
//       }
//   }
  PREREQ

  step {
    uddi {
      address      = "{{random_ipv4_network}}"
      cidr         = 16
      space        = infoblox_network_view.test.id
      dhcp_options = [{ type = "option", option_code = infoblox_dhcp_optiondefinition.test.id, option_value = "true" }]
    }
    check = {
      "uddi.dhcp_options.#"              = "1"
      "uddi.dhcp_options.0.type"         = "option"
      "uddi.dhcp_options.0.option_value" = "true"
    }
  }

  step {
    uddi {
      address      = "{{random_ipv4_network}}"
      cidr         = 16
      space        = infoblox_network_view.test.id
      dhcp_options = [{ type = "group", group = "{{uddi_option_group_1_id}}" }]
    }
    check = {
      "uddi.dhcp_options.#"       = "1"
      "uddi.dhcp_options.0.type"  = "group"
      "uddi.dhcp_options.0.group" = "{{uddi_option_group_1_id}}"
    }
  }

}

case "header_option_filename" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address                = "{{random_ipv4_network}}"
      cidr                   = 16
      space                  = infoblox_network_view.test.id
      header_option_filename = "testfile"
    }
    check = {
      "uddi.header_option_filename" = "testfile"
    }
  }

  step {
    uddi {
      address                = "{{random_ipv4_network}}"
      cidr                   = 16
      space                  = infoblox_network_view.test.id
      header_option_filename = "testfile1"
    }
    check = {
      "uddi.header_option_filename" = "testfile1"
    }
  }

}

case "header_option_server_address" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address                      = "{{random_ipv4_network}}"
      cidr                         = 16
      space                        = infoblox_network_view.test.id
      header_option_server_address = "1.1.1.1"
    }
    check = {
      "uddi.header_option_server_address" = "1.1.1.1"
    }
  }

  step {
    uddi {
      address                      = "{{random_ipv4_network}}"
      cidr                         = 16
      space                        = infoblox_network_view.test.id
      header_option_server_address = "2.2.2.2"
    }
    check = {
      "uddi.header_option_server_address" = "2.2.2.2"
    }
  }

}

case "header_option_server_name" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address                   = "{{random_ipv4_network}}"
      cidr                      = 16
      space                     = infoblox_network_view.test.id
      header_option_server_name = "test"
    }
    check = {
      "uddi.header_option_server_name" = "test"
    }
  }

  step {
    uddi {
      address                   = "{{random_ipv4_network}}"
      cidr                      = 16
      space                     = infoblox_network_view.test.id
      header_option_server_name = "test-1"
    }
    check = {
      "uddi.header_option_server_name" = "test-1"
    }
  }

}

case "hostname_rewrite_char" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address               = "{{random_ipv4_network}}"
      cidr                  = 16
      space                 = infoblox_network_view.test.id
      hostname_rewrite_char = "a"
    }
    check = {
      "uddi.hostname_rewrite_char" = "a"
    }
  }

  step {
    uddi {
      address               = "{{random_ipv4_network}}"
      cidr                  = 16
      space                 = infoblox_network_view.test.id
      hostname_rewrite_char = "c"
    }
    check = {
      "uddi.hostname_rewrite_char" = "c"
    }
  }

}

case "hostname_rewrite_enabled" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address                  = "{{random_ipv4_network}}"
      cidr                     = 16
      space                    = infoblox_network_view.test.id
      hostname_rewrite_enabled = true
    }
    check = {
      "uddi.hostname_rewrite_enabled" = "true"
    }
  }

  step {
    uddi {
      address                  = "{{random_ipv4_network}}"
      cidr                     = 16
      space                    = infoblox_network_view.test.id
      hostname_rewrite_enabled = false
    }
    check = {
      "uddi.hostname_rewrite_enabled" = "false"
    }
  }

}

case "hostname_rewrite_regex" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address                = "{{random_ipv4_network}}"
      cidr                   = 16
      space                  = infoblox_network_view.test.id
      hostname_rewrite_regex = "[^a-z]"
    }
    check = {
      "uddi.hostname_rewrite_regex" = "[^a-z]"
    }
  }

  step {
    uddi {
      address                = "{{random_ipv4_network}}"
      cidr                   = 16
      space                  = infoblox_network_view.test.id
      hostname_rewrite_regex = "[^g-hG-H0-9_.]"
    }
    check = {
      "uddi.hostname_rewrite_regex" = "[^g-hG-H0-9_.]"
    }
  }

}

case "inheritance_sources" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address             = "{{random_ipv4_network}}"
      cidr                = 16
      space               = infoblox_network_view.test.id
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
      address             = "{{random_ipv4_network}}"
      cidr                = 16
      space               = infoblox_network_view.test.id
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

case "federated_realms" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_federated_realm" "test" {
    uddi = {
      name = "{{random2}}"
    }
  }

  resource "infoblox_federated_realm" "test2" {
    uddi = {
      name = "{{random3}}"
    }
  }

  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address          = "{{random_ipv6_network_address}}"
      cidr             = 64
      space            = infoblox_network_view.test.id
      federated_realms = [infoblox_federated_realm.test.id]
    }
    check = {
      "uddi.federated_realms.#" = "1"
      "uddi.federated_realms.0" = infoblox_federated_realm.test.id
    }
  }

  step {
    uddi {
      address          = "{{random_ipv6_network_address}}"
      cidr             = 64
      space            = infoblox_network_view.test.id
      federated_realms = [infoblox_federated_realm.test2.id]
    }
    check = {
      "uddi.federated_realms.#" = "1"
      "uddi.federated_realms.0" = infoblox_federated_realm.test2.id
    }
  }

}

case "name" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address = "{{random_ipv4_network}}"
      cidr    = 16
      space   = infoblox_network_view.test.id
      name    = "test_name"
    }
    check = {
      "uddi.name" = "test_name"
    }
  }

  step {
    uddi {
      address = "{{random_ipv4_network}}"
      cidr    = 16
      space   = infoblox_network_view.test.id
      name    = "test_name_1"
    }
    check = {
      "uddi.name" = "test_name_1"
    }
  }

}

case "space" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "one" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address = "{{random_ipv4_network}}"
      cidr    = 16
      space   = infoblox_network_view.test.id
    }
  }

}

case "tags" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      address = "{{random_ipv4_network}}"
      cidr    = 16
      space   = infoblox_network_view.test.id
      tags    = { tag1 = "value1", tag2 = "value2" }
    }
    check = {
      "uddi.tags.tag1" = "value1"
      "uddi.tags.tag2" = "value2"
    }
  }

  step {
    uddi {
      address = "{{random_ipv4_network}}"
      cidr    = 16
      space   = infoblox_network_view.test.id
      tags    = { tag2 = "value2changed", tag3 = "value3" }
    }
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }

}

case "next_available_network" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }

  resource "infoblox_network_container" "alloc_parent" {
    uddi = {
      address = "{{random_ipv4_network}}"
      cidr    = 16
      name    = "{{random}}"
      space              = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      dynamic_allocation = {
        next_available_id = infoblox_network_container.alloc_parent.id
      }
      cidr    = 24
      comment = "Created by Dynamic Allocation"
      space   = infoblox_network_view.test.id
    }
    depends_on = [infoblox_network_container.alloc_parent]
    check = {
      "uddi.cidr"    = "24"
      "uddi.comment" = "Created by Dynamic Allocation"
    }
  }

}
