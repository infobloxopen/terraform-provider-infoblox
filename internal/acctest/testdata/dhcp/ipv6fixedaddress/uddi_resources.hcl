# Auto-generated resource acceptance-test cases for Ipv6fixedaddress.
case "basic" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_ipv6_network" "test" {
    uddi = {
      address = "2001:db8:{{random_hextet}}:{{random_int}}::"
      cidr = 64
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type  = "mac"
      match_value = "aa:aa:aa:aa:aa:aa"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.address"     = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      "uddi.match_type"  = "mac"
      "uddi.match_value" = "aa:aa:aa:aa:aa:aa"
    }
  }

}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_ipv6_network" "test" {
    uddi = {
      address = "2001:db8:{{random_hextet}}:{{random_int}}::"
      cidr = 64
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type  = "mac"
      match_value = "aa:aa:aa:aa:aa:aa"
    }
    depends_on = [infoblox_ipv6_network.test]
  }

}

case "address" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_ipv6_network" "test" {
    uddi = {
      address = "2001:db8:{{random_hextet}}:{{random_int}}::"
      cidr = 64
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type  = "mac"
      match_value = "aa:aa:aa:aa:aa:aa"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.address" = "2001:db8:{{random_hextet}}:{{random_int}}::10"
    }
  }

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = "2001:db8:{{random_hextet}}:{{random_int}}::11"
      match_type  = "mac"
      match_value = "bb:bb:bb:bb:bb:bb"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.address" = "2001:db8:{{random_hextet}}:{{random_int}}::11"
    }
  }

  # step {
  #   uddi {
  #     ip_space    = infoblox_network_view.test.id
  #     match_type  = "mac"
  #     match_value = "cc:cc:cc:cc:cc:cc"
  #   }
  #   depends_on = [infoblox_ipv6_network.test]
  #   check = {
  #     "uddi.address" = "2001:db8:{{random_hextet}}:{{random_int}}::1"
  #   }
  # }

}

case "comment" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_ipv6_network" "test" {
    uddi = {
      address = "2001:db8:{{random_hextet}}:{{random_int}}::"
      cidr = 64
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type  = "mac"
      match_value = "aa:aa:aa:aa:aa:aa"
      comment     = "this range is created by terraform"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.comment" = "this range is created by terraform"
    }
  }

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type  = "mac"
      match_value = "aa:aa:aa:aa:aa:aa"
      comment     = "update: this range is created by terraform"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.comment" = "update: this range is created by terraform"
    }
  }

}

case "disable_dhcp" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_ipv6_network" "test" {
    uddi = {
      address = "2001:db8:{{random_hextet}}:{{random_int}}::"
      cidr = 64
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      ip_space     = infoblox_network_view.test.id
      address      = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type   = "mac"
      match_value  = "aa:aa:aa:aa:aa:aa"
      disable_dhcp = false
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.disable_dhcp" = "false"
    }
  }

  step {
    uddi {
      ip_space     = infoblox_network_view.test.id
      address      = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type   = "mac"
      match_value  = "aa:aa:aa:aa:aa:aa"
      disable_dhcp = true
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.disable_dhcp" = "true"
    }
  }

}

# TODO: auto-extraction incomplete — please verify and fill in manually.
# Reason: requires_resource: infoblox_dhcp_option_group not yet implemented
case "dhcp_options" {
  backend     = "uddi"
  skip        = true
  skip_reason = "requires_resource: infoblox_dhcp_option_group not yet implemented"
  parallel    = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random3}}"
    }
  }
  resource "infoblox_ipv6_network" "test" {
    uddi = {
      address = "2001:db8:{{random_hextet}}:{{random_int}}::"
      cidr = 64
      space = infoblox_network_view.test.id
    }
  }
  resource "infoblox_dhcp_optiondefinition" "test" {
    uddi = {
      code = 234
      name = "test_dhcp_option_code"
      option_space = infoblox_dhcp_optionspace.test.id
      type = "boolean"
    }
  }
  resource "infoblox_dhcp_option_group_unknown" "test" {
    uddi = {
      name = "og-{{random}}"
      protocol = "ip4"
    }
  }
  resource "infoblox_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random}}"
      protocol = "ip4"
    }
  }
  PREREQ

  step {
    uddi {
      ip_space     = infoblox_network_view.test.id
      address      = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type   = "mac"
      match_value  = "aa:aa:aa:aa:aa:aa"
      name         = "{{random2}}"
      dhcp_options = [{ type = "option", option_code = infoblox_dhcp_optiondefinition.test.id, option_value = true }]
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.dhcp_options.#"              = "1"
      "uddi.dhcp_options.0.option_value" = "true"
    }
  }

  step {
    uddi {
      ip_space     = infoblox_network_view.test.id
      address      = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type   = "mac"
      match_value  = "aa:aa:aa:aa:aa:aa"
      name         = "{{random2}}"
      dhcp_options = [{ type = "group", group = infoblox_dhcp_option_group_unknown.test.id }]
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.dhcp_options.#" = "1"
    }
  }

}

