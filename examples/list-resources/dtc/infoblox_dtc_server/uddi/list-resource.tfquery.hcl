// List specific Dtc Servers using filters
list "infoblox_dtc_server" "list_dtc_server_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "dtc-server-basic"
    }
  }
  limit = 10
}

// List specific Dtc Servers using Tags
list "infoblox_dtc_server" "list_dtc_server_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "us-east-1"
    }
  }
}

// List Dtc Servers with resource details included
list "infoblox_dtc_server" "list_dtc_server_with_resource" {
  provider         = infoblox
  include_resource = true
}
