# TODO: The following prerequisites MUST exist on the grid before running these tests:
#   - notification_rest_template : Version5_REST_API_Session_Template  (used by template_instance case)
#   - notification_rest_template : Version5_REST_API_Session_Template1 (used by template_instance case)
#   - grid member                : infoblox.172_28_83_167              (used by outbound_members/outbound_member_type cases)

case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
    }
    check = {
      "nios.name"                   = "{{random}}"
      "nios.outbound_member_type"   = "GM"
      "nios.uri"                    = "https://example.com"
      "nios.log_level"              = "WARNING"
      "nios.server_cert_validation" = "CA_CERT"
      "nios.sync_disabled"          = "false"
      "nios.timeout"                = "30"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
      comment              = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
      comment              = "This is a updated comment"
    }
    check = {
      "nios.comment" = "This is a updated comment"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
      ext_attrs            = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
      ext_attrs            = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "log_level" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
      log_level            = "DEBUG"
    }
    check = {
      "nios.log_level" = "DEBUG"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
      log_level            = "ERROR"
    }
    check = {
      "nios.log_level" = "ERROR"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name                 = "{{random2}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "client_certificate_file" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                    = "{{random}}"
      outbound_member_type    = "GM"
      uri                     = "https://example.com"
      client_certificate_file = "{{testdata_path}}/notification/notification_rest_endpoint/dummy-bundle.pem"
    }
  }

  step {
    nios {
      name                    = "{{random}}"
      outbound_member_type    = "GM"
      uri                     = "https://example.com"
      client_certificate_file = "{{testdata_path}}/notification/notification_rest_endpoint/dummy-bundle2.pem"
    }
  }

}

case "outbound_member_type" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
    }
    check = {
      "nios.outbound_member_type" = "GM"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example-updated.com"
    }
    check = {
      "nios.outbound_member_type" = "GM"
      "nios.uri"                  = "https://example-updated.com"
    }
  }

}

case "outbound_members" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
    }
    check = {
      "nios.outbound_member_type" = "GM"
      "nios.outbound_members.#"   = "0"
    }
  }

}

case "server_cert_validation" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                   = "{{random}}"
      outbound_member_type   = "GM"
      uri                    = "https://example.com"
      server_cert_validation = "CA_CERT_NO_HOSTNAME"
    }
    check = {
      "nios.server_cert_validation" = "CA_CERT_NO_HOSTNAME"
    }
  }

  step {
    nios {
      name                   = "{{random}}"
      outbound_member_type   = "GM"
      uri                    = "https://example.com"
      server_cert_validation = "NO_VALIDATION"
    }
    check = {
      "nios.server_cert_validation" = "NO_VALIDATION"
    }
  }

}

case "sync_disabled" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
      sync_disabled        = false
    }
    check = {
      "nios.sync_disabled" = "false"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
      sync_disabled        = true
    }
    check = {
      "nios.sync_disabled" = "true"
    }
  }

}

case "template_instance" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
      template_instance    = { template = "Version5_REST_API_Session_Template" }
    }
    check = {
      "nios.template_instance.template" = "Version5_REST_API_Session_Template"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
      template_instance    = { template = "Version5_REST_API_Session_Template1" }
    }
    check = {
      "nios.template_instance.template" = "Version5_REST_API_Session_Template1"
    }
  }

}

case "timeout" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
      timeout              = 100
    }
    check = {
      "nios.timeout" = "100"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
      timeout              = 200
    }
    check = {
      "nios.timeout" = "200"
    }
  }

}

case "uri" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
    }
    check = {
      "nios.uri" = "https://example.com"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example-updated.com"
    }
    check = {
      "nios.uri" = "https://example-updated.com"
    }
  }

}

case "username" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
      username             = "example_username"
      password             = "example_password"
    }
    check = {
      "nios.username" = "example_username"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
      username             = "example_username_updated"
      password             = "example_password_updated"
    }
    check = {
      "nios.username" = "example_username_updated"
    }
  }

}

case "vendor_identifier" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
      vendor_identifier    = "WAPI"
    }
    check = {
      "nios.vendor_identifier" = "WAPI"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
      vendor_identifier    = "extattrsgg WAPI"
    }
    check = {
      "nios.vendor_identifier" = "extattrsgg WAPI"
    }
  }

}

case "wapi_user_name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
      wapi_user_name       = "example_wapi_username"
      wapi_user_password   = "example_wapi_password"
    }
    check = {
      "nios.wapi_user_name" = "example_wapi_username"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
      wapi_user_name       = "example_wapi_username_updated"
      wapi_user_password   = "example_wapi_password_updated"
    }
    check = {
      "nios.wapi_user_name" = "example_wapi_username_updated"
    }
  }

}