case "header_option_filename" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_ipv6_network" "test" {
    uddi = {
      address = "2001:db8:{{random_hextet}}:{{random_int}}::"
      cidr = 64
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      ip_space               = infoblox_network_view.test.id
      address                = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type             = "mac"
      match_value            = "aa:aa:aa:aa:aa:aa"
      header_option_filename = "header_option_filename"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.header_option_filename" = "header_option_filename"
    }
  }

  step {
    uddi {
      ip_space               = infoblox_network_view.test.id
      address                = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type             = "mac"
      match_value            = "aa:aa:aa:aa:aa:aa"
      header_option_filename = "header_option_filename_update"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.header_option_filename" = "header_option_filename_update"
    }
  }

}

case "header_option_server_address" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_ipv6_network" "test" {
    uddi = {
      address = "2001:db8:{{random_hextet}}:{{random_int}}::"
      cidr = 64
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      ip_space                     = infoblox_network_view.test.id
      address                      = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type                   = "mac"
      match_value                  = "aa:aa:aa:aa:aa:aa"
      header_option_server_address = "2001:db8:{{random_hextet}}:{{random_int}}::12"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.header_option_server_address" = "2001:db8:{{random_hextet}}:{{random_int}}::12"
    }
  }

  step {
    uddi {
      ip_space                     = infoblox_network_view.test.id
      address                      = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type                   = "mac"
      match_value                  = "aa:aa:aa:aa:aa:aa"
      header_option_server_address = "2001:db8:{{random_hextet}}:{{random_int}}::13"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.header_option_server_address" = "2001:db8:{{random_hextet}}:{{random_int}}::13"
    }
  }

}

case "header_option_server_name" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_ipv6_network" "test" {
    uddi = {
      address = "2001:db8:{{random_hextet}}:{{random_int}}::"
      cidr = 64
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      ip_space                  = infoblox_network_view.test.id
      address                   = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type                = "mac"
      match_value               = "aa:aa:aa:aa:aa:aa"
      header_option_server_name = "header_option_server_name"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.header_option_server_name" = "header_option_server_name"
    }
  }

  step {
    uddi {
      ip_space                  = infoblox_network_view.test.id
      address                   = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type                = "mac"
      match_value               = "aa:aa:aa:aa:aa:aa"
      header_option_server_name = "header_option_server_name_update"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.header_option_server_name" = "header_option_server_name_update"
    }
  }

}

case "hostname" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_ipv6_network" "test" {
    uddi = {
      address = "2001:db8:{{random_hextet}}:{{random_int}}::"
      cidr = 64
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type  = "mac"
      match_value = "aa:aa:aa:aa:aa:aa"
      hostname    = "hostname1"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.hostname" = "hostname1"
    }
  }

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type  = "mac"
      match_value = "aa:aa:aa:aa:aa:aa"
      hostname    = "hostname2"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.hostname" = "hostname2"
    }
  }

}

