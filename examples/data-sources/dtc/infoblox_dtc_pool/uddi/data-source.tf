// Retrieve a specific DTC Pool by filters
data "infoblox_dtc_pool" "get_dtc_pool_using_filters" {
  filters = {
    name = "dtc_pool"
  }
}

// Retrieve specific DTC Pools using Tag Filters
data "infoblox_dtc_pool" "get_dtc_pool_using_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all DTC Pools
data "infoblox_dtc_pool" "get_all_dtc_pools" {}
