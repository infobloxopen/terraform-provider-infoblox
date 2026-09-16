// List specific Shared Record Groups using filters
list "infoblox_sharedrecordgroup" "list_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example-shared-record-group"
    }
  }
  limit = 10
}

// List specific Shared Record Groups using Extensible Attributes
list "infoblox_sharedrecordgroup" "list_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Shared Record Groups with resource details included
list "infoblox_sharedrecordgroup" "list_with_resource" {
  provider         = infoblox
  include_resource = true
}
