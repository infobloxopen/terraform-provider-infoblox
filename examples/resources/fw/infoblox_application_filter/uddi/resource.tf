// Create an Application Filter matching by application name
resource "infoblox_application_filter" "example_by_name" {
  uddi = {
    name     = "example-app-filter-by-name"
    criteria = [{ name = "Microsoft 365" }]
  }
}

// Create an Application Filter matching by category
resource "infoblox_application_filter" "example_by_category" {
  uddi = {
    name        = "example-app-filter-by-category"
    description = "Filter for all email applications"
    criteria    = [{ category = "Email" }]
    tags = {
      Site = "location-1"
    }
  }
}
