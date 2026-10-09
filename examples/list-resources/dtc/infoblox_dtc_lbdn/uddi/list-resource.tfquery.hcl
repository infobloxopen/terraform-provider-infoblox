// List specific Dtc Lbdns using filters
list "infoblox_dtc_lbdn" "list_dtc_lbdn_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "Created by Terraform"
    }
  }
  limit = 10
}

// List specific Dtc Lbdns using Tags
list "infoblox_dtc_lbdn" "list_dtc_lbdn_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Dtc Lbdns with resource details included
list "infoblox_dtc_lbdn" "list_dtc_lbdn_with_resource" {
  provider         = infoblox
  include_resource = true
}
