# Hand-authored resource acceptance-test cases for DhcpHost.
# DhcpHost is a system-managed object; tests reference pre-existing hosts by name.
# Host IDs are resolved dynamically via data source lookup so tests are env-portable.

case "basic" {
  backend  = "uddi"
  parallel = true

  prerequisites_hcl = <<-HCL
    data "infoblox_dhcp_host" "host01" {
      filters = {
        "uddi.name" = "TF_TEST_HOST_01"
      }
    }
  HCL

  step {
    common {
      id = data.infoblox_dhcp_host.host01.results.0.id
    }
    uddi {
      server = "dhcp/server/090985ed-7e7d-11f1-85df-a6a5c5dbbe56"
    }
    check = {
      "uddi.server" = "dhcp/server/090985ed-7e7d-11f1-85df-a6a5c5dbbe56"
      "uddi.name"   = "TF_TEST_HOST_01"
    }
  }

}

case "server" {
  backend  = "uddi"
  parallel = true

  prerequisites_hcl = <<-HCL
    data "infoblox_dhcp_host" "host02" {
      filters = {
        "uddi.name" = "TF_TEST_HOST_02"
      }
    }
  HCL

  step {
    common {
      id = data.infoblox_dhcp_host.host02.results.0.id
    }
    uddi {
      server = "dhcp/server/090985ed-7e7d-11f1-85df-a6a5c5dbbe56"
    }
    check = {
      "uddi.server" = "dhcp/server/090985ed-7e7d-11f1-85df-a6a5c5dbbe56"
    }
  }

  step {
    common {
      id = data.infoblox_dhcp_host.host02.results.0.id
    }
    uddi {
      server = "dhcp/server/0b6ff225-7e7d-11f1-b89b-0e6a1c02f67a"
    }
    check = {
      "uddi.server" = "dhcp/server/0b6ff225-7e7d-11f1-b89b-0e6a1c02f67a"
    }
  }

}
