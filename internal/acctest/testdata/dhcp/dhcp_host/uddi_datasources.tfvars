# Hand-authored datasource acceptance-test cases for DhcpHost.
# DhcpHost is a system-managed object; tests reference pre-existing hosts by ID.
# Host dhcp/host/1390921 (Anycast_10_39_49_37) is used as the stable test fixture.
# The id is passed via a local so it is not a scalar literal and does not generate
# a pair-check assertion (ip_space filter may return multiple hosts).

case "filters" {
  backend = "uddi"

  prerequisites_hcl = <<-HCL
    locals {
      dhcp_host_id = "dhcp/host/1390921"
    }
  HCL

  filter {
    type = "filters"
    values = {
      "uddi.ip_space" = "uddi.ip_space"
    }
  }

  step {
    common {
      id = local.dhcp_host_id
    }
  }

}
