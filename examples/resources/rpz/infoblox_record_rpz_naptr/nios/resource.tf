// Note: RPZ zones must be pre-created in NIOS (infoblox_zone_rp is not managed by this provider).
// Create a Substitute (NAPTR Record) Rule with Basic Fields
resource "infoblox_record_rpz_naptr" "create_record_basic" {
  nios = {
    name        = "naptr.rpz.example.com"
    rp_zone     = "rpz.example.com"
    order       = 10
    preference  = 10
    replacement = "."

    // Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create a Substitute (NAPTR Record) Rule with Additional Fields
resource "infoblox_record_rpz_naptr" "create_record_additional_fields" {
  nios = {
    // Basic Fields
    name        = "naptr1.rpz.example.com"
    rp_zone     = "rpz.example.com"
    order       = 10
    preference  = 10
    replacement = "."

    // Additional Fields
    flags    = "U"
    services = "SIP+D2U"
    regexp   = "!^.*$!sip:jdoe@corpxyz.com!"
    ttl      = 3600
    disable  = false
    comment  = "NAPTR RPZ record created by Terraform"

    // Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create DNS View (Required as Parent)
// Note: RPZ zone "custom-rpz.example.com" must be pre-created in NIOS in this view.
resource "infoblox_view" "custom_view" {
  nios = {
    name = "custom-view"
  }
}

// Create a Substitute (NAPTR Record) Rule in a Custom View
resource "infoblox_record_rpz_naptr" "create_record_custom_view" {
  nios = {
    name        = "naptr.custom-rpz.example.com"
    rp_zone     = "custom-rpz.example.com"
    order       = 20
    preference  = 20
    replacement = "."
    view        = infoblox_view.custom_view.nios.name
  }
}
