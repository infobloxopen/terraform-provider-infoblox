// Retrieve a specific Named ACL by filters
data "infoblox_namedacl" "get_namedacl_using_filters" {
  filters = {
    name = "example_namedacl"
  }
}

// Retrieve specific Named ACLs using Tags
data "infoblox_namedacl" "get_namedacl_using_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all Named ACLs
data "infoblox_namedacl" "get_all_namedacls" {}
