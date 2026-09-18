// List specific Infra Hosts using filters
list "infoblox_infra_host" "list_infra_host_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_host"
    }
  }
  limit = 10
}

// List specific Infra Hosts using Tags
list "infoblox_infra_host" "list_infra_host_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Infra Hosts with resource details included
list "infoblox_infra_host" "list_infra_host_with_resource" {
  provider         = infoblox
  include_resource = true
}
