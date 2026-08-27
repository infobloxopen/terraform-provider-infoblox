# Hand-authored datasource acceptance-test cases for DhcpHost.
# DhcpHost is a system-managed object; tests reference pre-existing hosts by name.
# The id is resolved via data source lookup so tests are env-portable.

case "filters" {
  backend = "uddi"

  prerequisites_hcl = <<-HCL
    data "infoblox_dhcp_host" "host01" {
      filters = {
        "uddi.name" = "TF_TEST_HOST_01"
      }
    }
  HCL

  filter {
    type = "filters"
    values = {
      "uddi.name" = "uddi.name"
    }
  }

  step {
    common {
      id = data.infoblox_dhcp_host.host01.results.0.id
    }
    uddi {
      server = "dhcp/server/090985ed-7e7d-11f1-85df-a6a5c5dbbe56"
    }
  }

}
