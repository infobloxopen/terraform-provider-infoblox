// Retrieve a NS Group Stub Member by filters
data "infoblox_nsgroup_stubmember" "get_ns_group_stubmember_using_filters" {
  filters = {
    name = "example_ns_group"
  }
}

// Retrieve NS Group Stub Members using Extensible Attributes
data "infoblox_nsgroup_stubmember" "get_ns_group_stubmember_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all NS Group Stub Members
data "infoblox_nsgroup_stubmember" "get_all_ns_group_stubmembers" {}
