// Retrieve a specific DTC Server using filters
data "infoblox_dtc_server" "get_dtc_server_using_filters" {
  filters = {
    name = "dtc-server-basic"
  }
}

// Retrieve specific DTC Servers using tag filters
data "infoblox_dtc_server" "get_dtc_server_using_tag_filters" {
  tag_filters = {
    Site = "us-east-1"
  }
}

// Retrieve all DTC Servers
data "infoblox_dtc_server" "get_all_dtc_servers" {}
