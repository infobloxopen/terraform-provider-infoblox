// Retrieve a specific Upgrade Group by filters
data "infoblox_upgradegroup" "get_upgradegroup_using_filters" {
  filters = {
    name = "example-upgradegroup"
  }
}

// Retrieve all Upgrade Groups
data "infoblox_upgradegroup" "get_all_upgradegroups" {}
