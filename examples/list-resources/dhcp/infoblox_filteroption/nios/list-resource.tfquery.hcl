// List specific DHCP Option Filters using filters
list "infoblox_filteroption" "list_filteroptions_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "filteroption_example"
    }
  }
}

// List specific DHCP Option Filters using Extensible Attributes
list "infoblox_filteroption" "list_filteroptions_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List DHCP Option Filters with resource details included
list "infoblox_filteroption" "list_filteroptions_with_resource" {
  provider         = infoblox
  include_resource = true
}
