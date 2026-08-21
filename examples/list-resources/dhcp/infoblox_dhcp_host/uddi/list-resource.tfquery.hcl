// List DHCP hosts by IP space
list "infoblox_dhcp_host" "by_ip_space" {
  provider = infoblox
  config {
    filters = {
      "uddi.ip_space" = "ipam/ip_space/5202ccf2-f3b6-11ed-a04c-0acb29431b1f"
    }
  }
}

// List DHCP hosts by tag
list "infoblox_dhcp_host" "by_tag" {
  provider = infoblox
  config {
    tag_filters = {
      "host/host_ip" = "10.39.49.37"
    }
  }
}

// List DHCP hosts with full resource details
list "infoblox_dhcp_host" "with_resource" {
  provider         = infoblox
  include_resource = true
}
