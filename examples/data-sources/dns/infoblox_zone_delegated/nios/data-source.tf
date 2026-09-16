// Retrieve a specific DNS zone delegated record by filters
data "infoblox_zone_delegated" "get_zone_delegated_using_filters" {
  filters = {
    fqdn = "zone-delegated.example.com"
  }
}

// Retrieve specific DNS zone delegated records using Extensible Attributes
data "infoblox_zone_delegated" "get_zone_delegated_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all DNS zone delegated zones
data "infoblox_zone_delegated" "get_all_zone_delegated" {}
