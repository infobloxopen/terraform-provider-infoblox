list "infoblox_option_group" "list_option_group_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example-group"
    }
  }
  limit = 10
}

list "infoblox_option_group" "list_option_group_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

list "infoblox_option_group" "list_option_group_with_resource" {
  provider         = infoblox
  include_resource = true
}
