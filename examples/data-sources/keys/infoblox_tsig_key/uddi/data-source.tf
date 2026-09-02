// Retrieve a specific TSIG Key using filters
data "infoblox_tsig_key" "get_tsig_key_using_filters" {
  filters = {
    name = "tsig-key-basic.example.com."
  }
}

// Retrieve specific TSIG Keys using Tags
data "infoblox_tsig_key" "get_tsig_key_using_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all TSIG Keys
data "infoblox_tsig_key" "get_all_tsig_keys" {}
