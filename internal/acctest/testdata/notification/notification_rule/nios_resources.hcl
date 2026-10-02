# DNS_RPZ/event-deduplication cases need a syslog:endpoint (NIOS_SYSLOG_ENDPOINT_REF); the
# publish_settings case needs a pxgrid:endpoint (NIOS_PXGRID_ENDPOINT_REF). Neither
# infoblox_syslog_endpoint nor infoblox_pxgrid_endpoint is implemented yet, so these cases
# reference the grid's pre-provisioned objects via env var (populated by the shared
# integration-test setup in CI; see internal/acctest/integration_tests/setup/nios) and
# skip cleanly via skip_if_env_empty when run locally without those env vars set.

case "basic" {
  backend  = "nios"
  parallel = true
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
      notification_target  = "$${infoblox_notification_rest_endpoint.test_endpoint.id}"
      template_instance    = { template = "DHCP_Lease" }
    }
    check = {
      "nios.event_type"                     = "DHCP_LEASES"
      "nios.name"                           = "{{random}}"
      "nios.notification_action"            = "RESTAPI_TEMPLATE_INSTANCE"
      "nios.template_instance.template"     = "DHCP_Lease"
      "nios.expression_list.#"              = "3"
      "nios.expression_list.0.op"           = "AND"
      "nios.expression_list.0.op1_type"     = "LIST"
      "nios.expression_list.1.op"           = "EQ"
      "nios.expression_list.1.op1"          = "DHCP_LEASE_STATE"
      "nios.expression_list.1.op1_type"     = "FIELD"
      "nios.expression_list.1.op2"          = "DHCP_LEASE_STATE_ACTIVE"
      "nios.expression_list.1.op2_type"     = "STRING"
      "nios.expression_list.2.op"           = "ENDLIST"
      "nios.disable"                        = "false"
      "nios.enable_event_deduplication"     = "false"
      "nios.enable_event_deduplication_log" = "false"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
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
      notification_target  = "$${infoblox_notification_rest_endpoint.test_endpoint.id}"
      template_instance    = { template = "DHCP_Lease" }
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true
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
      notification_target  = "$${infoblox_notification_rest_endpoint.test_endpoint.id}"
      template_instance    = { template = "DHCP_Lease" }
      comment              = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      event_type           = "DHCP_LEASES"
      expression_list      = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DHCP_LEASE_STATE", op1_type = "FIELD", op2 = "DHCP_LEASE_STATE_ACTIVE", op2_type = "STRING" }, { op = "ENDLIST" }]
      name                 = "{{random}}"
      notification_action  = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target  = "$${infoblox_notification_rest_endpoint.test_endpoint.id}"
      template_instance    = { template = "DHCP_Lease" }
      comment              = "This is an updated comment"
    }
    check = {
      "nios.comment" = "This is an updated comment"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true
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
      notification_target  = "$${infoblox_notification_rest_endpoint.test_endpoint.id}"
      template_instance    = { template = "DHCP_Lease" }
      disable              = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      event_type           = "DHCP_LEASES"
      expression_list      = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DHCP_LEASE_STATE", op1_type = "FIELD", op2 = "DHCP_LEASE_STATE_ACTIVE", op2_type = "STRING" }, { op = "ENDLIST" }]
      name                 = "{{random}}"
      notification_action  = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target  = "$${infoblox_notification_rest_endpoint.test_endpoint.id}"
      template_instance    = { template = "DHCP_Lease" }
      disable              = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

}

case "enable_event_deduplication" {
  backend           = "nios"
  parallel          = true
  skip_if_env_empty = ["NIOS_SYSLOG_ENDPOINT_REF"]
  skip_reason       = "NIOS_SYSLOG_ENDPOINT_REF environment variable must be set for this test to run"

  step {
    nios {
      event_type                 = "DNS_RPZ"
      expression_list            = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DNS_RPZ_TYPE", op1_type = "FIELD", op2 = "DNS_RPZ_TYPE_IP", op2_type = "STRING" }, { op = "ENDLIST" }]
      name                       = "{{random}}"
      notification_action        = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target        = "{{nios_syslog_endpoint_ref}}"
      template_instance          = { template = "Version5_Syslog_Action_Template" }
      enable_event_deduplication = false
      event_deduplication_fields = ["SOURCE_IP", "QUERY_NAME"]
    }
    check = {
      "nios.enable_event_deduplication" = "false"
    }
  }

  step {
    nios {
      event_type                 = "DNS_RPZ"
      expression_list            = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DNS_RPZ_TYPE", op1_type = "FIELD", op2 = "DNS_RPZ_TYPE_IP", op2_type = "STRING" }, { op = "ENDLIST" }]
      name                       = "{{random}}"
      notification_action        = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target        = "{{nios_syslog_endpoint_ref}}"
      template_instance          = { template = "Version5_Syslog_Action_Template" }
      enable_event_deduplication = true
      event_deduplication_fields = ["SOURCE_IP", "QUERY_NAME"]
    }
    check = {
      "nios.enable_event_deduplication" = "true"
    }
  }

}

