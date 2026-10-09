// Retrieve a specific PTR record using filters
data "infoblox_record_ptr" "get_ptr_record_using_filters" {
  filters = {
    name_in_zone = "1.0.168"
  }
}

// Retrieve specific PTR records using tag filters
data "infoblox_record_ptr" "get_ptr_record_using_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all PTR records
data "infoblox_record_ptr" "get_all_ptr_records" {}
