// List specific Response Policy Zones using filters
list "infoblox_zone_rp" "list_response_policy_zones_using_filters" {
  provider = infoblox
  config {
    filters = {
      fqdn = "example1.com"
    }
  }
  limit = 10
}

// List specific Response Policy Zones using Extensible Attributes
list "infoblox_zone_rp" "list_response_policy_zones_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Response Policy Zones with resource details included
list "infoblox_zone_rp" "list_response_policy_zones_with_resource" {
  provider         = infoblox
  include_resource = true
}
