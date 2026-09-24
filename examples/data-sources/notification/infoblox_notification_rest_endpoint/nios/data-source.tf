// Retrieve a specific Notification REST Endpoint by filters
data "infoblox_notification_rest_endpoint" "get_notification_rest_endpoint_using_filters" {
  filters = {
    name = "example_notification_rest_endpoint"
  }
}

// Retrieve a specific Notification REST Endpoint using Extensible Attributes
data "infoblox_notification_rest_endpoint" "get_notification_rest_endpoint_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all Notification REST Endpoints
data "infoblox_notification_rest_endpoint" "get_all_notification_rest_endpoints" {}
