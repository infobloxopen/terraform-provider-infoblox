// Retrieve specific HA Groups using filters
data "infoblox_ha_group" "get_by_name" {
  filters = {
    name = "example-ha-group"
  }
}

// Retrieve specific HA Groups using tag filters
data "infoblox_ha_group" "get_by_tag" {
  tag_filters = {
    location = "site-1"
  }
}

// Retrieve all HA Groups
data "infoblox_ha_group" "get_all" {}
