# Auto-generated resource acceptance-test cases for Sharednetwork.

case "basic" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.1.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.1.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
    }
    check = {
      "nios.name"                            = "{{random}}"
      "nios.authority"                       = "false"
      "nios.ddns_generate_hostname"          = "false"
      "nios.ddns_server_always_updates"      = "true"
      "nios.ddns_ttl"                        = "0"
      "nios.ddns_update_fixed_addresses"     = "false"
      "nios.ddns_use_option81"               = "false"
      "nios.deny_bootp"                      = "false"
      "nios.disable"                         = "false"
      "nios.enable_ddns"                     = "false"
      "nios.enable_pxe_lease_time"           = "false"
      "nios.ignore_client_identifier"        = "false"
      "nios.ignore_dhcp_option_list_request" = "false"
      "nios.ignore_id"                       = "NONE"
      "nios.lease_scavenge_time"             = "-1"
      "nios.update_dns_on_lease_renewal"     = "false"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl     = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.1.1.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.1.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
    }
  }

}

case "authority" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.3.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.3.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
      authority    = true
    }
    check = {
      "nios.authority" = "true"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
      authority    = false
    }
    check = {
      "nios.authority" = "false"
    }
  }

}

case "bootfile" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.5.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.5.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
      bootfile     = "boot.txt"
    }
    check = {
      "nios.bootfile" = "boot.txt"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
      bootfile     = "boot_updated.txt"
    }
    check = {
      "nios.bootfile" = "boot_updated.txt"
    }
  }

}

case "bootserver" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.7.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.7.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
      bootserver   = "boot-server1"
    }
    check = {
      "nios.bootserver" = "boot-server1"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
      bootserver   = "boot-server2"
    }
    check = {
      "nios.bootserver" = "boot-server2"
    }
  }

}

case "comment" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.9.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.9.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
      comment      = "shared network comment"
    }
    check = {
      "nios.comment" = "shared network comment"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
      comment      = "updated shared network comment"
    }
    check = {
      "nios.comment" = "updated shared network comment"
    }
  }

}

case "ddns_generate_hostname" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.11.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.11.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name                   = "{{random}}"
      networks               = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view           = infoblox_network_view.test_view.nios.name
      ddns_generate_hostname = true
    }
    check = {
      "nios.ddns_generate_hostname" = "true"
    }
  }

  step {
    nios {
      name                   = "{{random}}"
      networks               = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view           = infoblox_network_view.test_view.nios.name
      ddns_generate_hostname = false
    }
    check = {
      "nios.ddns_generate_hostname" = "false"
    }
  }

}

case "ddns_server_always_updates" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.13.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.13.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name                       = "{{random}}"
      networks                   = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view               = infoblox_network_view.test_view.nios.name
      ddns_server_always_updates = true
      ddns_use_option81          = true
    }
    check = {
      "nios.ddns_server_always_updates" = "true"
    }
  }

  step {
    nios {
      name                       = "{{random}}"
      networks                   = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view               = infoblox_network_view.test_view.nios.name
      ddns_server_always_updates = false
      ddns_use_option81          = true
    }
    check = {
      "nios.ddns_server_always_updates" = "false"
    }
  }

}

case "ddns_ttl" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.15.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.15.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
      ddns_ttl     = 3600
    }
    check = {
      "nios.ddns_ttl" = "3600"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
      ddns_ttl     = 7200
    }
    check = {
      "nios.ddns_ttl" = "7200"
    }
  }

}

case "ddns_update_fixed_addresses" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.17.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.17.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name                        = "{{random}}"
      networks                    = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view                = infoblox_network_view.test_view.nios.name
      ddns_update_fixed_addresses = true
    }
    check = {
      "nios.ddns_update_fixed_addresses" = "true"
    }
  }

  step {
    nios {
      name                        = "{{random}}"
      networks                    = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view                = infoblox_network_view.test_view.nios.name
      ddns_update_fixed_addresses = false
    }
    check = {
      "nios.ddns_update_fixed_addresses" = "false"
    }
  }

}

