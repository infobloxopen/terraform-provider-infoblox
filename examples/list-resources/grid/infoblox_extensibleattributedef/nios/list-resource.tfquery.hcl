// List specific Extensible Attribute Definitions using filters
list "infoblox_extensibleattributedef" "list_extensibleattributedefs_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_ea_1"
    }
  }
}

// List Extensible Attribute Definitions with resource details included
list "infoblox_extensibleattributedef" "list_extensibleattributedefs_with_resource" {
  provider         = infoblox
  include_resource = true
}
