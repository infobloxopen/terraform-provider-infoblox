// Retrieve a specific Shared TXT Record by filters
data "infoblox_sharedrecord_txt" "get_sharedrecord_txt_using_filters" {
  filters = {
    name = "example-shared-record-txt"
  }
}

// Retrieve specific Shared TXT Records using Extensible Attributes
data "infoblox_sharedrecord_txt" "get_sharedrecord_txt_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all Shared TXT records
data "infoblox_sharedrecord_txt" "get_all_sharedrecord_txt" {}
