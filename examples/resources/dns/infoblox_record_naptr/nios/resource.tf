// Create Record NAPTR with Basic Fields
resource "infoblox_record_naptr" "create_record_basic" {
  nios = {
    name        = "naptr_record.example.com"
    order       = 10
    preference  = 10
    replacement = "."

    // Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create Record NAPTR with Additional Fields
resource "infoblox_record_naptr" "create_record_additional_fields" {
  nios = {
    // Basic Fields
    name        = "naptr_record1.example.com"
    order       = 10
    preference  = 10
    replacement = "."
    view        = "default"

    // Additional Fields
    flags              = "U"
    services           = "SIP+D2U"
    regexp             = "!^.*$!sip:jdoe@corpxyz.com!"
    ttl                = 3600
    creator            = "DYNAMIC"
    forbid_reclamation = false
    comment            = "NAPTR record created by Terraform"

    // Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
}
