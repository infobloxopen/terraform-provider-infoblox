// List a DHCP range by its start address
list "infoblox_range" "list_range_by_start_addr" {
  provider = infoblox
  config {
    filters = {
      start_addr = "10.0.0.10"
    }
  }
  limit = 10
}

// List DHCP ranges in a specific network view
// (the network view and its parent network must already exist, see the resource example)
list "infoblox_range" "list_range_by_network_view" {
  provider = infoblox
  config {
    filters = {
      network_view = "example-range-view"
    }
  }
}

// List DHCP ranges of a specific network
list "infoblox_range" "list_range_by_network" {
  provider = infoblox
  config {
    filters = {
      network      = "10.0.0.0/24"
      network_view = "example-range-view"
    }
  }
}

// List DHCP ranges filtered by an extensible attribute
list "infoblox_range" "list_range_by_ext_attr" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List all DHCP ranges with resource details included
list "infoblox_range" "list_range_with_resource" {
  provider         = infoblox
  include_resource = true
}
