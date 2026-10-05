// Create a Custom Option Space (required parent for the option definition)
resource "infoblox_dhcp_optionspace" "option_space" {
  uddi = {
    name = "example_option_space"
  }
}

// Create a Custom Option Definition (used as option_code in filteroption rules)
resource "infoblox_dhcp_optiondefinition" "option_definition" {
  uddi = {
    code         = 234
    name         = "example_option_code"
    type         = "boolean"
    option_space = infoblox_dhcp_optionspace.option_space.id
  }
}

// Create a DHCP Option Filter with the required fields
resource "infoblox_filteroption" "filteroption_basic_fields" {
  uddi = {
    name = "filteroption_example"
    rules = {
      match = "any"
      rules = [
        {
          compare      = "equals"
          option_code  = infoblox_dhcp_optiondefinition.option_definition.id
          option_value = "true"
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
          option_code  = infoblox_dhcp_optiondefinition.option_definition.id
          option_value = "true"
          # Offset applies only to the substring compare modes
          substring_offset = 2
        }
      ]
    }

    # DHCP options handed out to clients matching this filter
    dhcp_options = [
      {
        type         = "option"
        option_code  = infoblox_dhcp_optiondefinition.option_definition.id
        option_value = "true"
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
