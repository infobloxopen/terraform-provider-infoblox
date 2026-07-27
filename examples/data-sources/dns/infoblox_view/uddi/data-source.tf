// Retrieve a specific DNS view by filters
data "infoblox_view" "get_view_using_filters" {
  filters = {
    name = "example_dns_view"
  }
}

// Retrieve DNS views filtered by tag
data "infoblox_view" "get_view_using_tag_filters" {
  tag_filters = {
    site = "Site A"
  }
}

// Retrieve DNS views filtered by a NIOS-imported tag
data "infoblox_view" "get_view_using_nios_tag" {
  tag_filters = {
    "nios/imported" = "true"
  }
}

// Retrieve all DNS views
data "infoblox_view" "get_all_views" {}
