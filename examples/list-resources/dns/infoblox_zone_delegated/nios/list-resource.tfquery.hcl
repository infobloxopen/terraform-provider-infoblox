// List specific Delegated Zones using filters
list "infoblox_zone_delegated" "list_delegated_zones_using_filters" {
  provider = infoblox
  config {
    filters = {
      fqdn = "zone-delegated.example_auth.com"
    }
  }
  limit = 10
}

// List specific Delegated Zones using Extensible Attributes
list "infoblox_zone_delegated" "list_delegated_zones_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Delegated Zones with resource details included
list "infoblox_zone_delegated" "list_delegated_zones_with_resource" {
  provider         = infoblox
  include_resource = true
}
