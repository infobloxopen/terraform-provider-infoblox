// Retrieve a specific Extensible Attribute Definition by filters
data "infoblox_extensibleattributedef" "get_extensibleattributedef_using_filters" {
  filters = {
    name = "example_ea_1"
  }
}

// Retrieve all Extensible Attribute Definitions
data "infoblox_extensibleattributedef" "get_all_extensibleattributedefs" {}
