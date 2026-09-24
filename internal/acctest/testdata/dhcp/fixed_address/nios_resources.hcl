//TODO : OBJECTS TO BE PRESENT IN GRID FOR TESTS
// - Ipv4 Network - 15.0.0.0/24 , 16.0.0.0/24
// - Logic Filter Rules - example-mac-filter-1 , example-option-filter-1
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.1"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
    }
    check = {
      "nios.ipv4addr"                        = "15.0.0.1"
      "nios.match_client"                    = "CIRCUIT_ID"
      "nios.allow_telnet"                    = "false"
      "nios.always_update_dns"               = "false"
      "nios.client_identifier_prepend_zero"  = "false"
      "nios.deny_bootp"                      = "false"
      "nios.disable"                         = "false"
      "nios.disable_discovery"               = "false"
      "nios.enable_ddns"                     = "false"
      "nios.enable_pxe_lease_time"           = "false"
      "nios.ignore_dhcp_option_list_request" = "false"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

  step {
    nios {
      ipv4addr         = "15.0.0.2"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
    }
  }

}

case "agent_circuit_id" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.3"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "30"
    }
    check = {
      "nios.agent_circuit_id" = "30"
    }
  }

  step {
    nios {
      ipv4addr         = "15.0.0.3"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "32"
    }
    check = {
      "nios.agent_circuit_id" = "32"
    }
  }

  step {
    nios {
      ipv4addr         = "15.0.0.3"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "35"
      agent_remote_id  = "34"
    }
    check = {
      "nios.agent_circuit_id" = "35"
      "nios.agent_remote_id"  = "34"
    }
  }

}

case "agent_remote_id" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr        = "15.0.0.4"
      match_client    = "REMOTE_ID"
      agent_remote_id = "{{random_int}}"
    }
    check = {
      "nios.agent_remote_id" = "{{random_int}}"
    }
  }

  step {
    nios {
      ipv4addr        = "15.0.0.4"
      match_client    = "REMOTE_ID"
      agent_remote_id = "{{random_int2}}"
    }
    check = {
      "nios.agent_remote_id" = "{{random_int2}}"
    }
  }

  step {
    nios {
      ipv4addr         = "15.0.0.4"
      match_client     = "REMOTE_ID"
      agent_circuit_id = "{{random_int4}}"
      agent_remote_id  = "{{random_int3}}"
    }
    check = {
      "nios.agent_remote_id" = "{{random_int3}}"
    }
  }

}

case "allow_telnet" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.5"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      allow_telnet     = true
      cli_credentials  = [{ comment = "Comment for SSH Credentials", user = "NIOS_USER", password = "NIOS_PASSWORD", credential_type = "SSH", credential_group = "default" }, { user = "NIOS_USER", comment = "Comment for SSH Credentials", password = "NIOS_PASSWORD", credential_type = "TELNET", credential_group = "default" }]
      comment          = "Comment for SSH Credentials"
    }
    check = {
      "nios.allow_telnet" = "true"
    }
  }

}

case "always_update_dns" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr          = "15.0.0.6"
      match_client      = "CIRCUIT_ID"
      agent_circuit_id  = "{{random_int}}"
      always_update_dns = true
    }
    check = {
      "nios.always_update_dns" = "true"
    }
  }

  step {
    nios {
      ipv4addr          = "15.0.0.6"
      match_client      = "CIRCUIT_ID"
      agent_circuit_id  = "{{random_int}}"
      always_update_dns = false
    }
    check = {
      "nios.always_update_dns" = "false"
    }
  }

}

case "bootfile" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.7"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      bootfile         = "file"
    }
    check = {
      "nios.bootfile" = "file"
    }
  }

  step {
    nios {
      ipv4addr         = "15.0.0.7"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      bootfile         = "file1"
    }
    check = {
      "nios.bootfile" = "file1"
    }
  }

}

case "bootserver" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.8"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      bootserver       = "boot_server_example.com"
    }
    check = {
      "nios.bootserver" = "boot_server_example.com"
    }
  }

  step {
    nios {
      ipv4addr         = "15.0.0.8"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      bootserver       = "boot_server_updated_example.com"
    }
    check = {
      "nios.bootserver" = "boot_server_updated_example.com"
    }
  }

}

