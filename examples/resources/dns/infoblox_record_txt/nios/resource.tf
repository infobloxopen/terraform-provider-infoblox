// Create Record TXT with Basic Fields
resource "infoblox_record_txt" "create_record" {
  nios = {
    name = "example-txt-record.example.com"
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
    name = "example-txt-record-with-config.example.com"
    text = "Example TXT Record with Additional Config"

    // Additional Fields
    view    = "default"
    ttl     = 10
    creator = "DYNAMIC"
    comment = "Example TXT record"

    // Extensible Attributes
    ext_attrs = {
      Site = "location-2"
    }
  }
}
