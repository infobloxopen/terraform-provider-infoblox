// List specific AAAA Records using filters
list "infoblox_record_aaaa" "list_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "Created by Terraform"
    }
  }
  limit = 10
}

// List specific AAAA Records using Tags
list "infoblox_record_aaaa" "list_records_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      location = "site1"
    }
  }
}

// List AAAA Records with resource details included
list "infoblox_record_aaaa" "list_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
