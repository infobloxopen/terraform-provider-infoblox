// List specific DNS Views using filters
list "infoblox_view" "list_dns_views_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "default"
    }
  }
}

// List specific DNS Views using Extensible Attributes
list "infoblox_view" "list_dns_views_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List DNS Views with resource details included
list "infoblox_view" "list_dns_views_with_resource" {
  provider         = infoblox
  include_resource = true
}
