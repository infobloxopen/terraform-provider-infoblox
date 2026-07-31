// List specific CNAME Records using filters
list "infoblox_record_cname" "list_cname_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "alias.example.com"
    }
  }
}

// List specific CNAME Records using Extensible Attributes
list "infoblox_record_cname" "list_cname_records_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List CNAME Records with resource details included
list "infoblox_record_cname" "list_cname_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
