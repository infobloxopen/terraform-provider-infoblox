// Create an RP Zone (Required as Parent)
resource "infoblox_zone_rp" "example" {
  nios = {
    fqdn = "rpz-clientip.example.com"
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

// Create an RPZ CNAME Client IP Address Record with Additional Fields
resource "infoblox_record_rpz_cname_clientipaddress" "create_record_rpz_cname_clientipaddress_additional" {
  nios = {
    name      = "12.0.0.2.${infoblox_zone_rp.example.nios.fqdn}"
    canonical = "rpz-passthru"
    rp_zone   = infoblox_zone_rp.example.nios.fqdn
    comment   = "Allow this client IP to pass through the RPZ"
    disable   = false
    ttl       = 3600
    ext_attrs = {
      Site = "datacenter-1"
    }
  }
}

// Create an RPZ CNAME Client IP Address Record in a Custom View
resource "infoblox_view" "parent_view" {
  nios = {
    name = "custom-view"
  }
}

resource "infoblox_zone_rp" "parent_zone" {
  nios = {
    fqdn = "rpz-custom.example.com"
    view = infoblox_view.parent_view.nios.name
  }
}

resource "infoblox_record_rpz_cname_clientipaddress" "create_record_rpz_cname_clientipaddress_custom_view" {
  nios = {
    name      = "12.0.0.3.${infoblox_zone_rp.parent_zone.nios.fqdn}"
    canonical = "*"
    rp_zone   = infoblox_zone_rp.parent_zone.nios.fqdn
    view      = infoblox_view.parent_view.nios.name
  }
}
