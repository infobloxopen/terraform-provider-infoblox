// Retrieve a specific SRV record by filters
data "infoblox_record_srv" "get_srv_record_using_filters" {
  filters = {
    "name_in_zone" = "record_srv"
  }
}

// Retrieve specific SRV records using Tag Filters
data "infoblox_record_srv" "get_srv_record_using_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all SRV records
data "infoblox_record_srv" "get_all_srv_records" {}
