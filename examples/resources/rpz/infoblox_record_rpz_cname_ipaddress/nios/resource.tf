// Create an RP Zone (Required as Parent)
resource "infoblox_zone_rp" "example" {
  nios = {
    fqdn = "rpzip.example.com"
  }
}

// Create Record RPZ CNAME IP Address with Basic Fields
resource "infoblox_record_rpz_cname_ipaddress" "basic" {
  nios = {
    name      = "11.0.0.1.${infoblox_zone_rp.example.nios.fqdn}"
    canonical = "11.0.0.1"
    rp_zone   = infoblox_zone_rp.example.nios.fqdn
  }
}

// Create Record RPZ CNAME IP Address with Additional Fields
resource "infoblox_record_rpz_cname_ipaddress" "additional" {
  nios = {
    name      = "11.0.0.2.${infoblox_zone_rp.example.nios.fqdn}"
    canonical = "11.0.0.2"
    rp_zone   = infoblox_zone_rp.example.nios.fqdn
    ttl       = 10
    comment   = "Example RPZ CNAME IP address record"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Block IP Address (No Data) Rule
resource "infoblox_record_rpz_cname_ipaddress" "block_no_data" {
  nios = {
    name      = "11.0.0.3.${infoblox_zone_rp.example.nios.fqdn}"
    canonical = "*"
    rp_zone   = infoblox_zone_rp.example.nios.fqdn
  }
}

// Create Record RPZ CNAME IP Address in a Custom View
resource "infoblox_view" "parent_view" {
  nios = {
    name = "custom-view"
  }
}

resource "infoblox_zone_rp" "parent_zone" {
  nios = {
    fqdn = "rpzip-custom.example.com"
    view = infoblox_view.parent_view.nios.name
  }
}

resource "infoblox_record_rpz_cname_ipaddress" "create_record_rpz_cname_ipaddress_custom_view" {
  nios = {
    name      = "11.0.0.4.${infoblox_zone_rp.parent_zone.nios.fqdn}"
    canonical = "11.0.0.4"
    rp_zone   = infoblox_zone_rp.parent_zone.nios.fqdn
    view      = infoblox_view.parent_view.nios.name
  }
}