case "cli_credentials" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.9"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      cli_credentials  = [{ comment = "Comment for CLI Credentials", user = "NIOS_USER", password = "NIOS_PASSWORD", credential_type = "SSH", credential_group = "default" }]
    }
    check = {
      "nios.cli_credentials.0.comment"          = "Comment for CLI Credentials"
      "nios.cli_credentials.0.user"             = "NIOS_USER"
      "nios.cli_credentials.0.credential_type"  = "SSH"
      "nios.cli_credentials.0.credential_group" = "default"
    }
  }

  # Update: add TELNET credential at index 1
  step {
    nios {
      ipv4addr         = "15.0.0.9"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      cli_credentials  = [{ comment = "Comment for CLI Credentials", user = "NIOS_USER", password = "NIOS_PASSWORD", credential_type = "SSH", credential_group = "default" }, { comment = "Updated Comment for CLI Credentials", user = "NIOS_USER", password = "NIOS_PASSWORD", credential_type = "TELNET", credential_group = "default" }]
    }
    check = {
      "nios.cli_credentials.1.comment"          = "Updated Comment for CLI Credentials"
      "nios.cli_credentials.1.user"             = "NIOS_USER"
      "nios.cli_credentials.1.credential_type"  = "TELNET"
      "nios.cli_credentials.1.credential_group" = "default"
    }
  }

  # Update: change credential_type at index 1 to ENABLE_TELNET
  step {
    nios {
      ipv4addr         = "15.0.0.9"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      cli_credentials  = [{ comment = "Comment for CLI Credentials", user = "NIOS_USER", password = "NIOS_PASSWORD", credential_type = "SSH", credential_group = "default" }, { comment = "Updated Comment for CLI Credentials", user = "NIOS_USER", password = "NIOS_PASSWORD", credential_type = "ENABLE_TELNET", credential_group = "default" }]
    }
    check = {
      "nios.cli_credentials.1.comment"          = "Updated Comment for CLI Credentials"
      "nios.cli_credentials.1.user"             = "NIOS_USER"
      "nios.cli_credentials.1.credential_type"  = "ENABLE_TELNET"
      "nios.cli_credentials.1.credential_group" = "default"
    }
  }

  # Update: change credential_type at index 1 to ENABLE_SSH
  step {
    nios {
      ipv4addr         = "15.0.0.9"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      cli_credentials  = [{ comment = "Comment for CLI Credentials", user = "NIOS_USER", password = "NIOS_PASSWORD", credential_type = "SSH", credential_group = "default" }, { comment = "Updated Comment for CLI Credentials", user = "NIOS_USER", password = "NIOS_PASSWORD", credential_type = "ENABLE_SSH", credential_group = "default" }]
    }
    check = {
      "nios.cli_credentials.1.comment"          = "Updated Comment for CLI Credentials"
      "nios.cli_credentials.1.user"             = "NIOS_USER"
      "nios.cli_credentials.1.credential_type"  = "ENABLE_SSH"
      "nios.cli_credentials.1.credential_group" = "default"
    }
  }

  # Create without cli_credentials
  step {
    nios {
      ipv4addr         = "15.0.0.150"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
    }
    check = {
      "nios.cli_credentials.#" = "0"
    }
  }

  # Add cli_credentials (cliCred)
  step {
    nios {
      ipv4addr         = "15.0.0.151"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      cli_credentials  = [{ user = "user1", credential_type = "SSH", comment = "cli credential comment", password = "password1", credential_group = "default" }]
    }
    check = {
      "nios.cli_credentials.0.credential_type"  = "SSH"
      "nios.cli_credentials.0.user"             = "user1"
      "nios.cli_credentials.0.credential_group" = "default"
    }
  }

  # Update write-only password field (cliCred1)
  step {
    nios {
      ipv4addr         = "15.0.0.151"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      cli_credentials  = [{ user = "user1", credential_type = "SSH", comment = "cli credential comment", password = "password12", credential_group = "default" }]
    }
    check = {
      "nios.cli_credentials.0.comment"          = "cli credential comment"
      "nios.cli_credentials.0.user"             = "user1"
      "nios.cli_credentials.0.credential_type"  = "SSH"
      "nios.cli_credentials.0.credential_group" = "default"
    }
  }

  # Update non write-only user field (cliCred2)
  step {
    nios {
      ipv4addr         = "15.0.0.151"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      cli_credentials  = [{ user = "user2", credential_type = "SSH", comment = "cli credential comment update", password = "password12", credential_group = "default" }]
    }
    check = {
      "nios.cli_credentials.0.comment"          = "cli credential comment update"
      "nios.cli_credentials.0.user"             = "user2"
      "nios.cli_credentials.0.credential_type"  = "SSH"
      "nios.cli_credentials.0.credential_group" = "default"
    }
  }

}

