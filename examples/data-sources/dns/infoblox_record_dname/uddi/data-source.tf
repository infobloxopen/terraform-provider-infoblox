// Retrieve specific DNAME records using filters
data "infoblox_record_dname" "get_dname_record_using_filters" {
  filters = {
    "name_in_zone" = "dname"
  }
}

// Retrieve specific DNAME records using tags
data "infoblox_record_dname" "get_dname_record_using_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all DNAME records
data "infoblox_record_dname" "get_all_dname_records" {}
