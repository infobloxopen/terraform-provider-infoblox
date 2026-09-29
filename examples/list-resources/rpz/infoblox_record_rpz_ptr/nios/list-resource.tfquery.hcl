// List specific Record RPZ PTR Rules using filters
list "infoblox_record_rpz_ptr" "list_record_rpz_ptr_using_filters" {
  provider = infoblox
  config {
    filters = {
      ptrdname = "record1.rpz.example.com"
    }
  }
}

// List specific Record RPZ PTR Rules using Extensible Attributes
list "infoblox_record_rpz_ptr" "list_record_rpz_ptr_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Record RPZ PTR Rules with resource details included
list "infoblox_record_rpz_ptr" "list_record_rpz_ptr_with_resource" {
  provider         = infoblox
  include_resource = true
}
