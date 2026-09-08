// Retrieve a specific Named ACL by filters
data "infoblox_namedacl" "get_namedacl_using_filters" {
  filters = {
    name = "example_namedacl"
  }
}

// Retrieve specific Named ACLs using Extensible Attributes
data "infoblox_namedacl" "get_namedacl_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all Named ACLs
data "infoblox_namedacl" "get_all_namedacls" {}
