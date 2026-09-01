// Retrieve a specific DTC Pool by filters
data "infoblox_dtc_pool" "get_dtc_pool_using_filters" {
  filters = {
    name = "dtc_pool"
  }
}

// Retrieve specific DTC Pools using Extensible Attributes
data "infoblox_dtc_pool" "get_dtc_pool_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all DTC Pools
data "infoblox_dtc_pool" "get_all_dtc_pools" {}
