// List specific A Records using filters
list "infoblox_record_a" "list_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "Created by Terraform"
    }
  }
  limit = 10
}

// List specific A Records using Tags
list "infoblox_record_a" "list_records_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List A Records with resource details included
list "infoblox_record_a" "list_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
