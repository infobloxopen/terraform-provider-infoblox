// Retrieve a NS Group by filters
data "infoblox_nsgroup" "get_ns_group_using_filters" {
  filters = {
    name = "example_ns_group"
  }
}

// Retrieve a NS Group using Extensible Attributes
data "infoblox_nsgroup" "get_ns_group_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all NS Groups
data "infoblox_nsgroup" "get_all_ns_groups" {}
