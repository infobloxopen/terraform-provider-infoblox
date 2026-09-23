// List specific Shared SRV Records using filters
list "infoblox_sharedrecord_srv" "list_sharedrecord_srv_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "sharedrecord_srv.example.com"
    }
  }
  limit = 10
}

// List specific Shared SRV Records using Extensible Attributes
list "infoblox_sharedrecord_srv" "list_sharedrecord_srv_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Shared SRV Records with resource details included
list "infoblox_sharedrecord_srv" "list_sharedrecord_srv_with_resource" {
  provider         = infoblox
  include_resource = true
}
