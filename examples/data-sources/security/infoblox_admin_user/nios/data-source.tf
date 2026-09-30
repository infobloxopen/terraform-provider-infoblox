// Retrieve an Admin User by name
data "infoblox_admin_user" "by_name" {
  filters = {
    name = "tf-admin-example"
  }
}

// Retrieve Admin Users using Extensible Attributes
data "infoblox_admin_user" "by_ext_attrs" {
  ext_attr_filters = {
    Site = "us-west"
  }
}

// Retrieve all Admin Users
data "infoblox_admin_user" "all" {}
