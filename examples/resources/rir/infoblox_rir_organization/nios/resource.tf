// Manage a RIR Organization with Basic Fields
resource "infoblox_rir_organization" "basic" {
  nios = {
    id           = "ORG-CR17-IB"
    maintainer   = "infoblox"
    name         = "example_rir_organization"
    password     = "example-pass"
    sender_email = "support@infoblox.com"
    ext_attrs = {
      "RIPE Admin Contact"     = "ib-contact"
      "RIPE Country"           = "United Kingdom (GB)"
      "RIPE Technical Contact" = "EG123-IB"
      "RIPE Email"             = "support@infoblox.com"
    }
  }
}

// Manage a RIR Organization with Additional Fields
resource "infoblox_rir_organization" "with_additional_fields" {
  nios = {
    id           = "ORG-BB99-IB"
    maintainer   = "nios"
    name         = "example_rir_organization_additional"
    password     = "examplePass123"
    rir          = "RIPE"
    sender_email = "support@infoblox.com"
    ext_attrs = {
      "RIPE Admin Contact"     = "ib-contact"
      "RIPE Country"           = "United Kingdom (GB)"
      "RIPE Technical Contact" = "EG567-IB"
      "RIPE Email"             = "support@infoblox.com"
      "RIPE Remarks"           = "Example RIR Organization"
      "RIPE Organization Type" = "IANA"
      "RIPE Notify"            = "support@infoblox.com"
    }
  }
}
