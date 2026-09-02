// List specific Shared TXT Records using filters
list "infoblox_sharedrecord_txt" "list_sharedrecord_txt_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example-shared-record-txt"
    }
  }
  limit = 10
}

// List specific Shared TXT Records using Extensible Attributes
list "infoblox_sharedrecord_txt" "list_sharedrecord_txt_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Shared TXT Records with resource details included
list "infoblox_sharedrecord_txt" "list_sharedrecord_txt_with_resource" {
  provider         = infoblox
  include_resource = true
}
