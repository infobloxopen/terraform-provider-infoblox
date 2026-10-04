// List specific DNS Hosts using filters
list "infoblox_dns_host" "list_dns_host_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "dns_host_by_name"
    }
  }
  limit = 10
}

// List specific DNS Hosts using Tags
list "infoblox_dns_host" "list_dns_host_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List DNS Hosts with resource details included
list "infoblox_dns_host" "list_dns_host_with_resource" {
  provider         = infoblox
  include_resource = true
}
