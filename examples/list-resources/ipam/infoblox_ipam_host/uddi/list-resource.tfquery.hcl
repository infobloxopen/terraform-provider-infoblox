// List specific Ipam Hosts using filters
list "infoblox_ipam_host" "list_ipam_host_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "Created by Terraform"
    }
  }
  limit = 10
}

// List specific Ipam Hosts using Tags
list "infoblox_ipam_host" "list_ipam_host_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Ipam Hosts with resource details included
list "infoblox_ipam_host" "list_ipam_host_with_resource" {
  provider         = infoblox
  include_resource = true
}
