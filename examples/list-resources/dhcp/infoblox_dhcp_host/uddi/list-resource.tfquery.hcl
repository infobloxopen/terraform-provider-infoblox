// List DHCP hosts by name
list "infoblox_dhcp_host" "by_name" {
  provider = infoblox
  config {
    filters = {
      "uddi.name" = "example-host"
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
