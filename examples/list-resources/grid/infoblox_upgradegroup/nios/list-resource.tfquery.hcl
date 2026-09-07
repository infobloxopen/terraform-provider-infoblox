// List specific Upgrade Groups using filters
list "infoblox_upgradegroup" "list_upgradegroups_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example-upgradegroup"
    }
  }
}

// List Upgrade Groups with resource details included
list "infoblox_upgradegroup" "list_upgradegroups_with_resource" {
  provider         = infoblox
  include_resource = true
}
