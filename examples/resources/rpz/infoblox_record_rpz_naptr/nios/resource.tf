// Create an RPZ Zone (Required as Parent)
resource "infoblox_zone_rp" "example" {
  nios = {
    fqdn = "rpz.example.com"
  }
}

// Create Record RPZ NAPTR with Basic Fields
resource "infoblox_record_rpz_naptr" "create_record_rpz_naptr_basic" {
  nios = {
    name        = "naptr-record.${infoblox_zone_rp.example.nios.fqdn}"
    rp_zone     = infoblox_zone_rp.example.nios.fqdn
    order       = 10
    preference  = 10
    replacement = "."
  }
}

// Create Record RPZ NAPTR with Additional Fields
resource "infoblox_record_rpz_naptr" "create_record_rpz_naptr_additional" {
  nios = {
    name        = "naptr-record-2.${infoblox_zone_rp.example.nios.fqdn}"
    rp_zone     = infoblox_zone_rp.example.nios.fqdn
    order       = 20
    preference  = 20
    replacement = "."
    flags       = "U"
    services    = "SIP+D2U"
    regexp      = "!^.*$!sip:jdoe@corpxyz.com!"
    ttl         = 3600
    comment     = "NAPTR RPZ record created by Terraform"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create a Substitute (NAPTR Record) Rule in a Custom View
resource "infoblox_view" "custom" {
  nios = {
    name = "custom-view"
  }
}

resource "infoblox_zone_rp" "custom_view" {
  nios = {
    fqdn = "rpz-custom.example.com"
    view = infoblox_view.custom.nios.name
  }
}

resource "infoblox_record_rpz_naptr" "create_record_rpz_naptr_custom_view" {
  nios = {
    name        = "naptr-record.${infoblox_zone_rp.custom_view.nios.fqdn}"
    rp_zone     = infoblox_zone_rp.custom_view.nios.fqdn
    order       = 10
    preference  = 10
    replacement = "."
    view        = infoblox_view.custom.nios.name
  }
}
