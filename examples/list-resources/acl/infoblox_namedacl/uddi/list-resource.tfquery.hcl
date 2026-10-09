// List specific Named ACLs using filters
list "infoblox_namedacl" "list_namedacls_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_namedacl"
    }
  }
  limit = 10
}

// List specific Named ACLs using Tags
list "infoblox_namedacl" "list_namedacls_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Named ACLs with resource details included
list "infoblox_namedacl" "list_namedacls_with_resource" {
  provider         = infoblox
  include_resource = true
}
