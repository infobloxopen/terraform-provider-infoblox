// List specific NS Groups using filters
list "infoblox_nsgroup" "list_ns_groups_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_ns_group"
    }
  }
  limit = 10
}

// List specific NS Groups using Extensible Attributes
list "infoblox_nsgroup" "list_ns_groups_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List NS Groups with resource details included
list "infoblox_nsgroup" "list_ns_groups_with_resource" {
  provider         = infoblox
  include_resource = true
}
