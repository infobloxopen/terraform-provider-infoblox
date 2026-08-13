// Create Record Alias with Basic Fields
resource "infoblox_record_alias" "create_alias_record" {
  nios = {
    target_type = "A"
    view        = "default"
  }
}

// Create Record Alias with Additional Fields
resource "infoblox_record_alias" "create_alias_record_with_additional_fields" {
  nios = {
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
