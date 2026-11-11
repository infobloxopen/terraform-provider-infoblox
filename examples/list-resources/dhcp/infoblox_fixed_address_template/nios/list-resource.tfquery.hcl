// List specific Fixed Address Templates using filters
list "infoblox_fixed_address_template" "list_fixed_address_template_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_fixed_address_template_1"
    }
  }
}

// List specific Fixed Address Templates using Extensible Attributes
list "infoblox_fixed_address_template" "list_fixed_address_template_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Fixed Address Templates with resource details included
list "infoblox_fixed_address_template" "list_fixed_address_template_with_resource" {
  provider         = infoblox
  include_resource = true
  limit            = 10
}
