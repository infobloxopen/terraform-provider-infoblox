// List specific DHCP Option Filters using filters
list "infoblox_filteroption" "list_filteroptions_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "filteroption_example"
    }
  }
  limit = 10
}

// List specific DHCP Option Filters using Tags
list "infoblox_filteroption" "list_filteroptions_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      location = "site1"
    }
  }
}

// List DHCP Option Filters with resource details included
list "infoblox_filteroption" "list_filteroptions_with_resource" {
  provider         = infoblox
  include_resource = true
}
