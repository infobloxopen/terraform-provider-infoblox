// Create an Auth Zone (Required as Parent)
resource "infoblox_zone_auth" "parent_zone" {
  nios = {
    fqdn = "example.com"
  }
}

// Create Record NAPTR with Basic Fields
resource "infoblox_record_naptr" "create_record_basic" {
  nios = {
    name        = "naptr_record.${infoblox_zone_auth.parent_zone.nios.fqdn}"
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
    name        = "naptr_record1.${infoblox_zone_auth.parent_zone.nios.fqdn}"
    order       = 10
    preference  = 10
    replacement = "."

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
