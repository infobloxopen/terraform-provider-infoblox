// Get "SVCB" records filtered by an attribute
data "infoblox_record_svcb" "example_by_attribute" {
  filters = {
    "absolute_name_spec" = "abc.example.com"
  }
}

// Get "SVCB" records filtered by tag
data "infoblox_record_svcb" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Get all "SVCB" records
data "infoblox_record_svcb" "example_all" {}
