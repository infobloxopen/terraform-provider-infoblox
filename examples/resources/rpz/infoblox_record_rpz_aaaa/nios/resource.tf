// Create an RP Zone (Required as Parent)
resource "infoblox_zone_rp" "example" {
  nios = {
    fqdn = "rpz.example.com"
  }
}

// Create Record RPZ AAAA - basic substitution rule
resource "infoblox_record_rpz_aaaa" "create_record_rpz_aaaa_basic" {
  nios = {
    name     = "blocked.${infoblox_zone_rp.example.nios.fqdn}"
    ipv6addr = "2002:1f93::1"
    rp_zone  = infoblox_zone_rp.example.nios.fqdn
  }
}

// Create Record RPZ AAAA - with additional fields
resource "infoblox_record_rpz_aaaa" "create_record_rpz_aaaa_full" {
  nios = {
    name     = "redirect.${infoblox_zone_rp.example.nios.fqdn}"
    ipv6addr = "2002:1f93::2"
    rp_zone  = infoblox_zone_rp.example.nios.fqdn
    ttl      = 60
    comment  = "Example RPZ AAAA record"
    ext_attrs = {
      Site = "location-1"
    }
  }
}
