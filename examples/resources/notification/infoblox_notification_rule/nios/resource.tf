// Create a Notification REST Endpoint (Required as Parent)
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
      {
        op       = "EQ"
        op1      = "DHCP_LEASE_STATE"
        op1_type = "FIELD"
        op2      = "DHCP_LEASE_STATE_ACTIVE"
        op2_type = "STRING"
      },
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
      {
        op       = "EQ"
        op1      = "DHCP_LEASE_STATE"
        op1_type = "FIELD"
        op2      = "DHCP_LEASE_STATE_ACTIVE"
        op2_type = "STRING"
      },
      { op = "ENDLIST" },
    ]
    comment = "Notify on active DHCP leases"
    disable = false
  }
}

# TODO: The following prerequisite MUST exist on the grid before applying this example:
#   - syslog:endpoint named "example_syslog_endpoint" (infoblox_syslog_endpoint is not yet implemented)

// Create a Notification Rule for DNS RPZ events, sent to a syslog endpoint
resource "infoblox_notification_rule" "notification_rule_for_dns_rpz" {
  nios = {
    name                = "example_notification_rule_dns_rpz"
    event_type          = "DNS_RPZ"
    notification_action = "RESTAPI_TEMPLATE_INSTANCE"
    notification_target = "syslog:endpoint/<ref-of-example_syslog_endpoint>"
    template_instance = {
      template = "Version5_Syslog_Action_Template"
    }
    expression_list = [
      { op = "AND", op1_type = "LIST" },
      {
        op       = "EQ"
        op1      = "DNS_RPZ_TYPE"
        op1_type = "FIELD"
        op2      = "DNS_RPZ_TYPE_IP"
        op2_type = "STRING"
      },
      { op = "ENDLIST" },
    ]
  }
}

// Create a Notification Rule for IPAM events, published to a pxGrid endpoint
resource "infoblox_notification_rule" "notification_rule_for_ipam" {
  nios = {
    name                = "example_notification_rule_ipam"
    event_type          = "IPAM"
    notification_action = "RESTAPI_TEMPLATE_INSTANCE"
    notification_target = "pxgrid:endpoint/<ref-of-example_pxgrid_endpoint>"
    template_instance = {
      template = "IPAM_PxgridEvent"
    }
    publish_settings = {
      enabled_attributes = ["CLIENT_ID", "IPADDRESS"]
    }
    expression_list = [
      { op = "AND", op1_type = "LIST" },
      {
        op       = "EQ"
        op1      = "DHCP_LEASE_STATE"
        op1_type = "FIELD"
        op2      = "DHCP_LEASE_STATE_ACTIVE"
        op2_type = "STRING"
      },
      { op = "ENDLIST" },
    ]
  }
}
