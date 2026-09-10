// Create a Notification REST Endpoint with Basic Fields
resource "infoblox_notification_rest_endpoint" "notification_rest_endpoint_with_basic_fields" {
  nios = {
    name                 = "example_notification_rest_endpoint"
    outbound_member_type = "GM"
    uri                  = "https://example.com/notify"
  }
}

// Create a Notification REST Endpoint with Additional Fields
resource "infoblox_notification_rest_endpoint" "notification_rest_endpoint_with_additional_fields" {
  nios = {
    name                   = "example_notification_rest_endpoint_additional"
    outbound_member_type   = "GM"
    uri                    = "https://example.com/notify"
    comment                = "Comment for Notification REST Endpoint"
    log_level              = "DEBUG"
    server_cert_validation = "NO_VALIDATION"
    sync_disabled          = false
    timeout                = 60
    username               = "api_user"
    password               = "api_password"
    vendor_identifier      = "WAPI"

    ext_attrs = {
      Site = "location-1"
    }
  }
}
