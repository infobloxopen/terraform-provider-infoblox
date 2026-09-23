# Retrieve Anycast Configuration by Tag Filters
data "infoblox_anycast_config" "example_filter" {
  tag_filters = {
    tag1 = "value1"
  }
}

# Retrieve all the Anycast Configurations
data "infoblox_anycast_config" "example_all" {
}
