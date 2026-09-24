// List specific RIR Organizations using filters
list "infoblox_rir_organization" "list_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_rir_organization"
    }
  }
}

// List RIR Organizations with resource details included
list "infoblox_rir_organization" "list_with_resource" {
  provider         = infoblox
  include_resource = true
}
