// Retrieve a specific Range Template by filters
data "infoblox_range_template" "get_range_template_using_filters" {
  filters = {
    name = "example_range_template"
  }
}

// Retrieve specific Range Templates using Extensible Attributes
data "infoblox_range_template" "get_range_template_using_extensible_attributes" {
  ext_attr_filters = {
    "Tenant ID" = "tenant-1"
  }
}

// Retrieve all Range Templates
data "infoblox_range_template" "get_all_range_templates" {}
