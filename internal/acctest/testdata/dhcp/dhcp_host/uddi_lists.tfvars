# Hand-authored list acceptance-test cases for DhcpHost.
# DhcpHost is a system-managed object; tests reference pre-existing hosts by ID.
# Host dhcp/host/1390921 (Anycast_10_39_49_37) is used as the stable test fixture.

case "basic" {
  backend        = "uddi"
  min_tf_version = "1.14.0"

  step {
    common {
      id = "dhcp/host/1390921"
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}
