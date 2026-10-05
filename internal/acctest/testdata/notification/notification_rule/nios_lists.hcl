# NotificationRule — nios list cases
case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_notification_rest_endpoint" "test_endpoint" {
    nios = {
      name = "{{random}}"
      outbound_member_type = "GM"
      uri = "https://www.example.com"
    }
  }
  PREREQ

  step {
    nios {
      event_type           = "DHCP_LEASES"
      expression_list      = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DHCP_LEASE_STATE", op1_type = "FIELD", op2 = "DHCP_LEASE_STATE_ACTIVE", op2_type = "STRING" }, { op = "ENDLIST" }]
      name                 = "{{random}}"
      notification_action  = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target  = "${infoblox_notification_rest_endpoint.test_endpoint.id}"
      template_instance    = { template = "DHCP_Lease" }
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_notification_rest_endpoint" "test_endpoint" {
    nios = {
      name = "{{random}}"
      outbound_member_type = "GM"
      uri = "https://www.example.com"
    }
  }
  PREREQ

  step {
    nios {
      event_type           = "DHCP_LEASES"
      expression_list      = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DHCP_LEASE_STATE", op1_type = "FIELD", op2 = "DHCP_LEASE_STATE_ACTIVE", op2_type = "STRING" }, { op = "ENDLIST" }]
      name                 = "{{random}}"
      notification_action  = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target  = "${infoblox_notification_rest_endpoint.test_endpoint.id}"
      template_instance    = { template = "DHCP_Lease" }
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        name = "nios.name"
      }
    }
  }

}
