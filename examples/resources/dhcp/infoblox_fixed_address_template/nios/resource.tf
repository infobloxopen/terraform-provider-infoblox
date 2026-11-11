// Create a DHCP Option Filter to reference from the template
resource "infoblox_filteroption" "example_option_filter" {
  nios = {
    name = "example_option_filter"
  }
}

// Create a DHCP Fixed Address Template with Basic Fields
resource "infoblox_fixed_address_template" "fixed_address_template_basic_fields" {
  nios = {
    name = "example_fixed_address_template_1"
  }
}

// Create a DHCP Fixed Address Template with Additional Fields
resource "infoblox_fixed_address_template" "fixed_address_template_additional_fields" {
  nios = {
    name = "example_fixed_address_template_2"
    // Additional Fields
    comment             = "Fixed Address Template Created by Terraform"
    number_of_addresses = 3
    offset              = 2
    bootfile            = "pxelinux.0"
    bootserver          = "10.0.0.1"
    nextserver          = "10.0.0.2"
    ddns_hostname       = "ddns_host.name"
    ddns_domainname     = "example.com"
    enable_ddns         = true
    deny_bootp          = false
    options = [
      {
        name  = "time-offset"
        num   = 2
        value = "50"
      },
      {
        name  = "subnet-mask"
        value = "255.255.255.0"
      }
    ]
    logic_filter_rules = [
      {
        filter = infoblox_filteroption.example_option_filter.nios.name
        type   = "Option"
      }
    ]
    ext_attrs = {
      Site = "location-1"
    }
  }
}
