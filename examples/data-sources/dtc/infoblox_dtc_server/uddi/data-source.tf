data "infoblox_dtc_server" "get_dtc_server_using_filters" {
  filters = {
    name = "dtc-server-basic"
  }
}

data "infoblox_dtc_server" "get_dtc_server_using_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

data "infoblox_dtc_server" "get_all_dtc_servers" {}
