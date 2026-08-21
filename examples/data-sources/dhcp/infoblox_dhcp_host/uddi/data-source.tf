// Retrieve DHCP hosts filtered by IP space
data "infoblox_dhcp_host" "by_ip_space" {
  filters = {
    "uddi.ip_space" = "ipam/ip_space/5202ccf2-f3b6-11ed-a04c-0acb29431b1f"
  }
}

// Retrieve DHCP hosts filtered by tag
data "infoblox_dhcp_host" "by_tag" {
  tag_filters = {
    "host/host_ip" = "10.39.49.37"
  }
}

// Retrieve all DHCP hosts
data "infoblox_dhcp_host" "all" {}
