// List specific NS Group Delegations using filters
list "infoblox_nsgroup_delegation" "list_ns_group_delegations_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_ns_group_del"
    }
  }
  limit = 10
}

// List specific NS Group Delegations using Extensible Attributes
list "infoblox_nsgroup_delegation" "list_ns_group_delegations_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List NS Group Delegations with resource details included
list "infoblox_nsgroup_delegation" "list_ns_group_delegations_with_resource" {
  provider         = infoblox
  include_resource = true
}
