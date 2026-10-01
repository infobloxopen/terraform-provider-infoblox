// Create an Internal Domain List with Basic Fields
resource "infoblox_internal_domain_list" "example" {
  uddi = {
    name             = "example_list"
    internal_domains = ["example.somedomain.com", "187.13.5.64/32"]
  }
}

// Create an Internal Domain List with Additional Fields
resource "infoblox_internal_domain_list" "example_with_additional_fields" {
  uddi = {
    name             = "example_list_advanced"
    internal_domains = ["example.somedomain.com", "187.13.5.64/32"]
    description      = "Example of an Internal Domain Lists"
    tags = {
      Site = "location-1"
    }
  }
}
