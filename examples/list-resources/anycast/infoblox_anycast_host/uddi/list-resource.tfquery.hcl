// List specific Anycast Hosts using filters
list "infoblox_anycast_host" "list_anycast_host_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "Created by Terraform"
    }
  }
  limit = 10
}

// List specific Anycast Hosts using Tags
list "infoblox_anycast_host" "list_anycast_host_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Anycast Hosts with resource details included
list "infoblox_anycast_host" "list_anycast_host_with_resource" {
  provider         = infoblox
  include_resource = true
}
