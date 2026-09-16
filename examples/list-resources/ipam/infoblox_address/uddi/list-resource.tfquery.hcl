// List specific Addresses using filters
list "infoblox_address" "list_address_using_filters" {
  provider = infoblox
  config {
    filters = {
      address = "10.1.0.5"
    }
  }
  limit = 10
}

// List specific Addresses using Tags
list "infoblox_address" "list_address_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Addresses with resource details included
list "infoblox_address" "list_address_with_resource" {
  provider         = infoblox
  include_resource = true
}
