// Create a Notification REST Endpoint to use as the notification target
resource "infoblox_notification_rest_endpoint" "example" {
  nios = {
    name                 = "example_notification_rest_endpoint"
    outbound_member_type = "GM"
    uri                  = "https://example.com/notify"
  }
}

// Create a Notification Rule with Basic Fields
resource "infoblox_notification_rule" "notification_rule_with_basic_fields" {
  nios = {
    name                = "example_notification_rule"
    event_type          = "DHCP_LEASES"
    notification_action = "RESTAPI_TEMPLATE_INSTANCE"
    notification_target = infoblox_notification_rest_endpoint.example.id
    template_instance = {
      template = "DHCP_Lease"
    }
    expression_list = [
      { op = "AND", op1_type = "LIST" },
      { op = "EQ", op1 = "DHCP_LEASE_STATE", op1_type = "FIELD", op2 = "DHCP_LEASE_STATE_ACTIVE", op2_type = "STRING" },
      { op = "ENDLIST" },
    ]
  }
}

// Create a Notification Rule with Additional Fields
resource "infoblox_notification_rule" "notification_rule_with_additional_fields" {
  nios = {
    name                = "example_notification_rule_additional"
    event_type          = "DHCP_LEASES"
    notification_action = "RESTAPI_TEMPLATE_INSTANCE"
    notification_target = infoblox_notification_rest_endpoint.example.id
    template_instance = {
      template = "DHCP_Lease"
    }
    expression_list = [
      { op = "AND", op1_type = "LIST" },
      { op = "EQ", op1 = "DHCP_LEASE_STATE", op1_type = "FIELD", op2 = "DHCP_LEASE_STATE_ACTIVE", op2_type = "STRING" },
      { op = "ENDLIST" },
    ]
    comment = "Notify on active DHCP leases"
    disable = false
  }
}
