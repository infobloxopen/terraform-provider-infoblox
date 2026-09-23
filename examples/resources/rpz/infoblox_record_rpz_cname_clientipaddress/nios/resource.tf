// Parents
resource "infoblox_zone_rp" "example" {
  nios = {
    fqdn = "rpz-clientip.example.com"
  }
}

resource "infoblox_view" "example" {
  nios = {
    name = "custom-view"
  }
}

resource "infoblox_zone_rp" "example_custom_view" {
  nios = {
    fqdn = "rpz-custom.example.com"
    view = infoblox_view.example.nios.name
  }
}

// Create an RPZ CNAME Client IP Address Record — blocks the IP (empty canonical)
resource "infoblox_record_rpz_cname_clientipaddress" "create_record_rpz_cname_clientipaddress_basic" {
  nios = {
    name      = "12.0.0.1.${infoblox_zone_rp.example.nios.fqdn}"
    canonical = ""
    rp_zone   = infoblox_zone_rp.example.nios.fqdn
  }
}

// Create an RPZ CNAME Client IP Address Record with Additional Fields in a Custom View
resource "infoblox_record_rpz_cname_clientipaddress" "create_record_rpz_cname_clientipaddress_additional" {
  nios = {
    name      = "12.0.0.2.${infoblox_zone_rp.example_custom_view.nios.fqdn}"
    canonical = "rpz-passthru"
    rp_zone   = infoblox_zone_rp.example_custom_view.nios.fqdn
    view      = infoblox_view.example.nios.name
    comment   = "Allow this client IP to pass through the RPZ"
    disable   = false
    ttl       = 3600
    ext_attrs = {
      Site = "datacenter-1"
    }
  }
}
