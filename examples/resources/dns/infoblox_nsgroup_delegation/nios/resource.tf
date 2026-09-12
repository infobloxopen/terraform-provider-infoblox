// Manage an NS group delegation
resource "infoblox_nsgroup_delegation" "nsgroup_delegation_basic_fields" {
  nios = {
    name = "example_ns_group_del"
    delegate_to = [
      {
        address = "2.3.3.4"
        name    = "delegate_to_ns_group"
      }
    ]
  }
}

// Manage an NS group delegation with additional attributes
resource "infoblox_nsgroup_delegation" "nsgroup_delegation_with_additional_fields" {
  nios = {
    name = "example_ns_group_delegation"
    delegate_to = [
      {
        address = "2.3.3.5"
        name    = "delegate_to_ns_group"
      }
    ]
    comment = "Create NS Group Delegation"
    ext_attrs = {
      Site = "location-1"
    }
  }
}
