// List specific Fixed Addresses using filters
list "infoblox_fixed_address" "list_fixed_address_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_fixed_address"
    }
  }
  limit = 10
}

// List specific Fixed Addresses using Tags
list "infoblox_fixed_address" "list_fixed_address_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Fixed Addresses with resource details included
list "infoblox_fixed_address" "list_fixed_address_with_resource" {
  provider         = infoblox
  include_resource = true
}
