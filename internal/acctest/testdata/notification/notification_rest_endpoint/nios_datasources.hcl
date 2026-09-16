# Auto-generated datasource acceptance-test cases for NotificationRestEndpoint.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.log_level", "nios.name", "nios.outbound_member_type", "nios.server_cert_validation", "nios.sync_disabled", "nios.timeout", "nios.uri", "nios.username", "nios.vendor_identifier", "nios.wapi_user_name"]

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"

  filter {
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.comment", "nios.log_level", "nios.name", "nios.outbound_member_type", "nios.server_cert_validation", "nios.sync_disabled", "nios.timeout", "nios.uri", "nios.username", "nios.vendor_identifier", "nios.wapi_user_name"]

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
      ext_attrs            = { Site = "{{random}}" }
    }
  }

}
