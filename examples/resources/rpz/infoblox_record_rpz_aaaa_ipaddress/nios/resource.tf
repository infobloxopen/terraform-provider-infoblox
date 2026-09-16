// Create a custom DNS view (optional — omit to use the "default" view)
resource "infoblox_view" "example" {
  nios = {
    name = "example-view"
  }
}

// Create the parent Response Policy Zone in the custom view
resource "infoblox_zone_rp" "example" {
  nios = {
    fqdn = "rpz.example.com"
    view = infoblox_view.example.nios.name
  }
}

// Create an RPZ AAAA IP Address record with basic fields
resource "infoblox_record_rpz_aaaa_ipaddress" "create_rpz_aaaa_ipaddress_basic" {
  nios = {
    name     = "2001:db8::/64.${infoblox_zone_rp.example.nios.fqdn}"
    ipv6addr = "2001:db8::1"
    rp_zone  = infoblox_zone_rp.example.nios.fqdn
    view     = infoblox_view.example.nios.name
  }
}

// Create an RPZ AAAA IP Address record with additional fields
resource "infoblox_record_rpz_aaaa_ipaddress" "create_rpz_aaaa_ipaddress_additional" {
  nios = {
    name     = "2001:db8:1::/48.${infoblox_zone_rp.example.nios.fqdn}"
    ipv6addr = "2001:db8::2"
    rp_zone  = infoblox_zone_rp.example.nios.fqdn
    comment  = "Block traffic to 2001:db8:1::/48"
    disable  = false
    ttl      = 3600
    view     = infoblox_view.example.nios.name
    ext_attrs = {
      Site = "headquarters"
    }
  }
}
