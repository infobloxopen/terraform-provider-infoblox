// Get "HTTPS" records filtered by an attribute
data "infoblox_record_https" "example_by_attribute" {
  filters = {
    "absolute_name_spec" = "abc.example.com"
  }
}

// Get "HTTPS" records filtered by tag
data "infoblox_record_https" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Get all "HTTPS" records
data "infoblox_record_https" "example_all" {}
