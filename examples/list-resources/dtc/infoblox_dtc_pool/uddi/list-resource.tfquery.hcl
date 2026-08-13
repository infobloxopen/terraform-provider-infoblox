// List specific DtcPool records using filters
list "infoblox_dtc_pool" "list_dtc_pool_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "Created by Terraform"
    }
  }
  limit = 10
}

// List specific DtcPool records using Tags
list "infoblox_dtc_pool" "list_dtc_pool_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List DtcPool records with resource details included
list "infoblox_dtc_pool" "list_dtc_pool_with_resource" {
  provider         = infoblox
  include_resource = true
}
