// Retrieve a specific RIR Organization by filters
data "infoblox_rir_organization" "by_name" {
  filters = {
    name = "example_rir_organization"
  }
}

// Retrieve specific RIR Organizations using Extensible Attributes
data "infoblox_rir_organization" "by_ext_attr" {
  ext_attr_filters = {
    "RIPE Email" = "support@infoblox.com"
  }
}

// Retrieve all RIR Organizations
data "infoblox_rir_organization" "all" {}