case "client_identifier_prepend_zero" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr                       = "15.0.0.10"
      match_client                   = "CIRCUIT_ID"
      agent_circuit_id               = "{{random_int}}"
      client_identifier_prepend_zero = true
    }
    check = {
      "nios.client_identifier_prepend_zero" = "true"
    }
  }

  step {
    nios {
      ipv4addr                       = "15.0.0.10"
      match_client                   = "CIRCUIT_ID"
      agent_circuit_id               = "{{random_int}}"
      client_identifier_prepend_zero = false
    }
    check = {
      "nios.client_identifier_prepend_zero" = "false"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.11"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      comment          = "Comment for Fixed Address"
    }
    check = {
      "nios.comment" = "Comment for Fixed Address"
    }
  }

  step {
    nios {
      ipv4addr         = "15.0.0.11"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      comment          = "Updated Comment for Fixed Address"
    }
    check = {
      "nios.comment" = "Updated Comment for Fixed Address"
    }
  }

}

case "ddns_domainname" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.12"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      ddns_domainname  = "ddns_domain.name"
    }
    check = {
      "nios.ddns_domainname" = "ddns_domain.name"
    }
  }

  step {
    nios {
      ipv4addr         = "15.0.0.12"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      ddns_domainname  = "updated_ddns_domain.name"
    }
    check = {
      "nios.ddns_domainname" = "updated_ddns_domain.name"
    }
  }

}

case "ddns_hostname" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.13"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      ddns_hostname    = "ddns_host.name"
    }
    check = {
      "nios.ddns_hostname" = "ddns_host.name"
    }
  }

  step {
    nios {
      ipv4addr         = "15.0.0.13"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      ddns_hostname    = "updated_ddns_host.name"
    }
    check = {
      "nios.ddns_hostname" = "updated_ddns_host.name"
    }
  }

}

case "deny_bootp" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.14"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      deny_bootp       = true
    }
    check = {
      "nios.deny_bootp" = "true"
    }
  }

  step {
    nios {
      ipv4addr         = "15.0.0.14"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      deny_bootp       = false
    }
    check = {
      "nios.deny_bootp" = "false"
    }
  }

}

case "device_description" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr           = "15.0.0.15"
      match_client       = "CIRCUIT_ID"
      agent_circuit_id   = "{{random_int}}"
      device_description = "DEVICE_DESCRIPTION"
    }
    check = {
      "nios.device_description" = "DEVICE_DESCRIPTION"
    }
  }

  step {
    nios {
      ipv4addr           = "15.0.0.15"
      match_client       = "CIRCUIT_ID"
      agent_circuit_id   = "{{random_int}}"
      device_description = "DEVICE_DESCRIPTION_UPDATED"
    }
    check = {
      "nios.device_description" = "DEVICE_DESCRIPTION_UPDATED"
    }
  }

}

case "device_location" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.16"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      device_location  = "DEVICE_LOCATION"
    }
    check = {
      "nios.device_location" = "DEVICE_LOCATION"
    }
  }

  step {
    nios {
      ipv4addr         = "15.0.0.16"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      device_location  = "DEVICE_LOCATION_UPDATED"
    }
    check = {
      "nios.device_location" = "DEVICE_LOCATION_UPDATED"
    }
  }

}

