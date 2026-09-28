// List specific Notification REST Endpoints using filters
list "infoblox_notification_rest_endpoint" "list_notification_rest_endpoints_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_notification_rest_endpoint"
    }
  }
  limit = 10
}

// List specific Notification REST Endpoints using Extensible Attributes
list "infoblox_notification_rest_endpoint" "list_notification_rest_endpoints_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Notification REST Endpoints with resource details included
list "infoblox_notification_rest_endpoint" "list_notification_rest_endpoints_with_resource" {
  provider         = infoblox
  include_resource = true
}
