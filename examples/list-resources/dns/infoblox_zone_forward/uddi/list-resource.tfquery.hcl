// List specific Zone Forwards using filters
list "infoblox_zone_forward" "list_zone_forward_using_filters" {
  provider = infoblox
  config {
    filters = {
      fqdn = "domain.com."
    }
  }
  limit = 10
}

// List specific Zone Forwards using Tags
list "infoblox_zone_forward" "list_zone_forward_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Zone Forwards with resource details included
list "infoblox_zone_forward" "list_zone_forward_with_resource" {
  provider         = infoblox
  include_resource = true
}
