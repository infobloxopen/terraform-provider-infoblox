// Retrieve a specific DHCP Option Filter by filters
data "infoblox_filteroption" "get_filteroption_using_filters" {
  filters = {
    name = "filteroption_example"
  }
}

// Retrieve specific DHCP Option Filters using Tags
data "infoblox_filteroption" "get_filteroption_using_tag_filters" {
  tag_filters = {
    location = "site1"
  }
}

// Retrieve all DHCP Option Filters
data "infoblox_filteroption" "get_all_filteroptions" {}
