# Hand-authored resource acceptance-test cases for DhcpHost.
# DhcpHost is a system-managed object; resources reference pre-existing hosts by ID.
# Host dhcp/host/1390921 (Anycast_10_39_49_37) is used as the stable test fixture.

case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    common {
      id = "dhcp/host/1390921"
    }
    check = {
      "id"            = "dhcp/host/1390921"
      "uddi.ip_space" = "ipam/ip_space/5202ccf2-f3b6-11ed-a04c-0acb29431b1f"
    }
  }

}
