// List specific TSIG Keys using filters
list "infoblox_tsig_key" "list_tsig_key_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "TSIG Key created by Terraform"
    }
  }
  limit = 10
}

// List specific TSIG Keys using Tags
list "infoblox_tsig_key" "list_tsig_key_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List TSIG Keys with resource details included
list "infoblox_tsig_key" "list_tsig_key_with_resource" {
  provider         = infoblox
  include_resource = true
}
