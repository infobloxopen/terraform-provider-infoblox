// Retrieve a specific Upgrade Group by filters
data "infoblox_upgrade_group" "get_upgradegroup_using_filters" {
  filters = {
    name = "example-upgradegroup"
  }
}

// Retrieve all Upgrade Groups
data "infoblox_upgrade_group" "get_all_upgradegroups" {}
