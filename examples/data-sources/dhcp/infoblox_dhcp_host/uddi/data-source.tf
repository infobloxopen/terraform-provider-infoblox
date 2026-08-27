// Retrieve DHCP hosts filtered by name
data "infoblox_dhcp_host" "by_name" {
  filters = {
    "uddi.name" = "example-host"
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
