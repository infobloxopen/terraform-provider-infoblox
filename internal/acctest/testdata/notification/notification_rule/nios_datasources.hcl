# Auto-generated datasource acceptance-test cases for NotificationRule.
case "filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_notification_rest_endpoint" "test_endpoint" {
    nios = {
      name = "{{random}}"
      outbound_member_type = "GM"
      uri = "https://www.example.com"
    }
  }
  PREREQ

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.all_members", "nios.comment", "nios.disable", "nios.enable_event_deduplication", "nios.enable_event_deduplication_log", "nios.event_deduplication_lookback_period", "nios.event_priority", "nios.event_type", "nios.name", "nios.notification_action", "nios.notification_target"]

  step {
    nios {
      event_type           = "DHCP_LEASES"
      expression_list      = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DHCP_LEASE_STATE", op1_type = "FIELD", op2 = "DHCP_LEASE_STATE_ACTIVE", op2_type = "STRING" }, { op = "ENDLIST" }]
      name                 = "{{random}}"
      notification_action  = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target  = "$${infoblox_notification_rest_endpoint.test_endpoint.id}"
      template_instance    = { template = "DHCP_Lease" }
    }
  }

}
