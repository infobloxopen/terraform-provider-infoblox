# Auto-generated resource acceptance-test cases for Ipv6fixedaddress.
// Objects to be present on the GRID for testing
// IPv6 Option Filters - ipv6_option_filter and ipv6_option_filter1 

case "basic" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.ipv6addr"           = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      "nios.duid"               = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      "nios.network_view"       = "{{random}}"
      "nios.network"            = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      "nios.address_type"       = "ADDRESS"
      "nios.match_client"       = "DUID"
      "nios.allow_telnet"       = "false"
      "nios.disable"            = "false"
      "nios.disable_discovery"  = "false"
      "nios.preferred_lifetime" = "27000"
      "nios.valid_lifetime"     = "43200"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }

}

case "address_type" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      address_type = "ADDRESS"
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
    }
    check = {
      "nios.address_type" = "ADDRESS"
      "nios.ipv6addr"     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
    }
  }

  step {
    nios {
      network         = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view    = infoblox_network_view.parent_network_view.nios.name
      duid            = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      address_type    = "PREFIX"
      ipv6prefix      = "2001:db8:{{random_hextet}}:{{random_int}}::"
      ipv6prefix_bits = 64
    }
    check = {
      "nios.address_type" = "PREFIX"
    }
  }

  step {
    nios {
      network         = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view    = infoblox_network_view.parent_network_view.nios.name
      duid            = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      address_type    = "BOTH"
      ipv6addr        = "2001:db8:{{random_hextet}}:{{random_int}}::2"
      ipv6prefix      = "2001:db8:{{random_hextet}}:{{random_int2}}::"
      ipv6prefix_bits = 64
    }
    check = {
      "nios.address_type" = "BOTH"
    }
  }

}

case "allow_telnet" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr        = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid            = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network         = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view    = infoblox_network_view.parent_network_view.nios.name
      allow_telnet    = true
      cli_credentials = [
        { comment = "CLI CRED Comment", user = "NIOS_USER", password = "NIOS_PASSWORD", credential_type = "TELNET", credential_group = "default" },
         { comment = "CLI CRED Comment", user = "NIOS_USER", password = "NIOS_PASSWORD", credential_type = "SSH", credential_group = "default" }
         ]
      comment         = "CLI CRED Comment"
    }
    check = {
      "nios.allow_telnet" = "true"
    }
  }

  step {
    nios {
      ipv6addr        = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid            = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network         = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view    = infoblox_network_view.parent_network_view.nios.name
      allow_telnet    = false
      cli_credentials = [{ comment = "CLI CRED Comment", user = "NIOS_USER", password = "NIOS_PASSWORD", credential_type = "TELNET", credential_group = "default" }, 
      { comment = "CLI CRED Comment", user = "NIOS_USER", password = "NIOS_PASSWORD", credential_type = "SSH", credential_group = "default" }
      ]
      comment         = "CLI CRED Comment"
    }
    check = {
      "nios.allow_telnet" = "false"
    }
  }

}

