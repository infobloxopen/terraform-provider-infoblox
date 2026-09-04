// Create NS group Forward Stub Server with Basic Fields
resource "infoblox_nsgroup_forwardstubserver" "nsgroup_forward_stub_server_with_basic_fields" {
  nios = {
    name = "example_ns_group_forward_stub_server"
    external_servers = [
      {
        name    = "example.com"
        address = "2.3.4.4"
      }
    ]
  }
}

// Create NS Group Forward Stub Server with Additional Fields
resource "infoblox_nsgroup_forwardstubserver" "nsgroup_forward_stub_server_with_additional_fields" {
  nios = {
    name = "example_ns_group_forward_stub_server_additional_fields"
    external_servers = [
      {
        name    = "example.com"
        address = "2.3.4.4"
      }
    ]
    // Additional Fields
    comment = "Example NS Group Forward Stub Server"
    ext_attrs = {
      Site = "location-1"
    }
  }
}
