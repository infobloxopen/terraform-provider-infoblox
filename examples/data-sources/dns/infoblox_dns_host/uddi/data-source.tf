// Get all DNS Hosts
data "infoblox_dns_host" "all_hosts" {}

// Get DNS Host by name
data "infoblox_dns_host" "dns_host_by_name" {
  filters = {
    name = "dns_host_by_name"
  }
}

// Get DNS Hosts with the specific tags
data "infoblox_dns_host" "all_dns_hosts_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}
