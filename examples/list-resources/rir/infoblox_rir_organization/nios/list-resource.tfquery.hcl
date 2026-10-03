// List specific RIR Organizations using filters
list "infoblox_rir_organization" "list_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_rir_organization"
    }
  }
}

// List specific RIR Organizations using Extensible Attributes
list "infoblox_rir_organization" "list_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      "RIPE Email" = "support@infoblox.com"
    }
  }
}

// List RIR Organizations with resource details included
list "infoblox_rir_organization" "list_with_resource" {
  provider         = infoblox
  include_resource = true
}