case "enable_event_deduplication_log" {
  backend           = "nios"
  parallel          = true
  skip_if_env_empty = ["NIOS_SYSLOG_ENDPOINT_REF"]
  skip_reason       = "NIOS_SYSLOG_ENDPOINT_REF environment variable must be set for this test to run"

  step {
    nios {
      event_type                     = "DNS_RPZ"
      expression_list                = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DNS_RPZ_TYPE", op1_type = "FIELD", op2 = "DNS_RPZ_TYPE_IP", op2_type = "STRING" }, { op = "ENDLIST" }]
      name                           = "{{random}}"
      notification_action            = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target            = "{{nios_syslog_endpoint_ref}}"
      template_instance              = { template = "Version5_Syslog_Action_Template" }
      enable_event_deduplication_log = false
      event_deduplication_fields     = ["SOURCE_IP", "QUERY_NAME"]
    }
    check = {
      "nios.enable_event_deduplication_log" = "false"
    }
  }

  step {
    nios {
      event_type                     = "DNS_RPZ"
      expression_list                = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DNS_RPZ_TYPE", op1_type = "FIELD", op2 = "DNS_RPZ_TYPE_IP", op2_type = "STRING" }, { op = "ENDLIST" }]
      name                           = "{{random}}"
      notification_action            = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target            = "{{nios_syslog_endpoint_ref}}"
      template_instance              = { template = "Version5_Syslog_Action_Template" }
      enable_event_deduplication_log = true
      event_deduplication_fields     = ["SOURCE_IP", "QUERY_NAME"]
    }
    check = {
      "nios.enable_event_deduplication_log" = "true"
    }
  }

}

case "event_deduplication_fields" {
  backend           = "nios"
  parallel          = true
  skip_if_env_empty = ["NIOS_SYSLOG_ENDPOINT_REF"]
  skip_reason       = "NIOS_SYSLOG_ENDPOINT_REF environment variable must be set for this test to run"

  step {
    nios {
      event_type                 = "DNS_RPZ"
      expression_list            = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DNS_RPZ_TYPE", op1_type = "FIELD", op2 = "DNS_RPZ_TYPE_IP", op2_type = "STRING" }, { op = "ENDLIST" }]
      name                       = "{{random}}"
      notification_action        = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target        = "{{nios_syslog_endpoint_ref}}"
      template_instance          = { template = "Version5_Syslog_Action_Template" }
      event_deduplication_fields = ["SOURCE_IP"]
    }
    check = {
      "nios.event_deduplication_fields.#" = "1"
      "nios.event_deduplication_fields.0" = "SOURCE_IP"
    }
  }

  step {
    nios {
      event_type                 = "DNS_RPZ"
      expression_list            = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DNS_RPZ_TYPE", op1_type = "FIELD", op2 = "DNS_RPZ_TYPE_IP", op2_type = "STRING" }, { op = "ENDLIST" }]
      name                       = "{{random}}"
      notification_action        = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target        = "{{nios_syslog_endpoint_ref}}"
      template_instance          = { template = "Version5_Syslog_Action_Template" }
      event_deduplication_fields = ["SOURCE_IP", "QUERY_NAME"]
    }
    check = {
      "nios.event_deduplication_fields.#" = "2"
      "nios.event_deduplication_fields.0" = "SOURCE_IP"
      "nios.event_deduplication_fields.1" = "QUERY_NAME"
    }
  }

}

