// Create IPv6 Option Space and Option Definition (Required as Parent).
resource "infoblox_ipv6_dhcp_optionspace" "example" {
  uddi = {
    name = "ipv6_option_space_example"
  }
}

resource "infoblox_ipv6_dhcp_optiondefinition" "example" {
  uddi = {
    code         = 234
    name         = "ipv6_option_code_example"
    option_space = infoblox_ipv6_dhcp_optionspace.example.id
    type         = "boolean"
  }
}

// Create an IPv6 Filter Option with Basic fields
resource "infoblox_ipv6_filteroption" "ipv6_filter_option_basic_fields" {
  uddi = {
    name = "ipv6_filter_option_example"
    rules = {
      match = "any"
      rules = [
        {
          compare      = "equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.example.id
          option_value = "true"
        }
      ]
    }
  }
}

// Create an IPv6 Filter Option with Additional Fields
resource "infoblox_ipv6_filteroption" "ipv6_filter_option_with_additional_fields" {
  uddi = {
    name    = "ipv6_filter_option_example_2"
    comment = "Example IPv6 filter option"
    rules = {
      match = "all"
      rules = [
        {
          compare      = "text_substring"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.example.id
          option_value = "true"
          # Offset applies only to the substring compare modes
          substring_offset = 2
        }
      ]
    }

    # DHCPv6 options handed out to clients matching this filter
    dhcp_options = [
      {
        type         = "option"
        option_code  = infoblox_ipv6_dhcp_optiondefinition.example.id
        option_value = "true"
      }
    ]

    # Other optional fields
    lease_time = 3600
    tags = {
      location = "site1"
    }
  }
}
