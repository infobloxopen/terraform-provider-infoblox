// Create an Auth Zone (Required as Parent)
resource "infoblox_zone_auth" "parent_zone" {
  nios = {
    fqdn = "example_super_host_zone.com"
    view = "default"
  }
}

// Create a Record A (Required as Parent)
resource "infoblox_record_a" "parent_record_a" {
  nios = {
    name     = "parent_record.${infoblox_zone_auth.parent_zone.nios.fqdn}"
    ipv4addr = "10.20.1.2"
    view     = "default"
  }
}

// Create Super Host with Basic Fields
resource "infoblox_superhost" "create_super_host" {
  nios = {
    name = "example_super_host"
  }
}

// Create Super Host with Additional Fields
resource "infoblox_superhost" "create_super_host_with_additional_fields" {
  nios = {
    name = "example_super_host_with_associated_objects"

    // Additional Fields
    comment                   = "This is a Super Host example"
    delete_associated_objects = true
    disabled                  = false
    dns_associated_objects    = [infoblox_record_a.parent_record_a.id]

    // Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
}
