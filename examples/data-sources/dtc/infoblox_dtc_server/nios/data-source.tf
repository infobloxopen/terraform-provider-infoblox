data "infoblox_dtc_server" "get_dtc_server_using_filters" {
  filters = {
    name = "dtc-server-basic"
  }
}

data "infoblox_dtc_server" "get_dtc_server_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

data "infoblox_dtc_server" "get_all_dtc_servers" {}
