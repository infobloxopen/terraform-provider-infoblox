// List specific Access Codes using filters
list "infoblox_access_code" "list_access_code_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "Created by Terraform"
    }
  }
  limit = 10
}

// List specific Access Codes using Tags
list "infoblox_access_code" "list_access_code_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Access Codes with resource details included
list "infoblox_access_code" "list_access_code_with_resource" {
  provider         = infoblox
  include_resource = true
}
