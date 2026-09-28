# Retrieve Infra Host with Anycast Service Configured ( Required as Parent )
data "infoblox_infra_host" "parent" {
  filters = {
    display_name = "my_host"
  }
}

# Manage Anycast Configuration with basic fields
resource "infoblox_anycast_config" "example" {
  uddi = {
    name               = "anycast_config_example"
    service            = "DNS"
    anycast_ip_address = "192.1.1.1"
  }
}

# Manage Anycast Configuration with additional fields
resource "infoblox_anycast_config" "example_advanced" {
  uddi = {
    name               = "anycast_config_example2"
    service            = "DNS"
    anycast_ip_address = "192.2.2.2"

    # Other Optional Fields
    onprem_hosts = [
      {
        id   = data.infoblox_infra_host.parent.results[0].uddi.legacy_id
        name = data.infoblox_infra_host.parent.results[0].uddi.display_name
      }
    ]
    description = "anycast configuration example"
    tags = {
      Site = "location-1"
    }
  }
}