case "device_type" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.17"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      device_type      = "DEVICE_TYPE"
    }
    check = {
      "nios.device_type" = "DEVICE_TYPE"
    }
  }

  step {
    nios {
      ipv4addr         = "15.0.0.17"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      device_type      = "DEVICE_TYPE_UPDATED"
    }
    check = {
      "nios.device_type" = "DEVICE_TYPE_UPDATED"
    }
  }

}

case "device_vendor" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.18"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      device_vendor    = "DEVICE_VENDOR"
    }
    check = {
      "nios.device_vendor" = "DEVICE_VENDOR"
    }
  }

  step {
    nios {
      ipv4addr         = "15.0.0.18"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      device_vendor    = "DEVICE_VENDOR_UPDATED"
    }
    check = {
      "nios.device_vendor" = "DEVICE_VENDOR_UPDATED"
    }
  }

}

case "dhcp_client_identifier" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr               = "15.0.0.19"
      match_client           = "CLIENT_ID"
      dhcp_client_identifier = "DHCP_CLIENT_IDENTIFIER"
    }
    check = {
      "nios.dhcp_client_identifier" = "DHCP_CLIENT_IDENTIFIER"
    }
  }

  step {
    nios {
      ipv4addr               = "15.0.0.19"
      match_client           = "CLIENT_ID"
      dhcp_client_identifier = "DHCP_CLIENT_IDENTIFIER_UPDATED"
    }
    check = {
      "nios.dhcp_client_identifier" = "DHCP_CLIENT_IDENTIFIER_UPDATED"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.20"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      disable          = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

  step {
    nios {
      ipv4addr         = "15.0.0.20"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      disable          = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

}

case "disable_discovery" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr          = "15.0.0.21"
      match_client      = "CIRCUIT_ID"
      agent_circuit_id  = "{{random_int}}"
      disable_discovery = true
    }
    check = {
      "nios.disable_discovery" = "true"
    }
  }

  step {
    nios {
      ipv4addr          = "15.0.0.21"
      match_client      = "CIRCUIT_ID"
      agent_circuit_id  = "{{random_int}}"
      disable_discovery = false
    }
    check = {
      "nios.disable_discovery" = "false"
    }
  }

}

case "enable_ddns" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.22"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      enable_ddns      = true
    }
    check = {
      "nios.enable_ddns" = "true"
    }
  }

  step {
    nios {
      ipv4addr         = "15.0.0.22"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      enable_ddns      = false
    }
    check = {
      "nios.enable_ddns" = "false"
    }
  }

}

case "enable_immediate_discovery" {
  backend  = "nios"
  parallel = true
  skip                  = true
  skip_reason           = "Discovery Member is Required for the Test"

  step {
    nios {
      ipv4addr                   = "15.0.0.23"
      match_client               = "CIRCUIT_ID"
      agent_circuit_id           = "{{random_int}}"
      enable_immediate_discovery = true
    }
    check = {
      "nios.enable_immediate_discovery" = "true"
    }
  }

  step {
    nios {
      ipv4addr                   = "15.0.0.23"
      match_client               = "CIRCUIT_ID"
      agent_circuit_id           = "{{random_int}}"
      enable_immediate_discovery = false
    }
    check = {
      "nios.enable_immediate_discovery" = "false"
    }
  }

}

case "enable_pxe_lease_time" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr              = "15.0.0.24"
      match_client          = "CIRCUIT_ID"
      agent_circuit_id      = "{{random_int}}"
      enable_pxe_lease_time = true
      pxe_lease_time        = 3600
    }
    check = {
      "nios.enable_pxe_lease_time" = "true"
    }
  }

  step {
    nios {
      ipv4addr              = "15.0.0.24"
      match_client          = "CIRCUIT_ID"
      agent_circuit_id      = "{{random_int}}"
      enable_pxe_lease_time = false
      pxe_lease_time        = 3600
    }
    check = {
      "nios.enable_pxe_lease_time" = "false"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.25"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      ext_attrs        = { Site = "{{random}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random}}"
    }
  }

  step {
    nios {
      ipv4addr         = "15.0.0.25"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      ext_attrs        = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

}

case "ignore_dhcp_option_list_request" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr                        = "15.0.0.26"
      match_client                    = "CIRCUIT_ID"
      agent_circuit_id                = "{{random_int}}"
      ignore_dhcp_option_list_request = true
    }
    check = {
      "nios.ignore_dhcp_option_list_request" = "true"
    }
  }

  step {
    nios {
      ipv4addr                        = "15.0.0.26"
      match_client                    = "CIRCUIT_ID"
      agent_circuit_id                = "{{random_int}}"
      ignore_dhcp_option_list_request = false
    }
    check = {
      "nios.ignore_dhcp_option_list_request" = "false"
    }
  }

}

