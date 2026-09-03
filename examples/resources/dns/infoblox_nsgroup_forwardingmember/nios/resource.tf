// Create an NS Group Forwarding Member with Basic Fields
resource "infoblox_nsgroup_forwardingmember" "nsgroup_forwarding_member_basic_fields" {
  nios = {
    name = "example_nsgroup_forwarding_member"
    forwarding_servers = [
      {
        name = "infoblox.localdomain"
      }
    ]
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create an NS Group Forwarding Member with Additional Fields
resource "infoblox_nsgroup_forwardingmember" "nsgroup_forwarding_member_additional_fields" {
  nios = {
    name    = "example_nsgroup_forwarding_member1"
    comment = "nsgroup forwarding member with additional fields"
    forwarding_servers = [
      {
        name            = "infoblox.localdomain"
        forwarders_only = true
        forward_to = [
          {
            name    = "forwarder.localdomain"
            address = "2.3.4.45"
          }
        ]
        use_override_forwarders = true
      }
    ]
    ext_attrs = {
      Site = "location-1"
    }
  }
}