case "ddns_use_option81" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.19.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.19.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name              = "{{random}}"
      networks          = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view      = infoblox_network_view.test_view.nios.name
      ddns_use_option81 = true
    }
    check = {
      "nios.ddns_use_option81" = "true"
    }
  }

  step {
    nios {
      name              = "{{random}}"
      networks          = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view      = infoblox_network_view.test_view.nios.name
      ddns_use_option81 = false
    }
    check = {
      "nios.ddns_use_option81" = "false"
    }
  }

}

case "deny_bootp" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.21.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.21.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
      deny_bootp   = true
    }
    check = {
      "nios.deny_bootp" = "true"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
      deny_bootp   = false
    }
    check = {
      "nios.deny_bootp" = "false"
    }
  }

}

case "disable" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.23.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.23.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
      disable      = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
      disable      = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

}

case "enable_ddns" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.25.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.25.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
      enable_ddns  = true
    }
    check = {
      "nios.enable_ddns" = "true"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
      enable_ddns  = false
    }
    check = {
      "nios.enable_ddns" = "false"
    }
  }

}

case "enable_pxe_lease_time" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.27.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.27.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name                  = "{{random}}"
      networks              = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view          = infoblox_network_view.test_view.nios.name
      enable_pxe_lease_time = true
      pxe_lease_time        = 43200
    }
    check = {
      "nios.enable_pxe_lease_time" = "true"
    }
  }

  step {
    nios {
      name                  = "{{random}}"
      networks              = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view          = infoblox_network_view.test_view.nios.name
      enable_pxe_lease_time = false
      pxe_lease_time        = 43200
    }
    check = {
      "nios.enable_pxe_lease_time" = "false"
    }
  }

}

case "ext_attrs" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.29.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.29.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
      ext_attrs    = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
      ext_attrs    = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

# WARNING: the extractor could not auto-record the following line(s) from
# the Go helper. Some fields may not be correctly captured — please verify
# this case manually against the original test before running:
#   %s
case "ignore_client_identifier" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.31.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.31.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name                     = "{{random}}"
      networks                 = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view             = infoblox_network_view.test_view.nios.name
      ignore_client_identifier = true
      ignore_id                = "CLIENT"
    }
    check = {
      "nios.ignore_client_identifier" = "true"
    }
  }

  step {
    nios {
      name                     = "{{random}}"
      networks                 = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view             = infoblox_network_view.test_view.nios.name
      ignore_client_identifier = false
      ignore_id                = "NONE"
    }
    check = {
      "nios.ignore_client_identifier" = "false"
    }
  }

}

case "ignore_dhcp_option_list_request" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.33.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.33.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name                            = "{{random}}"
      networks                        = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view                    = infoblox_network_view.test_view.nios.name
      ignore_dhcp_option_list_request = true
    }
    check = {
      "nios.ignore_dhcp_option_list_request" = "true"
    }
  }

  step {
    nios {
      name                            = "{{random}}"
      networks                        = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view                    = infoblox_network_view.test_view.nios.name
      ignore_dhcp_option_list_request = false
    }
    check = {
      "nios.ignore_dhcp_option_list_request" = "false"
    }
  }

}

# WARNING: the extractor could not auto-record the following line(s) from
# the Go helper. Some fields may not be correctly captured — please verify
# this case manually against the original test before running:
#   %s
case "ignore_id" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.35.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.35.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name                     = "{{random}}"
      networks                 = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view             = infoblox_network_view.test_view.nios.name
      ignore_id                = "CLIENT"
      ignore_client_identifier = true
    }
    check = {
      "nios.ignore_id" = "CLIENT"
    }
  }

  step {
    nios {
      name                     = "{{random}}"
      networks                 = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view             = infoblox_network_view.test_view.nios.name
      ignore_id                = "NONE"
      ignore_client_identifier = false
    }
    check = {
      "nios.ignore_id" = "NONE"
    }
  }

}