case "ipv4addr" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.27"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
    }
    check = {
      "nios.ipv4addr" = "15.0.0.27"
    }
  }

  step {
    nios {
      ipv4addr         = "15.0.0.60"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
    }
    check = {
      "nios.ipv4addr" = "15.0.0.60"
    }
  }

}

case "func_call" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_network" "test_func_call" {
      nios = {
        network = "{{random_cidr_network}}"
      }
    }

    PREREQ

  step {
    nios {
      match_client       = "CIRCUIT_ID"
      agent_circuit_id   = "{{random_int}}"
      dynamic_allocation = { network = infoblox_network.test_func_call.nios.network, network_view = "default" }
      comment            = "next_available_ip"
    }
  }

  step {
    nios {
      match_client       = "CIRCUIT_ID"
      agent_circuit_id   = "{{random_int}}"
      dynamic_allocation = { network = infoblox_network.test_func_call.nios.network, network_view = "default" }
      comment            = "next_available_ip"
    }
  }

}

case "logic_filter_rules" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr           = "15.0.0.28"
      match_client       = "CIRCUIT_ID"
      agent_circuit_id   = "{{random_int}}"
      logic_filter_rules = [{ filter = "example-mac-filter-1", type = "MAC" }]
    }
    check = {
      "nios.logic_filter_rules.0.filter" = "example-mac-filter-1"
      "nios.logic_filter_rules.0.type"   = "MAC"
    }
  }

  step {
    nios {
      ipv4addr           = "15.0.0.28"
      match_client       = "CIRCUIT_ID"
      agent_circuit_id   = "{{random_int}}"
      logic_filter_rules = [{ filter = "example-option-filter-1", type = "Option" }]
    }
    check = {
      "nios.logic_filter_rules.0.filter" = "example-option-filter-1"
      "nios.logic_filter_rules.0.type"   = "Option"
    }
  }

}

case "mac" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr     = "15.0.0.29"
      match_client = "MAC_ADDRESS"
      mac          = "00:1a:2b:3c:4d:5e"
    }
    check = {
      "nios.mac" = "00:1a:2b:3c:4d:5e"
    }
  }

  step {
    nios {
      ipv4addr     = "15.0.0.29"
      match_client = "MAC_ADDRESS"
      mac          = "10:9a:dd:ee:ff:01"
    }
    check = {
      "nios.mac" = "10:9a:dd:ee:ff:01"
    }
  }

  step {
    nios {
      ipv4addr     = "15.0.0.29"
      match_client = "RESERVED"
      mac          = "00:00:00:00:00:00"
    }
    check = {
      "nios.match_client" = "RESERVED"
      "nios.mac"          = "00:00:00:00:00:00"
    }
  }

}

case "match_client" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.30"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
    }
    check = {
      "nios.match_client" = "CIRCUIT_ID"
    }
  }

  step {
    nios {
      ipv4addr        = "15.0.0.30"
      match_client    = "REMOTE_ID"
      agent_remote_id = "{{random_int2}}"
    }
    check = {
      "nios.match_client" = "REMOTE_ID"
    }
  }

}

case "ms_options" {
  backend     = "nios"
  skip        = true
  skip_reason = "Skipping test as appropriate MS options are not set up in the GRID"
  parallel    = true

  step {
    nios {
      ipv4addr         = "15.0.0.31"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      ms_options       = [{ name = "time-offset", num = 2, value = 50 }]
    }
    check = {
      "nios.ms_options" = "50"
    }
  }

  step {
    nios {
      ipv4addr         = "15.0.0.31"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      ms_options       = [{ name = "dhcp-lease-time", num = 56, value = 100 }]
    }
    check = {
      "nios.ms_options" = "100"
    }
  }

}

