data "infoblox_custom_redirect" "by_filters" {
  filters = {
    name = "example_custom_redirect"
  }
}

data "infoblox_custom_redirect" "by_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

data "infoblox_custom_redirect" "all" {}
