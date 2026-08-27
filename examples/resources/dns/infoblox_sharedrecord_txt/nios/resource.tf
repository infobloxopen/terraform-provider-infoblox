// Create a Shared TXT Record with Basic Fields
resource "infoblox_sharedrecord_txt" "shared_record_txt_with_basic_fields" {
  nios = {
    name = "example-shared-record-txt"
    text = "Example TXT Shared Record"
  }
}

// Create a Shared TXT Record with Additional Fields
resource "infoblox_sharedrecord_txt" "shared_record_txt_with_additional_fields" {
  nios = {
    name = "example-shared-record-txt2"
    text = "Example TXT Shared Record"

    // Additional Fields
    ext_attrs = {
      Site = "location-1"
    }
    comment = "Shared TXT Record created by Terraform"
    disable = false
    ttl     = 3600
  }
}
