// Create a BFD template with Basic Fields
resource "infoblox_bfdtemplate" "bfd_template_with_basic_fields" {
  nios = {
    name = "example_bfdtemplate"
  }
}

// Create a BFD template with Additional Fields
resource "infoblox_bfdtemplate" "bfd_template_with_additional_fields" {
  nios = {
    name                  = "example_bfdtemplate_2"
    authentication_key_id = 4
    authentication_type   = "METICULOUS-MD5"
    authentication_key    = "example-auth-key"
    detection_multiplier  = 5
    min_rx_interval       = 1000
    min_tx_interval       = 1000
  }
}
