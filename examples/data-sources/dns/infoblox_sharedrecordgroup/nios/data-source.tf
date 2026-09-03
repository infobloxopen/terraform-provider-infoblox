// Retrieve a specific Shared Record Group by filters
data "infoblox_sharedrecordgroup" "get_sharedrecordgroup_using_filters" {
  filters = {
    name = "example-shared-record-group"
  }
}

// Retrieve specific Shared Record Groups using Extensible Attributes
data "infoblox_sharedrecordgroup" "get_sharedrecordgroups_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all Shared Record Groups
data "infoblox_sharedrecordgroup" "get_all_sharedrecordgroup" {}
