// Retrieve a specific NS Group Forward Stub Server by filters
data "infoblox_nsgroup_forwardstubserver" "get_nsgroup_forward_stub_server_using_filters" {
  filters = {
    name = "example_ns_group_forward_stub_server"
  }
}

// Retrieve specific NS Group Forward Stub Servers using Extensible Attributes
data "infoblox_nsgroup_forwardstubserver" "get_nsgroup_forward_stub_server_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all NS Groups Forward Stub Servers
data "infoblox_nsgroup_forwardstubserver" "get_all_nsgroup_forward_stub_servers" {
}
