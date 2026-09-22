data "infoblox_option_group" "example_by_name" {
  filters = {
    name = "example-group"
  }
}

data "infoblox_option_group" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

data "infoblox_option_group" "example_all" {}
