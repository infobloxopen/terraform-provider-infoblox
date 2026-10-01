// Retrieve specific Cloud Discovery Providers by filters
data "infoblox_cloud_discovery_provider" "by_name" {
  filters = {
    name = "example-aws-provider"
  }
}

// Retrieve specific Cloud Discovery Providers using tag filters
data "infoblox_cloud_discovery_provider" "by_tag" {
  tag_filters = {
    environment = "production"
  }
}

// Retrieve all Cloud Discovery Providers
data "infoblox_cloud_discovery_provider" "all" {}
