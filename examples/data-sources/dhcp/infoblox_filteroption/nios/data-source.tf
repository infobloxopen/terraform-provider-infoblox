// Retrieve a specific DHCP Option Filter by filters
data "infoblox_filteroption" "get_filteroption_using_filters" {
  filters = {
    name = "filteroption_example"
  }
}

// Retrieve specific DHCP Option Filters using Extensible Attributes
data "infoblox_filteroption" "get_filteroption_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all DHCP Option Filters
data "infoblox_filteroption" "get_all_filteroptions" {}
