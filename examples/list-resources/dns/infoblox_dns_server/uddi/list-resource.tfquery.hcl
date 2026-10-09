// List specific Dns Servers using filters
list "infoblox_dns_server" "list_dns_server_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_server"
    }
  }
  limit = 10
}

// List specific Dns Servers using Tags
list "infoblox_dns_server" "list_dns_server_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Dns Servers with resource details included
list "infoblox_dns_server" "list_dns_server_with_resource" {
  provider         = infoblox
  include_resource = true
}
