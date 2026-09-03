// List NS Group Stub Members using filters
list "infoblox_nsgroup_stubmember" "list_ns_group_stub_member_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_stubmember"
    }
  }
  limit = 10
}

// List NS Group Stub Members using Extensible Attributes
list "infoblox_nsgroup_stubmember" "list_ns_group_stub_member_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List NS Group Stub Members with resource details included
list "infoblox_nsgroup_stubmember" "list_ns_group_stub_member_with_resource" {
  provider         = infoblox
  include_resource = true
}