case "ignore_mac_addresses" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.37.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.37.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name                 = "{{random}}"
      networks             = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view         = infoblox_network_view.test_view.nios.name
      ignore_mac_addresses = ["00:11:22:33:44:55", "66:77:88:99:aa:bb"]
    }
    check = {
      "nios.ignore_mac_addresses.#" = "2"
      "nios.ignore_mac_addresses.0" = "00:11:22:33:44:55"
      "nios.ignore_mac_addresses.1" = "66:77:88:99:aa:bb"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      networks             = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view         = infoblox_network_view.test_view.nios.name
      ignore_mac_addresses = ["00:11:22:33:44:88", "00:11:22:33:44:55"]
    }
    check = {
      "nios.ignore_mac_addresses.#" = "2"
      "nios.ignore_mac_addresses.0" = "00:11:22:33:44:88"
      "nios.ignore_mac_addresses.1" = "00:11:22:33:44:55"
    }
  }

}

case "lease_scavenge_time" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.39.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.39.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      networks            = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view        = infoblox_network_view.test_view.nios.name
      lease_scavenge_time = 86420
    }
    check = {
      "nios.lease_scavenge_time" = "86420"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      networks            = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view        = infoblox_network_view.test_view.nios.name
      lease_scavenge_time = 214440
    }
    check = {
      "nios.lease_scavenge_time" = "214440"
    }
  }

}

case "logic_filter_rules" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.41.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.41.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
    }
    check = {
      "nios.logic_filter_rules.#"        = "1"
      "nios.logic_filter_rules.0.filter" = "example-option-filter-1"
      "nios.logic_filter_rules.0.type"   = "Option"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
    }
    check = {
      "nios.logic_filter_rules.#"        = "1"
      "nios.logic_filter_rules.0.filter" = "example-option-filter-2"
      "nios.logic_filter_rules.0.type"   = "Option"
    }
  }

}

case "name" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.43.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.43.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name         = "{{random2}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "networks" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "202.45.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.45.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network3" {
    nios = {
      network      = "211.45.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network4" {
    nios = {
      network      = "212.45.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
    }
    check = {
      "nios.networks.#" = "2"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network3.id }, { ref = infoblox_network.test_network4.id }]
      network_view = infoblox_network_view.test_view.nios.name
    }
    check = {
      "nios.networks.#" = "2"
    }
  }

}

case "nextserver" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.47.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.47.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
      nextserver   = "nest-server1"
    }
    check = {
      "nios.nextserver" = "nest-server1"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
      nextserver   = "next-server2"
    }
    check = {
      "nios.nextserver" = "next-server2"
    }
  }

}

case "options" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.49.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.49.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
    }
    check = {
      "nios.options.#"       = "2"
      "nios.options.0.name"  = "domain-name"
      "nios.options.0.value" = "aa.bb.com"
      "nios.options.1.name"  = "dhcp-lease-time"
      "nios.options.1.value" = "72000"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view = infoblox_network_view.test_view.nios.name
    }
    check = {
      "nios.options.#"       = "2"
      "nios.options.0.name"  = "time-offset"
      "nios.options.0.value" = "50"
      "nios.options.1.name"  = "dhcp-lease-time"
      "nios.options.1.value" = "82000"
    }
  }

}

case "pxe_lease_time" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.51.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.51.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name           = "{{random}}"
      networks       = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view   = infoblox_network_view.test_view.nios.name
      pxe_lease_time = 3600
    }
    check = {
      "nios.pxe_lease_time" = "3600"
    }
  }

  step {
    nios {
      name           = "{{random}}"
      networks       = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view   = infoblox_network_view.test_view.nios.name
      pxe_lease_time = 7200
    }
    check = {
      "nios.pxe_lease_time" = "7200"
    }
  }

}

case "update_dns_on_lease_renewal" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.53.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "210.53.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name                        = "{{random}}"
      networks                    = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view                = infoblox_network_view.test_view.nios.name
      update_dns_on_lease_renewal = true
    }
    check = {
      "nios.update_dns_on_lease_renewal" = "true"
    }
  }

  step {
    nios {
      name                        = "{{random}}"
      networks                    = [{ ref = infoblox_network.test_network1.id }, { ref = infoblox_network.test_network2.id }]
      network_view                = infoblox_network_view.test_view.nios.name
      update_dns_on_lease_renewal = false
    }
    check = {
      "nios.update_dns_on_lease_renewal" = "false"
    }
  }

}
