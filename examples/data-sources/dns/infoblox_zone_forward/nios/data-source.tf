// Retrieve a specific DNS Zone Forward by filters
data "infoblox_zone_forward" "get_zone_forward_using_filters" {
  filters = {
    fqdn = "zone-forward1.example.com"
  }
}

// Retrieve specific DNS Zone Forward using Extensible Attributes
data "infoblox_zone_forward" "get_zone_forward_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all DNS Zone Forward
data "infoblox_zone_forward" "get_all_zone_forward" {}
