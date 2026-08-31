// List specific Record Svcbs using filters
list "infoblox_record_svcb" "list_record_svcb_using_filters" {
  provider = infoblox
  config {
    filters = {
      absolute_name_spec = "abc.example.com"
    }
  }
  limit = 10
}

// List specific Record Svcbs using Tags
list "infoblox_record_svcb" "list_record_svcb_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Record Svcbs with resource details included
list "infoblox_record_svcb" "list_record_svcb_with_resource" {
  provider         = infoblox
  include_resource = true
}
