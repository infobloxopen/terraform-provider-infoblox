// Retrieve a specific NS Group Delegation by filters
data "infoblox_nsgroup_delegation" "get_ns_group_delegation_using_filters" {
  filters = {
    name = "example_ns_group_delegation"
  }
}

// Retrieve specific NS Group Delegations using Extensible Attributes
data "infoblox_nsgroup_delegation" "get_ns_group_delegation_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all NS Group Delegations
data "infoblox_nsgroup_delegation" "get_all_ns_group_delegations" {
}
