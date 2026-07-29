// List specific CAA Records using filters
list "infoblox_record_caa" "list_caa_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "CAA Record created by Terraform"
    }
  }
  limit = 10
}

// List specific CAA Records using Tags
list "infoblox_record_caa" "list_caa_records_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List CAA Records with resource details included
list "infoblox_record_caa" "list_caa_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
