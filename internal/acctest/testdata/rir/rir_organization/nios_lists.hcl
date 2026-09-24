# RirOrganization — nios list cases
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      id           = "ORG-CB{{random_int}}-IBTEST"
      maintainer   = "infoblox"
      name         = "{{random}}"
      password     = "test-pass"
      rir          = "RIPE"
      sender_email = "support@infoblox.com"
      ext_attrs = {
        "RIPE Admin Contact"     = "ib-contact"
        "RIPE Country"           = "United Kingdom (GB)"
        "RIPE Technical Contact" = "TEST123-IB"
        "RIPE Email"             = "support@infoblox.com"
      }
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend  = "nios"
  parallel = true

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  step {
    nios {
      id           = "ORG-CB{{random_int}}-IBTEST"
      maintainer   = "infoblox"
      name         = "{{random}}"
      password     = "test-pass"
      rir          = "RIPE"
      sender_email = "support@infoblox.com"
      ext_attrs = {
        "RIPE Admin Contact"     = "ib-contact"
        "RIPE Country"           = "United Kingdom (GB)"
        "RIPE Technical Contact" = "TEST123-IB"
        "RIPE Email"             = "support@infoblox.com"
      }
    }
  }

}

case "ext_attr_filters" {
  backend     = "nios"
  skip        = true
  skip_reason = "NIOS API does not allow non-RIPE extensible attributes on RIR Organization objects"
}
