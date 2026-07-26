// List specific DNS Views using filters
list "infoblox_view" "list_dns_views_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_dns_view"
    }
  }
}

// List specific DNS Views using tag filters
list "infoblox_view" "list_dns_views_using_tag_filters" {
  provider = infoblox
  config {
    tag_filters = {
      site = "Site A"
    }
  }
}

// List DNS Views with resource details included
list "infoblox_view" "list_dns_views_with_resource" {
  provider         = infoblox
  include_resource = true
}
