// Create a DHCP Option Filter with Basic Fields
resource "infoblox_filteroption" "filteroption_basic_fields" {
  nios = {
    name = "filteroption_example"
  }
}

// Create a DHCP Option Filter with Additional Fields
resource "infoblox_filteroption" "filteroption_additional_fields" {
  nios = {
    name        = "filteroption_example_2"
    comment     = "Example DHCP option filter"
    expression  = "(option domain-name=\"example.com\")"
    lease_time  = 3600
    next_server = "1.1.1.1"
    bootfile    = "pxelinux.0"
    bootserver  = "1.1.1.2"
    option_list = [
      {
        name  = "time-offset"
        num   = 2
        value = "1200"
      }
    ]
    ext_attrs = {
      Site = "location-1"
    }
  }
}
