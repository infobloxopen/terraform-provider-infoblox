// List Admin Users using name filter
list "infoblox_admin_user" "list_by_name" {
  provider = infoblox
  config {
    filters = {
      name = "tf-admin-example"
    }
  }
  limit = 10
}

// List Admin Users using Extensible Attributes
list "infoblox_admin_user" "list_by_ext_attrs" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "us-west"
    }
  }
}

// List all Admin Users with resource details
list "infoblox_admin_user" "list_all" {
  provider         = infoblox
  include_resource = true
}
