// List NS Group Forwarding Members using filters
list "infoblox_nsgroup_forwardingmember" "list_ns_group_forwarding_member_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_nsgroup_forwarding_member"
    }
  }
  limit = 10
}

// List NS Group Forwarding Members using Extensible Attributes
list "infoblox_nsgroup_forwardingmember" "list_ns_group_forwarding_member_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List NS Group Forwarding Members with resource details included
list "infoblox_nsgroup_forwardingmember" "list_ns_group_forwarding_member_with_resource" {
  provider         = infoblox
  include_resource = true
}
