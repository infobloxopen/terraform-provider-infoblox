// Create a Category Filter with Basic Fields
resource "infoblox_category_filter" "create_category_filter" {
  uddi = {
    name       = "example_category_filter"
    categories = ["College"]
  }
}

// Create a Category Filter with Additional Fields
resource "infoblox_category_filter" "create_category_filter_with_additional_fields" {
  uddi = {
    name        = "example_category_filter_advanced"
    categories  = ["College", "Tutoring"]
    description = "Category filter for education-related content"

    tags = {
      Site = "location-1"
    }
  }
}
