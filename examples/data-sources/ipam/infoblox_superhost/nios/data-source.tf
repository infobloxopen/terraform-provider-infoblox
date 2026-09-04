// Retrieve a specific Super Host by filters
data "infoblox_superhost" "get_super_host_using_filters" {
  filters = {
    name = "example_super_host"
  }
}

// Retrieve specific Super Hosts using Extensible Attributes
data "infoblox_superhost" "get_super_host_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all Super Hosts
data "infoblox_superhost" "get_all_super_hosts" {}
