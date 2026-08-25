// Retrieve a specific DTC Server using filters
data "infoblox_dtc_server" "get_dtc_server_using_filters" {
  filters = {
    name = "dtc-server-basic"
  }
}

// Retrieve specific DTC Servers using Extensible Attributes
data "infoblox_dtc_server" "get_dtc_server_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "us-east-1"
  }
}

// Retrieve all DTC Servers
data "infoblox_dtc_server" "get_all_dtc_servers" {}
