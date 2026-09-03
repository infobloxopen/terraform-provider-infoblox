// List DNS NS Group Forward Stub Servers using filters
list "infoblox_nsgroup_forwardstubserver" "list_ns_group_forward_stub_server_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_ns_group_forward_stub_server"
    }
  }
  limit = 10
}

// List DNS NS Group Forward Stub Servers using Extensible Attributes
list "infoblox_nsgroup_forwardstubserver" "list_ns_group_forward_stub_server_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List DNS NS Group Forward Stub Servers with resource details included
list "infoblox_nsgroup_forwardstubserver" "list_ns_group_forward_stub_server_with_resource" {
  provider         = infoblox
  include_resource = true
}
