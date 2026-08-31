// Retrieve a specific NAT Group by filters
data "infoblox_natgroup" "get_natgroup_using_filters" {
  filters = {
    name = "natgroup-example"
  }
}

// Retrieve all NAT Groups
data "infoblox_natgroup" "get_all_natgroups" {}
