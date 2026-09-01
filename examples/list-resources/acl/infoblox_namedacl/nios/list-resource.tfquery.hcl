// List specific Named ACLs using filters
list "infoblox_namedacl" "list_namedacls_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_namedacl"
    }
  }
}

// List specific Named ACLs using Extensible Attributes
list "infoblox_namedacl" "list_namedacls_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Named ACLs with resource details included
list "infoblox_namedacl" "list_namedacls_with_resource" {
  provider         = infoblox
  include_resource = true
}
