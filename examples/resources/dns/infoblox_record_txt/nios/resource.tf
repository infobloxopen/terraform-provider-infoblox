// Create an Auth Zone (Required as Parent)
resource "infoblox_zone_auth" "parent_zone" {
  nios = {
    fqdn = "example.com"
  }
}

// Create Record TXT with Basic Fields
resource "infoblox_record_txt" "create_record" {
  nios = {
    name = "example-txt-record.${infoblox_zone_auth.parent_zone.nios.fqdn}"
    text = "Example TXT Record"

    // Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create Record TXT with Additional Fields
resource "infoblox_record_txt" "create_with_additional_config" {
  nios = {
    name = "example-txt-record-with-config.${infoblox_zone_auth.parent_zone.nios.fqdn}"
    text = "Example TXT Record with Additional Config"

    // Additional Fields
    ttl     = 10
    creator = "DYNAMIC"
    comment = "Example TXT record"

    // Extensible Attributes
    ext_attrs = {
      Site = "location-2"
    }
  }
}
