# Auto-generated datasource acceptance-test cases for RirOrganization.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.id", "nios.maintainer", "nios.name", "nios.rir", "nios.sender_email"]

  step {
    nios {
      id           = "ORG-CB{{random_int}}-IBTEST"
      maintainer   = "infoblox"
      password     = tostring("test-pass")
      name         = "{{random}}"
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
