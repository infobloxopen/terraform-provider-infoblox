// Create an Auth Zone (Required as Parent)
resource "infoblox_zone_auth" "parent_zone" {
  nios = {
    fqdn = "example.com"
    view = "default"
  }
}

// Create a CAA Record with Basic Fields
resource "infoblox_record_caa" "record_caa_with_basic_fields" {
  nios = {
    name     = "caa-record.${infoblox_zone_auth.parent_zone.nios.fqdn}"
    ca_flag  = 1
    ca_tag   = "issue"
    ca_value = "digicert.com"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create a CAA Record with Additional Fields
resource "infoblox_record_caa" "record_caa_with_additional_fields" {
  nios = {
    name     = "caa-record1.${infoblox_zone_auth.parent_zone.nios.fqdn}"
    ca_flag  = 1
    ca_tag   = "issue"
    ca_value = "digicert.com"

    // Additional Fields
    comment            = "CAA Record created by Terraform"
    view               = "default"
    creator            = "STATIC"
    disable            = false
    ttl                = 10
    ddns_protected     = false
    forbid_reclamation = false
    ext_attrs = {
      Site = "location-1"
    }
  }
}
