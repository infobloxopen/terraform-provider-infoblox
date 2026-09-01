// List specific SVCB Records using filters
list "infoblox_record_svcb" "list_record_svcb_using_filters" {
  provider = infoblox
  config {
    filters = {
      absolute_name_spec = "record.example.com."
    }
  }
  limit = 10
}

// List specific SVCB Records using Tags
list "infoblox_record_svcb" "list_record_svcb_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List SVCB Records with resource details included
list "infoblox_record_svcb" "list_record_svcb_with_resource" {
  provider         = infoblox
  include_resource = true
}
