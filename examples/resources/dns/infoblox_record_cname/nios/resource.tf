// Create Record CNAME with Basic Fields
resource "infoblox_record_cname" "create_record_basic" {
  nios = {
    name      = "example_record.example.com"
    canonical = "example-canonical-name.example.com"

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
    name      = "example_record2.example.com"
    canonical = "example-canonical-name2.example.com"
    view      = "default"

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
