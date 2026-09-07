// List Super Hosts using filters
list "infoblox_superhost" "list_super_hosts_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_super_host"
    }
  }
  limit = 10
}

// List Super Hosts using Extensible Attributes
list "infoblox_superhost" "list_super_hosts_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List all Super Hosts with resource details
list "infoblox_superhost" "list_super_hosts_with_resource" {
  provider         = infoblox
  include_resource = true
}
