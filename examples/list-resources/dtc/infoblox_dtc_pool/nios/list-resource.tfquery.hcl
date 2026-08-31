// List specific DTC Pool records using filters
list "infoblox_dtc_pool" "list_dtc_pool_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "dtc_pool"
    }
  }
  limit = 10
}

// List specific DTC Pool records using Extensible Attributes
list "infoblox_dtc_pool" "list_dtc_pool_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List DTC Pool records with resource details included
list "infoblox_dtc_pool" "list_dtc_pool_with_resource" {
  provider         = infoblox
  include_resource = true
}
