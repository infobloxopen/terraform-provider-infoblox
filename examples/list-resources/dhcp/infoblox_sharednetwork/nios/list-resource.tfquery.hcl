// List shared networks by name
list "infoblox_sharednetwork" "list_sharednetwork_by_name" {
  provider = infoblox
  config {
    filters = {
      name = "example_shared_network"
    }
  }
  limit = 10
}

// List shared networks in a specific network view
list "infoblox_sharednetwork" "list_sharednetwork_by_network_view" {
  provider = infoblox
  config {
    filters = {
      network_view = "default"
    }
  }
}

// List shared networks filtered by an extensible attribute
list "infoblox_sharednetwork" "list_sharednetwork_by_ext_attr" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "us-east-1"
    }
  }
}

// List all shared networks with resource details included
list "infoblox_sharednetwork" "list_sharednetwork_with_resource" {
  provider         = infoblox
  include_resource = true
}