case "ms_server" {
  backend  = "nios"
  parallel = true
  skip        = true
  skip_reason = "Skipping this test as it requires Range to be supported"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network" "example_network" {
    nios = {
      network = "150.0.0.0/24"
      network_view = "default"
      comment = "Created by Terraform for FixedAddress MS Server Test"
      members = [{ struct = "msdhcpserver", ipv4addr = "10.10.10.10" }]
    }
  }
  resource "infoblox_dhcp_range" "test_ms_server_range" {
    nios = {
      start_addr = "150.0.0.35"
      end_addr = "150.0.0.40"
      ms_server = { ipv4addr = infoblox_network.example_network.nios.members[0].ipv4addr }
      server_association_type = "MS_SERVER"
    }
  }
  PREREQ

  step {
    nios {
      ipv4addr     = "150.0.0.32"
      match_client = "MAC_ADDRESS"
      mac          = "00:1a:6b:3c:4d:5e"
      ms_server    = { ipv4addr = infoblox_network.example_network.nios.members[0].ipv4addr }
      network_view = "default"
      name         = "example_fixed_address"
    }
    depends_on = [infoblox_dhcp_range.test_ms_server_range]
    check = {
      "nios.ms_server.ipv4addr" = "10.10.10.10"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.33"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      name             = "example_fixed_address"
    }
    check = {
      "nios.name" = "example_fixed_address"
    }
  }

  step {
    nios {
      ipv4addr         = "15.0.0.33"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      name             = "example_updated_fixed_address"
    }
    check = {
      "nios.name" = "example_updated_fixed_address"
    }
  }

}

case "network" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.34"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      network          = "15.0.0.0/24"
    }
    check = {
      "nios.network" = "15.0.0.0/24"
    }
  }

  step {
    nios {
      ipv4addr         = "16.0.0.224"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      network          = "16.0.0.0/24"
    }
    check = {
      "nios.network" = "16.0.0.0/24"
    }
  }

}

case "network_view" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.35"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      network_view     = "default"
    }
    check = {
      "nios.network_view" = "default"
    }
  }

}

case "nextserver" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.36"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      nextserver       = "example_next_server.com"
    }
    check = {
      "nios.nextserver" = "example_next_server.com"
    }
  }

  step {
    nios {
      ipv4addr         = "15.0.0.36"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      nextserver       = "example_updated_next_server.com"
    }
    check = {
      "nios.nextserver" = "example_updated_next_server.com"
    }
  }

}

case "options" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.237"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      options          = [{ name = "time-offset", num = 2, value = "50" }, { name = "subnet-mask", value = "1.1.1.1" }]
    }
    check = {
      "nios.options.0.name"  = "time-offset"
      "nios.options.0.num"   = "2"
      "nios.options.0.value" = "50"
      "nios.options.1.name"  = "subnet-mask"
      "nios.options.1.value" = "1.1.1.1"
    }
  }

  step {
    nios {
      ipv4addr         = "15.0.0.237"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      options          = [{ num = 51, value = "7200" }, { name = "subnet-mask", value = "1.1.1.1" }]
    }
    check = {
      "nios.options.0.name"  = "dhcp-lease-time"
      "nios.options.0.value" = "7200"
      "nios.options.1.name"  = "subnet-mask"
      "nios.options.1.value" = "1.1.1.1"
    }
  }

}

case "pxe_lease_time" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.38"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      pxe_lease_time   = 3600
    }
    check = {
      "nios.pxe_lease_time" = "3600"
    }
  }

  step {
    nios {
      ipv4addr         = "15.0.0.38"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      pxe_lease_time   = 4800
    }
    check = {
      "nios.pxe_lease_time" = "4800"
    }
  }

}

case "reserved_interface" {
  backend     = "nios"
  skip        = true
  skip_reason = "t.Skip: Skipping test as reserved_interface is not implemented yet"
  parallel    = true

  step {
    nios {
      ipv4addr           = "15.0.0.39"
      match_client       = "CIRCUIT_ID"
      agent_circuit_id   = "{{random_int}}"
      reserved_interface = "ref"
    }
    check = {
      "nios.reserved_interface" = "ref"
    }
  }

  step {
    nios {
      ipv4addr           = "15.0.0.39"
      match_client       = "CIRCUIT_ID"
      agent_circuit_id   = "{{random_int}}"
      reserved_interface = "ref2"
    }
    check = {
      "nios.reserved_interface" = "ref2"
    }
  }

}

