// List specific DNAME Records using filters
list "infoblox_record_dname" "list_dname_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      name_in_zone = "dname"
    }
  }
  limit = 10
}

// List specific DNAME Records using Tags
list "infoblox_record_dname" "list_dname_records_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List DNAME Records with resource details included
list "infoblox_record_dname" "list_dname_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