case "cli_credentials" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:5d:40:0c:39:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
      cli_credentials = [
        { comment = "Comment for CLI Credentials", user = "NIOS_USER", password = "NIOS_PASSWORD", credential_type = "SSH", credential_group = "default" },
      ]
    }
    check = {
      "nios.cli_credentials.#"                  = "1"
      "nios.cli_credentials.0.comment"          = "Comment for CLI Credentials"
      "nios.cli_credentials.0.user"             = "NIOS_USER"
      "nios.cli_credentials.0.credential_type"  = "SSH"
      "nios.cli_credentials.0.credential_group" = "default"
    }
  }

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:5d:40:0c:39:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
      cli_credentials = [
        { comment = "Comment for SSH Credentials", user = "NIOS_USER", password = "NIOS_PASSWORD", credential_type = "SSH", credential_group = "default" },
        { comment = "Updated Comment for CLI Credentials", user = "NIOS_USER", password = "NIOS_PASSWORD", credential_type = "TELNET", credential_group = "default" },
      ]
    }
    check = {
      "nios.cli_credentials.#"                  = "2"
      "nios.cli_credentials.1.comment"          = "Updated Comment for CLI Credentials"
      "nios.cli_credentials.1.user"             = "NIOS_USER"
      "nios.cli_credentials.1.credential_type"  = "TELNET"
      "nios.cli_credentials.1.credential_group" = "default"
    }
  }

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:5d:40:0c:39:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
      cli_credentials = [
        { comment = "Comment for SSH Credentials", user = "NIOS_USER", password = "NIOS_PASSWORD", credential_type = "SSH", credential_group = "default" },
        { comment = "Updated Comment for CLI Credentials", user = "NIOS_USER", password = "NIOS_PASSWORD", credential_type = "ENABLE_SSH", credential_group = "default" },
      ]
    }
    check = {
      "nios.cli_credentials.#"                  = "2"
      "nios.cli_credentials.1.comment"          = "Updated Comment for CLI Credentials"
      "nios.cli_credentials.1.user"             = "NIOS_USER"
      "nios.cli_credentials.1.credential_type"  = "ENABLE_SSH"
      "nios.cli_credentials.1.credential_group" = "default"
    }
  }

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:5d:40:0c:39:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
      cli_credentials = [
        { comment = "Comment for SSH Credentials", user = "NIOS_USER", password = "NIOS_PASSWORD", credential_type = "SSH", credential_group = "default" },
        { comment = "Updated Comment for CLI Credentials", user = "NIOS_USER", password = "NIOS_PASSWORD", credential_type = "ENABLE_TELNET", credential_group = "default" },
      ]
    }
    check = {
      "nios.cli_credentials.#"                  = "2"
      "nios.cli_credentials.1.comment"          = "Updated Comment for CLI Credentials"
      "nios.cli_credentials.1.user"             = "NIOS_USER"
      "nios.cli_credentials.1.credential_type"  = "ENABLE_TELNET"
      "nios.cli_credentials.1.credential_group" = "default"
    }
  }

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::2"
      duid         = "00:01:00:01:1d:2b:3c:5d:40:0c:39:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.cli_credentials.#" = "0"
    }
  }

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::2"
      duid         = "00:01:00:01:1d:2b:3c:5d:40:0c:39:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
      cli_credentials = [
        { comment = "cli credential comment", user = "user1", password = "password1", credential_type = "SSH", credential_group = "default" },
      ]
    }
    check = {
      "nios.cli_credentials.#"                  = "1"
      "nios.cli_credentials.0.comment"          = "cli credential comment"
      "nios.cli_credentials.0.user"             = "user1"
      "nios.cli_credentials.0.credential_type"  = "SSH"
      "nios.cli_credentials.0.credential_group" = "default"
    }
  }

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::2"
      duid         = "00:01:00:01:1d:2b:3c:5d:40:0c:39:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
      cli_credentials = [
        { comment = "cli credential comment update", user = "user2", password = "password12", credential_type = "SSH", credential_group = "default" },
      ]
    }
    check = {
      "nios.cli_credentials.#"                  = "1"
      "nios.cli_credentials.0.comment"          = "cli credential comment update"
      "nios.cli_credentials.0.user"             = "user2"
      "nios.cli_credentials.0.credential_type"  = "SSH"
      "nios.cli_credentials.0.credential_group" = "default"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
      comment      = "IPV6 Fixed Address Comment"
    }
    check = {
      "nios.comment" = "IPV6 Fixed Address Comment"
    }
  }

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
      comment      = "IPV6 Fixed Address Comment Updated"
    }
    check = {
      "nios.comment" = "IPV6 Fixed Address Comment Updated"
    }
  }

}

case "device_description" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr           = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid               = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network            = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view       = infoblox_network_view.parent_network_view.nios.name
      device_description = "{{random2}}"
    }
    check = {
      "nios.device_description" = "{{random2}}"
    }
  }

  step {
    nios {
      ipv6addr           = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid               = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network            = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view       = infoblox_network_view.parent_network_view.nios.name
      device_description = "{{random3}}"
    }
    check = {
      "nios.device_description" = "{{random3}}"
    }
  }

}

case "device_location" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr        = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid            = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network         = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view    = infoblox_network_view.parent_network_view.nios.name
      device_location = "{{random2}}"
    }
    check = {
      "nios.device_location" = "{{random2}}"
    }
  }

  step {
    nios {
      ipv6addr        = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid            = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network         = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view    = infoblox_network_view.parent_network_view.nios.name
      device_location = "{{random3}}"
    }
    check = {
      "nios.device_location" = "{{random3}}"
    }
  }

}

case "device_type" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
      device_type  = "{{random2}}"
    }
    check = {
      "nios.device_type" = "{{random2}}"
    }
  }

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
      device_type  = "{{random3}}"
    }
    check = {
      "nios.device_type" = "{{random3}}"
    }
  }

}

