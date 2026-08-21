// Create a DHCP Option Filter with the required fields
resource "infoblox_filteroption" "filteroption_basic_fields" {
  uddi = {
    name = "filteroption_example"
    rules = {
      match = "any"
      rules = [
        {
          compare      = "equals"
          option_code  = "dhcp/option_code/de50b0db-01cc-4da8-8213-aefd0880340f"
          option_value = "value1"
        }
      ]
    }
  }
}

// Create a DHCP Option Filter with Additional Fields
resource "infoblox_filteroption" "filteroption_additional_fields" {
  uddi = {
    name    = "filteroption_example_2"
    comment = "Example DHCP option filter"
    rules = {
      match = "all"
      rules = [
        {
          compare      = "text_substring"
          option_code  = "dhcp/option_code/de50b0db-01cc-4da8-8213-aefd0880340f"
          option_value = "value1"
          # Offset applies only to the substring compare modes
          substring_offset = 2
        }
      ]
    }

    # DHCP options handed out to clients matching this filter
    dhcp_options = [
      {
        type         = "option"
        option_code  = "dhcp/option_code/de50b0db-01cc-4da8-8213-aefd0880340f"
        option_value = "value1"
      }
    ]

    # Other optional fields
    lease_time                   = 3600
    header_option_filename       = "pxelinux.0"
    header_option_server_address = "192.168.1.10"
    header_option_server_name    = "tf-infoblox.example.com."
    tags = {
      location = "site1"
    }
  }
}
