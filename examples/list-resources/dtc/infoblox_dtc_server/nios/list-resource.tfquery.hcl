// List specific DTC Servers using filters
list "infoblox_dtc_server" "list_dtc_server_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "dtc-server-basic"
    }
  }
  limit = 10
}

// List specific DTC Servers using Extensible Attributes
list "infoblox_dtc_server" "list_dtc_server_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "us-east-1"
    }
  }
}

// List DTC Servers with resource details included
list "infoblox_dtc_server" "list_dtc_server_with_resource" {
  provider         = infoblox
  include_resource = true
}