case "event_deduplication_lookback_period" {
  backend           = "nios"
  parallel          = true
  skip_if_env_empty = ["NIOS_SYSLOG_ENDPOINT_REF"]
  skip_reason       = "NIOS_SYSLOG_ENDPOINT_REF environment variable must be set for this test to run"

  step {
    nios {
      event_type                          = "DNS_RPZ"
      expression_list                     = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DNS_RPZ_TYPE", op1_type = "FIELD", op2 = "DNS_RPZ_TYPE_IP", op2_type = "STRING" }, { op = "ENDLIST" }]
      name                                = "{{random}}"
      notification_action                 = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target                 = "{{nios_syslog_endpoint_ref}}"
      template_instance                   = { template = "Version5_Syslog_Action_Template" }
      event_deduplication_lookback_period = 500
      event_deduplication_fields          = ["SOURCE_IP", "QUERY_NAME"]
    }
    check = {
      "nios.event_deduplication_lookback_period" = "500"
    }
  }

  step {
    nios {
      event_type                          = "DNS_RPZ"
      expression_list                     = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DNS_RPZ_TYPE", op1_type = "FIELD", op2 = "DNS_RPZ_TYPE_IP", op2_type = "STRING" }, { op = "ENDLIST" }]
      name                                = "{{random}}"
      notification_action                 = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target                 = "{{nios_syslog_endpoint_ref}}"
      template_instance                   = { template = "Version5_Syslog_Action_Template" }
      event_deduplication_lookback_period = 600
      event_deduplication_fields          = ["SOURCE_IP", "QUERY_NAME"]
    }
    check = {
      "nios.event_deduplication_lookback_period" = "600"
    }
  }

}

case "event_priority" {
  backend     = "nios"
  skip        = true
  skip_reason = "event_priority can only be configured for notification rules with the SCHEDULE event type, which itself requires additional grid configuration not available in this environment"
  parallel    = true

  step {
    nios {
      event_type           = "DHCP_LEASES"
      expression_list      = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DHCP_LEASE_STATE", op1_type = "FIELD", op2 = "DHCP_LEASE_STATE_ACTIVE", op2_type = "STRING" }, { op = "ENDLIST" }]
      name                 = "{{random}}"
      notification_action  = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target  = "{{nios_syslog_endpoint_ref}}"
      template_instance    = { template = "DHCP_Lease" }
      event_priority       = "NORMAL"
    }
    check = {
      "nios.event_priority" = "NORMAL"
    }
  }

  step {
    nios {
      event_type           = "DHCP_LEASES"
      expression_list      = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DHCP_LEASE_STATE", op1_type = "FIELD", op2 = "DHCP_LEASE_STATE_ACTIVE", op2_type = "STRING" }, { op = "ENDLIST" }]
      name                 = "{{random}}"
      notification_action  = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target  = "{{nios_syslog_endpoint_ref}}"
      template_instance    = { template = "DHCP_Lease" }
      event_priority       = "HIGH"
    }
    check = {
      "nios.event_priority" = "HIGH"
    }
  }

}

case "event_type" {
  backend           = "nios"
  parallel          = true
  skip_if_env_empty = ["NIOS_SYSLOG_ENDPOINT_REF"]
  skip_reason       = "NIOS_SYSLOG_ENDPOINT_REF environment variable must be set for this test to run"
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
      notification_target  = "$${infoblox_notification_rest_endpoint.test_endpoint.id}"
      template_instance    = { template = "DHCP_Lease" }
    }
    check = {
      "nios.event_type" = "DHCP_LEASES"
    }
  }

  step {
    nios {
      event_type           = "DNS_RPZ"
      expression_list      = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DNS_RPZ_TYPE", op1_type = "FIELD", op2 = "DNS_RPZ_TYPE_IP", op2_type = "STRING" }, { op = "ENDLIST" }]
      name                 = "{{random}}"
      notification_action  = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target  = "{{nios_syslog_endpoint_ref}}"
      template_instance    = { template = "Version5_Syslog_Action_Template" }
    }
    check = {
      "nios.event_type" = "DNS_RPZ"
    }
  }

}

