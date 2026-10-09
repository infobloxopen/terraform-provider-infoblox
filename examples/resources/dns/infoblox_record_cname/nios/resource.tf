// Create an Auth Zone (Required as Parent)
resource "infoblox_zone_auth" "parent_zone" {
  nios = {
    fqdn = "example.com"
  }
}

// Create Record CNAME with Basic Fields
resource "infoblox_record_cname" "create_record_basic" {
  nios = {
    name      = "example_record.${infoblox_zone_auth.parent_zone.nios.fqdn}"
    canonical = "example-canonical-name.${infoblox_zone_auth.parent_zone.nios.fqdn}"

    // Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create Record CNAME with Additional Fields
resource "infoblox_record_cname" "create_record_additional_fields" {
  nios = {
    // Basic Fields
    name      = "example_record2.${infoblox_zone_auth.parent_zone.nios.fqdn}"
    canonical = "example-canonical-name2.${infoblox_zone_auth.parent_zone.nios.fqdn}"

    // Additional Fields
    ttl                = 3600
    creator            = "DYNAMIC"
    forbid_reclamation = false

    // Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
}