case "device_vendor" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr      = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid          = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network       = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view  = infoblox_network_view.parent_network_view.nios.name
      device_vendor = "{{random2}}"
    }
    check = {
      "nios.device_vendor" = "{{random2}}"
    }
  }

  step {
    nios {
      ipv6addr      = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid          = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network       = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view  = infoblox_network_view.parent_network_view.nios.name
      device_vendor = "{{random3}}"
    }
    check = {
      "nios.device_vendor" = "{{random3}}"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
      disable      = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
      disable      = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

}

case "disable_discovery" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr          = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid              = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network           = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view      = infoblox_network_view.parent_network_view.nios.name
      disable_discovery = true
    }
    check = {
      "nios.disable_discovery" = "true"
    }
  }

  step {
    nios {
      ipv6addr          = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid              = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network           = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view      = infoblox_network_view.parent_network_view.nios.name
      disable_discovery = false
    }
    check = {
      "nios.disable_discovery" = "false"
    }
  }

}

case "domain_name" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
      domain_name  = "{{random2}}.com"
    }
    check = {
      "nios.domain_name" = "{{random2}}.com"
    }
  }

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
      domain_name  = "{{random3}}.com"
    }
    check = {
      "nios.domain_name" = "{{random3}}.com"
    }
  }

}

case "domain_name_servers" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr            = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid                = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network             = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view        = infoblox_network_view.parent_network_view.nios.name
      domain_name_servers = ["2001:4860:4860::8888", "2001:4860:4860::8844"]
    }
    check = {
      "nios.domain_name_servers.#" = "2"
      "nios.domain_name_servers.0" = "2001:4860:4860::8888"
      "nios.domain_name_servers.1" = "2001:4860:4860::8844"
    }
  }

  step {
    nios {
      ipv6addr            = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid                = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network             = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view        = infoblox_network_view.parent_network_view.nios.name
      domain_name_servers = ["2620:fe::9", "2001:4860:4860::6844"]
    }
    check = {
      "nios.domain_name_servers.#" = "2"
      "nios.domain_name_servers.0" = "2620:fe::9"
      "nios.domain_name_servers.1" = "2001:4860:4860::6844"
    }
  }

}

case "duid" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.duid" = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
    }
  }

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:11:11:2b:3c:4d:00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.duid" = "00:01:00:11:11:2b:3c:4d:00:0c:29:ab:cd:ef"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
      ext_attrs    = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
      ext_attrs    = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "ipv6addr" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.ipv6addr" = "2001:db8:{{random_hextet}}:{{random_int}}::1"
    }
  }

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::2"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.ipv6addr" = "2001:db8:{{random_hextet}}:{{random_int}}::2"
    }
  }

}

case "ipv6prefix" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      duid            = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      address_type    = "PREFIX"
      ipv6prefix      = "2001:db8:{{random_hextet}}:{{random_int}}::"
      ipv6prefix_bits = 64
      network         = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view    = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.ipv6prefix" = "2001:db8:{{random_hextet}}:{{random_int}}::"
    }
  }

  step {
    nios {
      duid            = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      address_type    = "PREFIX"
      ipv6prefix      = "2001:db8:{{random_hextet}}:{{random_int2}}::"
      ipv6prefix_bits = 64
      network         = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view    = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.ipv6prefix" = "2001:db8:{{random_hextet}}:{{random_int2}}::"
    }
  }

}

case "ipv6prefix_bits" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      duid            = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      address_type    = "BOTH"
      ipv6addr        = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      ipv6prefix      = "2001:db8:{{random_hextet}}:{{random_int}}::"
      ipv6prefix_bits = 64
      network         = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view    = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.ipv6prefix_bits" = "64"
    }
  }

  step {
    nios {
      duid            = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      address_type    = "BOTH"
      ipv6addr        = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      ipv6prefix      = "2001:db8:{{random_hextet}}:{{random_int}}::"
      ipv6prefix_bits = 65
      network         = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view    = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.ipv6prefix_bits" = "65"
    }
  }

}

case "logic_filter_rules" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr           = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid               = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network            = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view       = infoblox_network_view.parent_network_view.nios.name
      logic_filter_rules = [{ filter = "ipv6_option_filter", type = "Option" }]
    }
    check = {
      "nios.logic_filter_rules.#"        = "1"
      "nios.logic_filter_rules.0.filter" = "ipv6_option_filter"
      "nios.logic_filter_rules.0.type"   = "Option"
    }
  }

  step {
    nios {
      ipv6addr           = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid               = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network            = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view       = infoblox_network_view.parent_network_view.nios.name
      logic_filter_rules = [{ filter = "ipv6_option_filter1", type = "Option" }]
    }
    check = {
      "nios.logic_filter_rules.#"        = "1"
      "nios.logic_filter_rules.0.filter" = "ipv6_option_filter1"
      "nios.logic_filter_rules.0.type"   = "Option"
    }
  }

}