case "expression_list" {
  backend           = "nios"
  parallel          = true
  skip_if_env_empty = ["NIOS_SYSLOG_ENDPOINT_REF"]
  skip_reason       = "NIOS_SYSLOG_ENDPOINT_REF environment variable must be set for this test to run"
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
      notification_target  = "$${infoblox_notification_rest_endpoint.test_endpoint.id}"
      template_instance    = { template = "DHCP_Lease" }
    }
    check = {
      "nios.expression_list.#"          = "3"
      "nios.expression_list.0.op"       = "AND"
      "nios.expression_list.0.op1_type" = "LIST"
      "nios.expression_list.1.op"       = "EQ"
      "nios.expression_list.1.op1"      = "DHCP_LEASE_STATE"
      "nios.expression_list.1.op1_type" = "FIELD"
      "nios.expression_list.1.op2"      = "DHCP_LEASE_STATE_ACTIVE"
      "nios.expression_list.1.op2_type" = "STRING"
      "nios.expression_list.2.op"       = "ENDLIST"
    }
  }

  step {
    nios {
      event_type           = "DNS_RPZ"
      expression_list      = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DNS_RPZ_TYPE", op1_type = "FIELD", op2 = "DNS_RPZ_TYPE_IP", op2_type = "STRING" }, { op = "ENDLIST" }]
      name                 = "{{random}}"
      notification_action  = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target  = "{{nios_syslog_endpoint_ref}}"
      template_instance    = { template = "Version5_Syslog_Action_Template" }
    }
    check = {
      "nios.expression_list.#"          = "3"
      "nios.expression_list.0.op"       = "AND"
      "nios.expression_list.0.op1_type" = "LIST"
      "nios.expression_list.1.op"       = "EQ"
      "nios.expression_list.1.op1"      = "DNS_RPZ_TYPE"
      "nios.expression_list.1.op1_type" = "FIELD"
      "nios.expression_list.1.op2"      = "DNS_RPZ_TYPE_IP"
      "nios.expression_list.1.op2_type" = "STRING"
      "nios.expression_list.2.op"       = "ENDLIST"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true
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
      notification_target  = "$${infoblox_notification_rest_endpoint.test_endpoint.id}"
      template_instance    = { template = "DHCP_Lease" }
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

}

case "notification_action" {
  backend     = "nios"
  skip        = true
  skip_reason = "The only exercisable notification_action value without additional infra is RESTAPI_TEMPLATE_INSTANCE. Confirmed live against the grid's pxgrid:endpoint (with both publish_settings and subscribe_settings already configured): every event_type/pxgrid:endpoint combination for CISCOISE_PUBLISH/CISCOISE_QUARANTINE is rejected with \"The notification action specified for the endpoint type TYPE_PXGRID is incorrect\", distinct from the endpoint error that DOES work for the publish_settings field case. This points to a required real Cisco ISE session-directory integration not reproducible here"
  parallel    = true
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
      notification_target  = "$${infoblox_notification_rest_endpoint.test_endpoint.id}"
      template_instance    = { template = "DHCP_Lease" }
    }
    check = {
      "nios.notification_action" = "RESTAPI_TEMPLATE_INSTANCE"
    }
  }

}

case "notification_target" {
  backend           = "nios"
  parallel          = true
  skip_if_env_empty = ["NIOS_SYSLOG_ENDPOINT_REF"]
  skip_reason       = "NIOS_SYSLOG_ENDPOINT_REF environment variable must be set for this test to run"
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
      notification_target  = "$${infoblox_notification_rest_endpoint.test_endpoint.id}"
      template_instance    = { template = "DHCP_Lease" }
    }
    check_pair = {
      "nios.notification_target" = infoblox_notification_rest_endpoint.test_endpoint.id
    }
  }

  step {
    nios {
      event_type           = "DNS_RPZ"
      expression_list      = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DNS_RPZ_TYPE", op1_type = "FIELD", op2 = "DNS_RPZ_TYPE_IP", op2_type = "STRING" }, { op = "ENDLIST" }]
      name                 = "{{random}}"
      notification_action  = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target  = "{{nios_syslog_endpoint_ref}}"
      template_instance    = { template = "Version5_Syslog_Action_Template" }
    }
    check = {
      "nios.notification_target" = "{{nios_syslog_endpoint_ref}}"
    }
  }

}

