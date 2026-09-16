// Retrieve an NS Group Forwarding Member by filters
data "infoblox_nsgroup_forwardingmember" "get_nsgroup_forwarding_member_using_filters" {
  filters = {
    name = "example_nsgroup_forwarding_member"
  }
}

// Retrieve an NS Group Forwarding Member using Extensible Attributes
data "infoblox_nsgroup_forwardingmember" "get_nsgroup_forwarding_member_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all NS Group Forwarding Members
data "infoblox_nsgroup_forwardingmember" "get_all_ns_group_forwarding_members" {}