case "mac_address" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      match_client = "MAC_ADDRESS"
      mac_address  = "00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.mac_address" = "00:0c:29:ab:cd:ef"
    }
  }

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      match_client = "MAC_ADDRESS"
      mac_address  = "01:2c:39:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.mac_address" = "01:2c:39:ab:cd:ef"
    }
  }

}

case "match_client" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      match_client = "DUID"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.match_client" = "DUID"
    }
  }

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      mac_address  = "00:0c:29:ab:cd:ef"
      match_client = "MAC_ADDRESS"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.match_client" = "MAC_ADDRESS"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      name         = "{{random2}}"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      name         = "{{random3}}"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.name" = "{{random3}}"
    }
  }

}

case "network" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6network1" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_ipv6_network" "test_ipv6network2" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int2}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6network1.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.network" = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
    }
  }

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int2}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6network2.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.network" = "2001:db8:{{random_hextet}}:{{random_int2}}::/64"
    }
  }

}

case "options" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
      options = [
        { name = "dhcp6.domain-search", num = 24, value = "\"aa.bb.com\"", vendor_class = "DHCPv6" },
        { name = "dhcp6.sntp-servers", num = 31, value = "2001:4860:4860::8888", vendor_class = "DHCPv6" },
      ]
    }
    check = {
      "nios.options.0.name"  = "dhcp6.domain-search"
      "nios.options.0.value" = "\"aa.bb.com\""
      "nios.options.1.name"  = "dhcp6.sntp-servers"
      "nios.options.1.value" = "2001:4860:4860::8888"
    }
  }

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
      options = [
        { name = "dhcp6.domain-search", num = 24, value = "\"bb.cc.com\"", vendor_class = "DHCPv6" },
        { name = "dhcp6.sntp-servers", num = 31, value = "2001:4860:4860::8008", vendor_class = "DHCPv6" },
      ]
    }
    check = {
      "nios.options.0.name"  = "dhcp6.domain-search"
      "nios.options.0.value" = "\"bb.cc.com\""
      "nios.options.1.name"  = "dhcp6.sntp-servers"
      "nios.options.1.value" = "2001:4860:4860::8008"
    }
  }

}

case "preferred_lifetime" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr           = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid               = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      preferred_lifetime = 6200
      valid_lifetime     = 43200
      network            = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view       = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.preferred_lifetime" = "6200"
    }
  }

  step {
    nios {
      ipv6addr           = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid               = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      preferred_lifetime = 4800
      valid_lifetime     = 43200
      network            = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view       = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.preferred_lifetime" = "4800"
    }
  }

}

case "template" {
  backend     = "nios"
  parallel    = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  # resource "infoblox_ipv6_fixed_address_template_unknown" "test" {
  #   nios = {
  #     name = "{{random}}"
  #   }
  # }
  PREREQ

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      template     = "ipv6-fa-template"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.template" = "ipv6-fa-template"
    }
  }

}

case "reserved_interface" {
  backend     = "nios"
  skip        = true
  skip_reason = "t.Skip: Skipping test as reserved_interface is not implemented yet"
  parallel    = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr           = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid               = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      reserved_interface = "{{random2}}"
      network            = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view       = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.reserved_interface" = "{{random2}}"
    }
  }

  step {
    nios {
      ipv6addr           = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid               = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      reserved_interface = "{{random3}}"
      network            = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view       = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.reserved_interface" = "{{random3}}"
    }
  }

}

