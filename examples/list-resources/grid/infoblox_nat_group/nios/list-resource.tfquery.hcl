// List specific NAT Groups using filters
list "infoblox_nat_group" "list_natgroups_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "natgroup.example"
    }
  }
}

// List NAT Groups with resource details included
list "infoblox_nat_group" "list_natgroups_with_resource" {
  provider         = infoblox
  include_resource = true
}
