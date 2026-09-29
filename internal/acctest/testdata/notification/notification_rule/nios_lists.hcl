# NotificationRule — nios list cases
# TODO: The following prerequisites MUST exist on the grid before running these tests:
#   - syslog:endpoint : syslog:endpoint/b25lLmVuZHBvaW50JDE:syslogendpoint123 (name: syslogendpoint123)
#     (infoblox_syslog_endpoint is not yet implemented)

case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      event_type           = "DNS_RPZ"
      expression_list      = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DNS_RPZ_TYPE", op1_type = "FIELD", op2 = "DNS_RPZ_TYPE_IP", op2_type = "STRING" }, { op = "ENDLIST" }]
      name                 = "{{random}}"
      notification_action  = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target  = "syslog:endpoint/b25lLmVuZHBvaW50JDE:syslogendpoint123"
      template_instance    = { template = "Version5_Syslog_Action_Template" }
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

  step {
    nios {
      event_type           = "DNS_RPZ"
      expression_list      = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DNS_RPZ_TYPE", op1_type = "FIELD", op2 = "DNS_RPZ_TYPE_IP", op2_type = "STRING" }, { op = "ENDLIST" }]
      name                 = "{{random}}"
      notification_action  = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target  = "syslog:endpoint/b25lLmVuZHBvaW50JDE:syslogendpoint123"
      template_instance    = { template = "Version5_Syslog_Action_Template" }
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
