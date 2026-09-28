// Create an RP Zone (Required as Parent)
resource "infoblox_zone_rp" "example" {
  nios = {
    fqdn = "rpz.example.com"
  }
}

// Create Record RPZ A with Basic Fields
resource "infoblox_record_rpz_a" "create_record_rpz_a_basic" {
  nios = {
    name     = "a-record.${infoblox_zone_rp.example.nios.fqdn}"
    ipv4addr = "192.168.1.1"
    rp_zone  = infoblox_zone_rp.example.nios.fqdn

    // Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create Record RPZ A with Additional Fields
resource "infoblox_record_rpz_a" "create_record_rpz_a_additional" {
  nios = {
    // Basic Fields
    name     = "a-record-2.${infoblox_zone_rp.example.nios.fqdn}"
    ipv4addr = "192.168.1.2"
    rp_zone  = infoblox_zone_rp.example.nios.fqdn

    // Additional Fields
    ttl     = 3600
    disable = false
    comment = "RPZ A record created by Terraform"

    // Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create a Record RPZ A in a Custom View
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

resource "infoblox_record_rpz_a" "create_record_rpz_a_custom_view" {
  nios = {
    name     = "a-record.${infoblox_zone_rp.parent_zone.nios.fqdn}"
    ipv4addr = "192.168.2.1"
    rp_zone  = infoblox_zone_rp.parent_zone.nios.fqdn
    view     = infoblox_view.parent_view.nios.name
  }
}
