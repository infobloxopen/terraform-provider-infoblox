// List specific Access Codes using filters
list "infoblox_access_code" "list_access_code_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example-access-code"
    }
  }
  limit = 10
}

// List all Access Codes
list "infoblox_access_code" "list_all_access_codes" {
  provider = infoblox
}

// List Access Codes with resource details included
list "infoblox_access_code" "list_access_code_with_resource" {
  provider         = infoblox
  include_resource = true
}
