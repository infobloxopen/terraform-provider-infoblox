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

// Create an RPZ A IP Address record with basic fields
resource "infoblox_record_rpz_a_ipaddress" "create_rpz_a_ipaddress_basic" {
  nios = {
    name     = "10.10.0.0/16.${infoblox_zone_rp.example.nios.fqdn}"
    ipv4addr = "10.10.0.1"
    rp_zone  = infoblox_zone_rp.example.nios.fqdn
    view     = infoblox_view.example.nios.name
  }
}

// Create an RPZ A IP Address record with additional fields
resource "infoblox_record_rpz_a_ipaddress" "create_rpz_a_ipaddress_additional" {
  nios = {
    name     = "192.168.1.0/24.${infoblox_zone_rp.example.nios.fqdn}"
    ipv4addr = "192.168.1.1"
    rp_zone  = infoblox_zone_rp.example.nios.fqdn
    comment  = "Block traffic to 192.168.1.0/24"
    disable  = false
    ttl      = 3600
    view     = infoblox_view.example.nios.name
    ext_attrs = {
      Site = "headquarters"
    }
  }
}
