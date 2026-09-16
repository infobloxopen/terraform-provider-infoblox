// List specific Zone Delegateds using filters
list "infoblox_zone_delegated" "list_zone_delegated_using_filters" {
  provider = infoblox
  config {
    filters = {
      fqdn = "del.domain.com."
    }
  }
  limit = 10
}

// List specific Zone Delegateds using Tags
list "infoblox_zone_delegated" "list_zone_delegated_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Zone Delegateds with resource details included
list "infoblox_zone_delegated" "list_zone_delegated_with_resource" {
  provider         = infoblox
  include_resource = true
}
