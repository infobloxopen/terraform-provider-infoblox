// NOTE: The shared record group "shared_group" must already exist on the grid.
// shared_record_group is a required, immutable field on every shared record.

// Create a Shared AAAA Record with Basic Fields
resource "infoblox_sharedrecord_aaaa" "shared_record_aaaa_with_basic_fields" {
  nios = {
    name                = "sharedrecord_aaaa_basic"
    ipv6addr            = "2001:db8::1"
    shared_record_group = "shared_group"
  }
}

// Create a Shared AAAA Record with Additional Fields
resource "infoblox_sharedrecord_aaaa" "shared_record_aaaa_with_additional_fields" {
  nios = {
    name                = "sharedrecord_aaaa_additional_fields"
    ipv6addr            = "2001:db8::10"
    shared_record_group = "shared_group"

    // Additional Fields
    ext_attrs = {
      Site = "location-1"
    }

    comment = "Example Sharedrecord AAAA"
    disable = false
    ttl     = 7200
  }
}
