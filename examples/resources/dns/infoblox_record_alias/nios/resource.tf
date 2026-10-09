// Create an Auth Zone (Required as Parent)
resource "infoblox_zone_auth" "example" {
  nios = {
    fqdn = "example.com"
  }
}

// Create Record Alias with Basic Fields
resource "infoblox_record_alias" "create_alias_record" {
  nios = {
    name        = "alias-record.${infoblox_zone_auth.example.nios.fqdn}"
    target_name = "server.${infoblox_zone_auth.example.nios.fqdn}"
    target_type = "A"
    view        = "default"
  }
}

// Create Record Alias with Additional Fields
resource "infoblox_record_alias" "create_alias_record_with_additional_fields" {
  nios = {
    name        = "alias-record-extra.${infoblox_zone_auth.example.nios.fqdn}"
    target_name = "server.${infoblox_zone_auth.example.nios.fqdn}"
    target_type = "A"
    view        = "default"

    // Optional fields
    comment = "Alias record with additional parameters"
    disable = false
    ext_attrs = {
      Site = "location-1"
    }
    ttl = 20
  }
}
