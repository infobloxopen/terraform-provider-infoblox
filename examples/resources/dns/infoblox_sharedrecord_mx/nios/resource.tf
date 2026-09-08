// NOTE: The shared record group "shared_group" must already exist on the grid.
// shared_record_group is a required, immutable field on every shared record.
// Record names avoid underscores so they satisfy a "Strict Hostname Checking"
// record name policy if the group enforces one.

// Create a Shared MX Record with Basic Fields
resource "infoblox_sharedrecord_mx" "sharedrecord_mx_basic_fields" {
  nios = {
    mail_exchanger      = "mail.example.com"
    name                = "sharedrecord-mx-basic"
    preference          = 10
    shared_record_group = "shared_group"
  }
}

// Create a Shared MX Record with Additional Fields
resource "infoblox_sharedrecord_mx" "sharedrecord_mx_additional_fields" {
  nios = {
    mail_exchanger      = "mail.example.com"
    name                = "sharedrecord-mx-additional-fields"
    preference          = 20
    shared_record_group = "shared_group"
    comment             = "Example MX Shared Record"
    disable             = true
    ext_attrs = {
      Site = "location-1"
    }
    ttl = 7200
  }
}