case "snmp3_credential" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr         = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid             = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      snmp3_credential = { user = "snmp", authentication_protocol = "MD5", authentication_password = "snmp1234", privacy_protocol = "3DES", privacy_password = "snmp1234", comment = "SNMP3 Credential Comment", credential_group = "default" }
      network          = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view     = infoblox_network_view.parent_network_view.nios.name
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
      ipv6addr         = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid             = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      snmp3_credential = { user = "SNMP3_USER_UPDATE", authentication_protocol = "SHA-224", authentication_password = "AUTH_PASSWORD", privacy_protocol = "AES-256", privacy_password = "PRIVACY_PASSWORD", comment = "SNMP3 Credential Comment Updated", credential_group = "default" }
      network          = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view     = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.snmp3_credential.user"                    = "SNMP3_USER_UPDATE"
      "nios.snmp3_credential.authentication_protocol" = "SHA-224"
      "nios.snmp3_credential.privacy_protocol"        = "AES-256"
      "nios.snmp3_credential.comment"                 = "SNMP3 Credential Comment Updated"
      "nios.snmp3_credential.credential_group"        = "default"
    }
  }

  step {
    nios {
      ipv6addr         = "2001:db8:{{random_hextet}}:{{random_int}}::2"
      duid             = "00:01:00:01:1d:2b:3c:4d:01:1c:29:ab:cd:ef"
      snmp3_credential = { user = "user1", authentication_protocol = "SHA", authentication_password = "authPass", privacy_protocol = "AES", privacy_password = "privPass", comment = "SNMP3 Credential Comment", credential_group = "default" }
      network          = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view     = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.snmp3_credential.user"                    = "user1"
      "nios.snmp3_credential.authentication_protocol" = "SHA"
      "nios.snmp3_credential.privacy_protocol"        = "AES"
      "nios.snmp3_credential.comment"                 = "SNMP3 Credential Comment"
      "nios.snmp3_credential.credential_group"        = "default"
    }
  }

  step {
    nios {
      ipv6addr         = "2001:db8:{{random_hextet}}:{{random_int}}::2"
      duid             = "00:01:00:01:1d:2b:3c:4d:01:1c:29:ab:cd:ef"
      snmp3_credential = { user = "user1", authentication_protocol = "SHA", authentication_password = "authPass345", privacy_protocol = "AES", privacy_password = "privPass345", comment = "SNMP3 Credential Comment", credential_group = "default" }
      network          = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view     = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.snmp3_credential.user"                    = "user1"
      "nios.snmp3_credential.authentication_protocol" = "SHA"
      "nios.snmp3_credential.privacy_protocol"        = "AES"
      "nios.snmp3_credential.comment"                 = "SNMP3 Credential Comment"
      "nios.snmp3_credential.credential_group"        = "default"
    }
  }

  step {
    nios {
      ipv6addr         = "2001:db8:{{random_hextet}}:{{random_int}}::2"
      duid             = "00:01:00:01:1d:2b:3c:4d:01:1c:29:ab:cd:ef"
      snmp3_credential = { user = "user2", authentication_protocol = "SHA", authentication_password = "authPass", privacy_protocol = "AES", privacy_password = "privPass", comment = "SNMP3 Credential Comment", credential_group = "default" }
      network          = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view     = infoblox_network_view.parent_network_view.nios.name
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
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr        = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid            = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      snmp_credential = { community_string = "COMMUNITY_STRING", comment = "SNMP Credential Comment", credential_group = "default" }
      network         = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view    = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.snmp_credential.community_string" = "COMMUNITY_STRING"
      "nios.snmp_credential.comment"          = "SNMP Credential Comment"
      "nios.snmp_credential.credential_group" = "default"
    }
  }

  step {
    nios {
      ipv6addr        = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid            = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      snmp_credential = { community_string = "COMMUNITY_STRING_UPDATED", comment = "SNMP Credential Comment Updated", credential_group = "default" }
      network         = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view    = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.snmp_credential.community_string" = "COMMUNITY_STRING_UPDATED"
      "nios.snmp_credential.comment"          = "SNMP Credential Comment Updated"
      "nios.snmp_credential.credential_group" = "default"
    }
  }

}

case "valid_lifetime" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      ipv6addr       = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid           = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      valid_lifetime = 42800
      network        = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view   = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.valid_lifetime" = "42800"
    }
  }

  step {
    nios {
      ipv6addr       = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      duid           = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      valid_lifetime = 56000
      network        = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view   = infoblox_network_view.parent_network_view.nios.name
    }
    check = {
      "nios.valid_lifetime" = "56000"
    }
  }

}

case "dynamic_allocation" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      duid         = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
      network_view = infoblox_network_view.parent_network_view.nios.name
      dynamic_allocation = {
        network      = infoblox_ipv6_network.test_ipv6_network.nios.network
        network_view = infoblox_network_view.parent_network_view.nios.name
      }
      comment = "Created by Dynamic Allocation"
    }
    check = {
      "nios.comment"      = "Created by Dynamic Allocation"
      "nios.address_type" = "ADDRESS"
    }
  }

}
