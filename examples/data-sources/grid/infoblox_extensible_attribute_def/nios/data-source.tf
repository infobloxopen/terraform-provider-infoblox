// Retrieve a specific Extensible Attribute Definition by filters
data "infoblox_extensible_attribute_def" "get_extensibleattributedef_using_filters" {
  filters = {
    name = "example_ea_1"
  }
}

// Retrieve all Extensible Attribute Definitions
data "infoblox_extensible_attribute_def" "get_all_extensibleattributedefs" {}