case "restart_if_needed" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr          = "15.0.0.40"
      match_client      = "CIRCUIT_ID"
      agent_circuit_id  = "{{random_int}}"
      restart_if_needed = true
    }
    check = {
      "nios.restart_if_needed" = "true"
    }
  }

  step {
    nios {
      ipv4addr          = "15.0.0.40"
      match_client      = "CIRCUIT_ID"
      agent_circuit_id  = "{{random_int}}"
      restart_if_needed = false
    }
    check = {
      "nios.restart_if_needed" = "false"
    }
  }

}

case "snmp3_credential" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "16.0.0.121"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      snmp3_credential = { user = "snmp", authentication_protocol = "MD5", authentication_password = "snmp1234", privacy_protocol = "3DES", privacy_password = "snmp1234", comment = "SNMP3 Credential Comment", credential_group = "default" }
    }
    check = {
      "nios.snmp3_credential.user"                    = "snmp"
      "nios.snmp3_credential.authentication_protocol" = "MD5"
      "nios.snmp3_credential.privacy_protocol"        = "3DES"
      "nios.snmp3_credential.comment"                 = "SNMP3 Credential Comment"
      "nios.snmp3_credential.credential_group"        = "default"
    }
  }

  step {
    nios {
      ipv4addr         = "16.0.0.121"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      snmp3_credential = { user = "SNMP3_USER_UPDATE", authentication_protocol = "SHA-224", authentication_password = "AUTH_PASSWORD", privacy_protocol = "AES-256", privacy_password = "PRIVACY_PASSWORD", comment = "SNMP3 Credential Comment Updated", credential_group = "default" }
    }
    check = {
      "nios.snmp3_credential.user"                    = "SNMP3_USER_UPDATE"
      "nios.snmp3_credential.authentication_protocol" = "SHA-224"
      "nios.snmp3_credential.privacy_protocol"        = "AES-256"
      "nios.snmp3_credential.comment"                 = "SNMP3 Credential Comment Updated"
      "nios.snmp3_credential.credential_group"        = "default"
    }
  }

  # Add SNMP3 credentials
  step {
    nios {
      ipv4addr         = "16.0.0.132"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int2}}"
      snmp3_credential = { user = "user1", authentication_protocol = "SHA", authentication_password = "authPass", privacy_protocol = "AES", privacy_password = "privPass", comment = "SNMP3 Credential Comment" }
    }
    check = {
      "nios.snmp3_credential.user"                    = "user1"
      "nios.snmp3_credential.authentication_protocol" = "SHA"
      "nios.snmp3_credential.privacy_protocol"        = "AES"
      "nios.snmp3_credential.comment"                 = "SNMP3 Credential Comment"
      "nios.snmp3_credential.credential_group"        = "default"
    }
  }

  # Update (write-only) authentication_password
  step {
    nios {
      ipv4addr         = "16.0.0.132"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int2}}"
      snmp3_credential = { user = "user1", authentication_protocol = "SHA", authentication_password = "authPass123", privacy_protocol = "AES", privacy_password = "privPass", comment = "SNMP3 Credential Comment" }
    }
    check = {
      "nios.snmp3_credential.user"                    = "user1"
      "nios.snmp3_credential.authentication_protocol" = "SHA"
      "nios.snmp3_credential.privacy_protocol"        = "AES"
      "nios.snmp3_credential.comment"                 = "SNMP3 Credential Comment"
      "nios.snmp3_credential.credential_group"        = "default"
    }
  }

  # Update (write-only) privacy_password
  step {
    nios {
      ipv4addr         = "16.0.0.132"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int2}}"
      snmp3_credential = { user = "user1", authentication_protocol = "SHA", authentication_password = "authPass123", privacy_protocol = "AES", privacy_password = "privPass123", comment = "SNMP3 Credential Comment" }
    }
    check = {
      "nios.snmp3_credential.user"                    = "user1"
      "nios.snmp3_credential.authentication_protocol" = "SHA"
      "nios.snmp3_credential.privacy_protocol"        = "AES"
      "nios.snmp3_credential.comment"                 = "SNMP3 Credential Comment"
      "nios.snmp3_credential.credential_group"        = "default"
    }
  }

  # Update both write-only fields (authentication_password and privacy_password)
  step {
    nios {
      ipv4addr         = "16.0.0.132"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int2}}"
      snmp3_credential = { user = "user1", authentication_protocol = "SHA", authentication_password = "authPass345", privacy_protocol = "AES", privacy_password = "privPass345", comment = "SNMP3 Credential Comment" }
    }
    check = {
      "nios.snmp3_credential.user"                    = "user1"
      "nios.snmp3_credential.authentication_protocol" = "SHA"
      "nios.snmp3_credential.privacy_protocol"        = "AES"
      "nios.snmp3_credential.comment"                 = "SNMP3 Credential Comment"
      "nios.snmp3_credential.credential_group"        = "default"
    }
  }

  # Update (non write-only) user field — revert to original creds
  step {
    nios {
      ipv4addr         = "16.0.0.132"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int2}}"
      snmp3_credential = { user = "user1", authentication_protocol = "SHA", authentication_password = "authPass", privacy_protocol = "AES", privacy_password = "privPass", comment = "SNMP3 Credential Comment" }
    }
    check = {
      "nios.snmp3_credential.user"                    = "user1"
      "nios.snmp3_credential.authentication_protocol" = "SHA"
      "nios.snmp3_credential.privacy_protocol"        = "AES"
      "nios.snmp3_credential.comment"                 = "SNMP3 Credential Comment"
      "nios.snmp3_credential.credential_group"        = "default"
    }
  }

  # Update user to user2
  step {
    nios {
      ipv4addr         = "16.0.0.132"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int2}}"
      snmp3_credential = { user = "user2", authentication_protocol = "SHA", authentication_password = "authPass", privacy_protocol = "AES", privacy_password = "privPass", comment = "SNMP3 Credential Comment" }
    }
    check = {
      "nios.snmp3_credential.user"                    = "user2"
      "nios.snmp3_credential.authentication_protocol" = "SHA"
      "nios.snmp3_credential.privacy_protocol"        = "AES"
      "nios.snmp3_credential.comment"                 = "SNMP3 Credential Comment"
      "nios.snmp3_credential.credential_group"        = "default"
    }
  }

}

