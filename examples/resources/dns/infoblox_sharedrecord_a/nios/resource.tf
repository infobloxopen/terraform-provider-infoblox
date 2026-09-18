// Create a Shared Record Group (Required as Parent)
resource "infoblox_sharedrecordgroup" "example" {
  nios = {
    name = "example-shared-record-group"
  }
}

// Create a Shared A Record with Basic Fields
resource "infoblox_sharedrecord_a" "shared_record_a_with_basic_fields" {
  nios = {
    name                = "sharedrecord_a_basic"
    ipv4addr            = "10.0.0.10"
    shared_record_group = infoblox_sharedrecordgroup.example.nios.name
  }
}

// Create a Shared A Record with Additional Fields
resource "infoblox_sharedrecord_a" "shared_record_a_with_additional_fields" {
  nios = {
    name                = "sharedrecord_a_additional_fields"
    ipv4addr            = "20.0.0.0"
    shared_record_group = infoblox_sharedrecordgroup.example.nios.name

    // Additional Fields
    ext_attrs = {
      Site = "location-1"
    }
    comment = "Example Sharedrecord A"
    disable = false
    ttl     = 7200
  }
}
