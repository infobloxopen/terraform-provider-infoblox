// List specific Range Templates using filters
list "infoblox_range_template" "list_rangetemplate_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_range_template"
    }
  }
  limit = 10
}

// List specific Range Templates using Extensible Attributes
list "infoblox_range_template" "list_rangetemplate_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      "Tenant ID" = "tenant-1"
    }
  }
}

// List Range Templates with resource details included
list "infoblox_range_template" "list_rangetemplate_with_resource" {
  provider         = infoblox
  include_resource = true
}
