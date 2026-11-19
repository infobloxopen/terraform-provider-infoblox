// Retrieve a specific Fixed Address Template by filters
data "infoblox_fixed_address_template" "get_fixed_address_template_using_filters" {
  filters = {
    name = "example_fixed_address_template_1"
  }
}

// Retrieve specific Fixed Address Templates using Extensible Attributes
data "infoblox_fixed_address_template" "get_fixed_address_template_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all Fixed Address Templates
data "infoblox_fixed_address_template" "get_all_fixed_address_templates" {}
