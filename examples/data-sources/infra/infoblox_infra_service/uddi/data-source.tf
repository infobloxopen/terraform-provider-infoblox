// Retrieve a specific Infra Service using filters
data "infoblox_infra_service" "get_infra_service_using_filters" {
  filters = {
    name = "example-infra-service"
  }
}

// Retrieve specific Infra Services using Tags
data "infoblox_infra_service" "get_infra_service_using_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all Infra Services
data "infoblox_infra_service" "get_all_infra_services" {}
