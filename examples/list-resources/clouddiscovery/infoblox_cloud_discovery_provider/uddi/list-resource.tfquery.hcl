// List specific Cloud Discovery Providers using filters
list "infoblox_cloud_discovery_provider" "list_cloud_discovery_provider_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "Created by Terraform"
    }
  }
  limit = 10
}

// List specific Cloud Discovery Providers using Tags
list "infoblox_cloud_discovery_provider" "list_cloud_discovery_provider_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Cloud Discovery Providers with resource details included
list "infoblox_cloud_discovery_provider" "list_cloud_discovery_provider_with_resource" {
  provider         = infoblox
  include_resource = true
}