case "inheritance_sources" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_ipv6_network" "test" {
    uddi = {
      address = "2001:db8:{{random_hextet}}:{{random_int}}::"
      cidr = 64
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      ip_space                     = infoblox_network_view.test.id
      address                      = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type                   = "mac"
      match_value                  = "aa:aa:aa:aa:aa:aa"
      inheritance_sources          = { header_option_filename = { action = "inherit" }, header_option_server_address = { action = "inherit" }, header_option_server_name = { action = "inherit" } }
      header_option_filename       = "header_option_filename"
      header_option_server_address = "2001:db8:{{random_hextet}}:{{random_int}}::12"
      header_option_server_name    = "header_option_server_name"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.inheritance_sources.header_option_filename.action"       = "inherit"
      "uddi.inheritance_sources.header_option_server_address.action" = "inherit"
      "uddi.inheritance_sources.header_option_server_name.action"    = "inherit"
    }
  }

  step {
    uddi {
      ip_space                     = infoblox_network_view.test.id
      address                      = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type                   = "mac"
      match_value                  = "aa:aa:aa:aa:aa:aa"
      inheritance_sources          = { header_option_filename = { action = "override" }, header_option_server_address = { action = "override" }, header_option_server_name = { action = "override" } }
      header_option_filename       = "header_option_filename"
      header_option_server_address = "2001:db8:{{random_hextet}}:{{random_int}}::12"
      header_option_server_name    = "header_option_server_name"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.inheritance_sources.header_option_filename.action"       = "override"
      "uddi.inheritance_sources.header_option_server_address.action" = "override"
      "uddi.inheritance_sources.header_option_server_name.action"    = "override"
    }
  }

}

case "ip_space" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "one" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_ipv6_network" "test" {
    uddi = {
      address = "2001:db8:{{random_hextet}}:{{random_int}}::"
      cidr = 64
      space = infoblox_network_view.one.id
    }
  }
  PREREQ

  step {
    uddi {
      address     = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type  = "mac"
      match_value = "aa:aa:aa:aa:aa:aa"
      ip_space    = infoblox_network_view.one.id
    }
    depends_on = [infoblox_ipv6_network.test]
  }

  step {
    uddi {
      address     = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type  = "mac"
      match_value = "aa:aa:aa:aa:aa:aa"
      ip_space    = infoblox_network_view.two.id
    }
    depends_on = [infoblox_ipv6_network.test]
  }

}

case "match_type_and_match_value" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_ipv6_network" "test" {
    uddi = {
      address = "2001:db8:{{random_hextet}}:{{random_int}}::"
      cidr = 64
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type  = "client_hex"
      match_value = "aa"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.match_type"  = "client_hex"
      "uddi.match_value" = "aa"
    }
  }

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type  = "client_hex"
      match_value = "bb"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.match_type"  = "client_hex"
      "uddi.match_value" = "bb"
    }
  }

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type  = "client_text"
      match_value = "clienttext"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.match_type"  = "client_text"
      "uddi.match_value" = "clienttext"
    }
  }

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type  = "mac"
      match_value = "aa:aa:aa:aa:aa:aa"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.match_type"  = "mac"
      "uddi.match_value" = "aa:aa:aa:aa:aa:aa"
    }
  }

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type  = "relay_hex"
      match_value = "aa"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.match_type"  = "relay_hex"
      "uddi.match_value" = "aa"
    }
  }

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type  = "relay_text"
      match_value = "relaytext"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.match_type"  = "relay_text"
      "uddi.match_value" = "relaytext"
    }
  }

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type  = "client_hex"
      match_value = "aa"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.match_type"  = "client_hex"
      "uddi.match_value" = "aa"
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_ipv6_network" "test" {
    uddi = {
      address = "2001:db8:{{random_hextet}}:{{random_int}}::"
      cidr = 64
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type  = "mac"
      match_value = "aa:aa:aa:aa:aa:aa"
      name        = "example_fixed_address"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.name" = "example_fixed_address"
    }
  }

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type  = "mac"
      match_value = "aa:aa:aa:aa:aa:aa"
      name        = "example_fixed_address_updated"
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.name" = "example_fixed_address_updated"
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_ipv6_network" "test" {
    uddi = {
      address = "2001:db8:{{random_hextet}}:{{random_int}}::"
      cidr = 64
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type  = "mac"
      match_value = "aa:aa:aa:aa:aa:aa"
      tags        = { tag1 = "value1", tag2 = "value2" }
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.tags.tag1" = "value1"
      "uddi.tags.tag2" = "value2"
    }
  }

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = "2001:db8:{{random_hextet}}:{{random_int}}::10"
      match_type  = "mac"
      match_value = "aa:aa:aa:aa:aa:aa"
      tags        = { tag2 = "value2changed", tag3 = "value3" }
    }
    depends_on = [infoblox_ipv6_network.test]
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }

}

case "next_available_id_count" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_ipv6_network" "test" {
    uddi = {
      address = "2001:db8:{{random_hextet}}:{{random_int}}::"
      cidr = 64
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      match_type  = "mac"
      match_value = "aa:aa:aa:aa:aa:$${count.index + 10}"
    }
  }

}
