// List specific Fixed Addresses using filters
list "infoblox_fixed_address" "list_fixed_addresses_using_filters" {
  provider = infoblox
  config {
    filters = {
      ipv4addr = "16.0.0.20"
    }
  }
  limit = 10
}

// List specific Fixed Addresses using Extensible Attributes
list "infoblox_fixed_address" "list_fixed_addresses_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Fixed Addresses with resource details included
list "infoblox_fixed_address" "list_fixed_addresses_with_resource" {
  provider         = infoblox
  include_resource = true
}