case "snmp_credential" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ipv4addr         = "15.0.0.42"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      snmp_credential  = { community_string = "COMMUNITY_STRING", comment = "SNMP Credential Comment", credential_group = "default" }
    }
    check = {
      "nios.snmp_credential.community_string" = "COMMUNITY_STRING"
      "nios.snmp_credential.comment"          = "SNMP Credential Comment"
      "nios.snmp_credential.credential_group" = "default"
    }
  }

  step {
    nios {
      ipv4addr         = "15.0.0.42"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      snmp_credential  = { community_string = "COMMUNITY_STRING_UPDATED", comment = "SNMP Credential Comment Updated", credential_group = "default" }
    }
    check = {
      "nios.snmp_credential.community_string" = "COMMUNITY_STRING_UPDATED"
      "nios.snmp_credential.comment"          = "SNMP Credential Comment Updated"
      "nios.snmp_credential.credential_group" = "default"
    }
  }

}

# TODO: auto-extraction incomplete — please verify and fill in manually.
# Reason: requires_resource: infoblox_fixed_address_template not yet implemented
case "template" {
  backend     = "nios"
  skip        = true
  skip_reason = "requires_resource: infoblox_fixed_address_template not yet implemented"
  parallel    = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_fixed_address_template_unknown" "test" {
    nios = {
    }
  }
  PREREQ

  step {
    nios {
      ipv4addr     = "15.0.0.43"
      match_client = "MAC_ADDRESS"
      mac          = "00:1d:2e:3f:4a:5c"
      template     = "$${nios_dhcp_fixedaddresstemplate.test.name}"
    }
  }

}