case "publish_settings" {
  backend           = "nios"
  parallel          = true
  skip_if_env_empty = ["NIOS_PXGRID_ENDPOINT_REF"]
  skip_reason       = "NIOS_PXGRID_ENDPOINT_REF environment variable must be set for this test to run"

  step {
    nios {
      event_type          = "DHCP_LEASES"
      expression_list     = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DHCP_LEASE_STATE", op1_type = "FIELD", op2 = "DHCP_LEASE_STATE_ACTIVE", op2_type = "STRING" }, { op = "ENDLIST" }]
      name                = "{{random}}"
      notification_action = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target = "{{pxgrid_endpoint_ref}}"
      template_instance    = { template = "IPAM_PxgridEvent" }
      publish_settings    = { enabled_attributes = ["CLIENT_ID", "IPADDRESS"] }
    }
    check = {
      "nios.publish_settings.enabled_attributes.#" = "2"
      "nios.publish_settings.enabled_attributes.0" = "CLIENT_ID"
      "nios.publish_settings.enabled_attributes.1" = "IPADDRESS"
    }
  }

  step {
    nios {
      event_type          = "DHCP_LEASES"
      expression_list     = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DHCP_LEASE_STATE", op1_type = "FIELD", op2 = "DHCP_LEASE_STATE_ACTIVE", op2_type = "STRING" }, { op = "ENDLIST" }]
      name                = "{{random}}"
      notification_action = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target = "{{pxgrid_endpoint_ref}}"
      template_instance    = { template = "IPAM_PxgridEvent" }
      publish_settings    = { enabled_attributes = ["CLIENT_ID", "IPADDRESS", "LEASE_STATE"] }
    }
    check = {
      "nios.publish_settings.enabled_attributes.#" = "3"
      "nios.publish_settings.enabled_attributes.0" = "CLIENT_ID"
      "nios.publish_settings.enabled_attributes.1" = "IPADDRESS"
      "nios.publish_settings.enabled_attributes.2" = "LEASE_STATE"
    }
  }

}

case "scheduled_event" {
  backend     = "nios"
  skip        = true
  skip_reason = "scheduled_event requires event_type SCHEDULE, which is not exercised by any other case and needs additional date/time config verified against a real grid clock; left unimplemented, matching the legacy nios provider"
  parallel    = true

  step {
    nios {
      event_type        = "SCHEDULE"
      expression_list   = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DHCP_LEASE_STATE", op1_type = "FIELD", op2 = "DHCP_LEASE_STATE_ACTIVE", op2_type = "STRING" }, { op = "ENDLIST" }]
      name              = "{{random}}"
      notification_action  = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target  = "{{nios_syslog_endpoint_ref}}"
      template_instance = { template = "DHCP_Lease" }
      scheduled_event   = { weekdays = ["TUESDAY", "WEDNESDAY", "MONDAY"], frequency = "WEEKLY", every = 15, minutes_past_hour = 6, disable = false, repeat = "RECUR", hour_of_day = 20 }
    }
    check = {
      "nios.scheduled_event.frequency" = "WEEKLY"
    }
  }

  step {
    nios {
      event_type        = "SCHEDULE"
      expression_list   = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DHCP_LEASE_STATE", op1_type = "FIELD", op2 = "DHCP_LEASE_STATE_ACTIVE", op2_type = "STRING" }, { op = "ENDLIST" }]
      name              = "{{random}}"
      notification_action  = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target  = "{{nios_syslog_endpoint_ref}}"
      template_instance = { template = "DHCP_Lease" }
      scheduled_event   = { minutes_past_hour = 6, repeat = "ONCE", day_of_month = 30, month = 1, year = 2026, hour_of_day = 20 }
    }
    check = {
      "nios.scheduled_event.repeat" = "ONCE"
    }
  }

}

case "template_instance" {
  backend           = "nios"
  parallel          = true
  skip_if_env_empty = ["NIOS_SYSLOG_ENDPOINT_REF"]
  skip_reason       = "NIOS_SYSLOG_ENDPOINT_REF environment variable must be set for this test to run"
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
      notification_target  = "$${infoblox_notification_rest_endpoint.test_endpoint.id}"
      template_instance    = { template = "DHCP_Lease" }
    }
    check = {
      "nios.template_instance.template" = "DHCP_Lease"
    }
  }

  step {
    nios {
      event_type           = "DNS_RPZ"
      expression_list      = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "DNS_RPZ_TYPE", op1_type = "FIELD", op2 = "DNS_RPZ_TYPE_IP", op2_type = "STRING" }, { op = "ENDLIST" }]
      name                 = "{{random}}"
      notification_action  = "RESTAPI_TEMPLATE_INSTANCE"
      notification_target  = "{{nios_syslog_endpoint_ref}}"
      template_instance    = { template = "Version5_Syslog_Action_Template" }
    }
    check = {
      "nios.template_instance.template" = "Version5_Syslog_Action_Template"
    }
  }

}
