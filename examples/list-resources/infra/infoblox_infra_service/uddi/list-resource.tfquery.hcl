// List specific Infra Services using filters
list "infoblox_infra_service" "list_infra_service_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_service"
    }
  }
  limit = 10
}

// List specific Infra Services using Tags
list "infoblox_infra_service" "list_infra_service_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Infra Services with resource details included
list "infoblox_infra_service" "list_infra_service_with_resource" {
  provider         = infoblox
  include_resource = true
}
