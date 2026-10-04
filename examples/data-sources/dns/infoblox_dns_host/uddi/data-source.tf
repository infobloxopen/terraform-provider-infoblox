// Retrieve DNS Host via filters
data "infoblox_dns_host" "dns_host_by_name" {
  filters = {
    name = "dns_host_by_name"
  }
}

// Retrieve DNS Hosts with the specific tags
data "infoblox_dns_host" "all_dns_hosts_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all DNS Hosts
data "infoblox_dns_host" "all_hosts" {}
