// Create a basic active-active HA Group
// Note: replace host IDs with real DHCP host resource identifiers (dhcp/host/<uuid>)
resource "infoblox_ha_group" "example" {
  uddi = {
    name = "example-ha-group"
    mode = "active-active"
    hosts = [
      {
        host = "dhcp/host/00000000-0000-0000-0000-000000000001"
        role = "active"
      },
      {
        host = "dhcp/host/00000000-0000-0000-0000-000000000002"
        role = "active"
      }
    ]
  }
}

// Create an active-passive HA Group with additional fields
resource "infoblox_ha_group" "example_with_options" {
  uddi = {
    name    = "example-ha-group-passive"
    mode    = "active-passive"
    comment = "HA Group created with Terraform"
    hosts = [
      {
        host = "dhcp/host/00000000-0000-0000-0000-000000000001"
        role = "active"
      },
      {
        host = "dhcp/host/00000000-0000-0000-0000-000000000002"
        role = "passive"
      }
    ]
    tags = {
      location = "site-1"
    }
  }
}
