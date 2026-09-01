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
