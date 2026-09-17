// Retrieve SVCB records filtered by an attribute
data "infoblox_record_svcb" "example_by_attribute" {
  filters = {
    "absolute_name_spec" = "record.example.com."
  }
}

// Retrieve SVCB records filtered by tag
data "infoblox_record_svcb" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all SVCB records
data "infoblox_record_svcb" "example_all" {}
