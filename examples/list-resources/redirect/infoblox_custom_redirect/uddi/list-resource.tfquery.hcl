// List specific Custom Redirects using filters
list "infoblox_custom_redirect" "list_custom_redirect_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "Created by Terraform"
    }
  }
  limit = 10
}

// List specific Custom Redirects using Tags
list "infoblox_custom_redirect" "list_custom_redirect_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Custom Redirects with resource details included
list "infoblox_custom_redirect" "list_custom_redirect_with_resource" {
  provider         = infoblox
  include_resource = true
}
