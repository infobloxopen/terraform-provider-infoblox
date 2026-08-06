// Create Record MX with Basic Fields
resource "infoblox_record_mx" "create_record_basic" {
  nios = {
    name           = "mx_record.example.com"
    mail_exchanger = "mail.example.com"
    preference     = 10

    // Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create Record MX with Additional Fields
resource "infoblox_record_mx" "create_record_additional_fields" {
  nios = {
    // Basic Fields
    name           = "mx_record1.example.com"
    mail_exchanger = "mail1.example.com"
    preference     = 10
    view           = "default"

    // Additional Fields
    ttl                = 3600
    creator            = "DYNAMIC"
    forbid_reclamation = false
    comment            = "MX record created by Terraform"

    // Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
}
