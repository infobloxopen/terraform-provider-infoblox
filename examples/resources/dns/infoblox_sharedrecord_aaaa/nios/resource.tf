// Create a Shared Record Group (Required as Parent)
resource "infoblox_sharedrecordgroup" "example" {
  nios = {
    name = "example-shared-record-group"
  }
}

// Create a Shared AAAA Record with Basic Fields
resource "infoblox_sharedrecord_aaaa" "shared_record_aaaa_with_basic_fields" {
  nios = {
    name                = "sharedrecord_aaaa_basic"
    ipv6addr            = "2001:db8::1"
    shared_record_group = infoblox_sharedrecordgroup.example.nios.name
  }
}

// Create a Shared AAAA Record with Additional Fields
resource "infoblox_sharedrecord_aaaa" "shared_record_aaaa_with_additional_fields" {
  nios = {
    name                = "sharedrecord_aaaa_additional_fields"
    ipv6addr            = "2001:db8::10"
    shared_record_group = infoblox_sharedrecordgroup.example.nios.name

    // Additional Fields
    ext_attrs = {
      Site = "location-1"
    }

    comment = "Example Sharedrecord AAAA"
    disable = false
    ttl     = 7200
  }
}
