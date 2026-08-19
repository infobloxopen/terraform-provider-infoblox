// Create an RPZ AAAA IP Address record with basic fields
resource "infoblox_record_rpz_aaaa_ipaddress" "create_rpz_aaaa_ipaddress_basic" {
  nios = {
    name     = "2001:db8::/64.rpz.example.com"
    ipv6addr = "2001:db8::1"
    rp_zone  = "rpz.example.com"
  }
}

// Create an RPZ AAAA IP Address record with additional fields
resource "infoblox_record_rpz_aaaa_ipaddress" "create_rpz_aaaa_ipaddress_additional" {
  nios = {
    name     = "2001:db8:1::/48.rpz.example.com"
    ipv6addr = "2001:db8::2"
    rp_zone  = "rpz.example.com"
    comment  = "Block traffic to 2001:db8:1::/48"
    disable  = false
    ttl      = 3600
    view     = "default"
    ext_attrs = {
      Site = "headquarters"
    }
  }
}
