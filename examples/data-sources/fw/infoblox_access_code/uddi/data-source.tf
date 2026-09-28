// Retrieve a specific Access Code by name
data "infoblox_access_code" "get_by_name" {
  filters = {
    name = "example-access-code"
  }
}

// Retrieve all Access Codes
data "infoblox_access_code" "get_all" {}
