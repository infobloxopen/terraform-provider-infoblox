// List specific Service Restart Groups using filters
list "infoblox_servicerestart_group" "list_servicerestart_groups_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "servicerestart_group.example"
    }
  }
  limit = 10
}

// List specific Service Restart Groups using Extensible Attributes
list "infoblox_servicerestart_group" "list_servicerestart_groups_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Service Restart Groups with resource details included
list "infoblox_servicerestart_group" "list_servicerestart_groups_with_resource" {
  provider         = infoblox
  include_resource = true
}
