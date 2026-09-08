// List specific PTR Records using filters
list "infoblox_record_ptr" "list_ptr_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "Created by Terraform"
    }
  }
  limit = 10
}

// List specific PTR Records using Tags
list "infoblox_record_ptr" "list_ptr_records_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List PTR Records with resource details included
list "infoblox_record_ptr" "list_ptr_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
