// List specific Auth Nsgs using filters
list "infoblox_auth_nsg" "list_auth_nsg_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_auth_nsg"
    }
  }
  limit = 10
}

// List specific Auth Nsgs using Tags
list "infoblox_auth_nsg" "list_auth_nsg_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Auth Nsgs with resource details included
list "infoblox_auth_nsg" "list_auth_nsg_with_resource" {
  provider         = infoblox
  include_resource = true
}
