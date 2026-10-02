// Retrieve a specific NAT Group by filters
data "infoblox_nat_group" "get_natgroup_using_filters" {
  filters = {
    name = "natgroup-example"
  }
}

// Retrieve all NAT Groups
data "infoblox_nat_group" "get_all_natgroups" {}
