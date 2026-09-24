// Create an RP Zone (Required as Parent)
resource "infoblox_zone_rp" "example" {
  nios = {
    fqdn = "rpz.example.com"
  }
}

// Create an RPZ CNAME IP Address DN Record with Basic Fields
resource "infoblox_record_rpz_cname_ipaddressdn" "create_record_rpz_cname_ipaddressdn_basic" {
  nios = {
    name      = "10.10.0.0/24.${infoblox_zone_rp.example.nios.fqdn}"
    canonical = "block.${infoblox_zone_rp.example.nios.fqdn}"
    rp_zone   = infoblox_zone_rp.example.nios.fqdn
  }
}

// Create an RPZ CNAME IP Address DN Record with Additional Fields
resource "infoblox_record_rpz_cname_ipaddressdn" "create_record_rpz_cname_ipaddressdn_additional" {
  nios = {
    name      = "192.168.1.0/24.${infoblox_zone_rp.example.nios.fqdn}"
    canonical = "redirect.${infoblox_zone_rp.example.nios.fqdn}"
    rp_zone   = infoblox_zone_rp.example.nios.fqdn
    comment   = "Block IP network from resolving"
    disable   = false
    ttl       = 3600
    ext_attrs = {
      Site = "datacenter-1"
    }
  }
}

// Create an RPZ CNAME IP Address DN Record in a Custom View
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

resource "infoblox_record_rpz_cname_ipaddressdn" "create_record_rpz_cname_ipaddressdn_custom_view" {
  nios = {
    name      = "10.10.0.0/24.${infoblox_zone_rp.parent_zone.nios.fqdn}"
    canonical = "block.${infoblox_zone_rp.parent_zone.nios.fqdn}"
    rp_zone   = infoblox_zone_rp.parent_zone.nios.fqdn
    view      = infoblox_view.parent_view.nios.name
  }
}
