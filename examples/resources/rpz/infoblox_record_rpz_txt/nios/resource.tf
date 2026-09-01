// Create an RPZ Zone (Required as Parent)
resource "infoblox_zone_rp" "example" {
  nios = {
    fqdn = "rpz.example.com"
  }
}

// Create Record RPZ TXT with Basic Fields
resource "infoblox_record_rpz_txt" "create_record_rpz_txt_basic" {
  nios = {
    name    = "blocked.${infoblox_zone_rp.example.nios.fqdn}"
    text    = "Example text"
    rp_zone = infoblox_zone_rp.example.nios.fqdn
  }
}

// Create Record RPZ TXT with Additional Fields
resource "infoblox_record_rpz_txt" "create_record_rpz_txt_additional" {
  nios = {
    name    = "blocked-with-ttl.${infoblox_zone_rp.example.nios.fqdn}"
    text    = "Example text with Additional Config"
    rp_zone = infoblox_zone_rp.example.nios.fqdn
    ttl     = 10
    comment = "Example RPZ TXT record"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create Record RPZ TXT in a Custom View
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

resource "infoblox_record_rpz_txt" "create_record_rpz_txt_custom_view" {
  nios = {
    name    = "blocked.${infoblox_zone_rp.custom_view.nios.fqdn}"
    text    = "Example text in custom view"
    rp_zone = infoblox_zone_rp.custom_view.nios.fqdn
    view    = infoblox_view.custom.nios.name
  }
}
