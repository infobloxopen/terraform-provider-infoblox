// List specific IPv6 Range Templates using filters
list "infoblox_ipv6_range_template" "list_ipv6_range_template_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_range_template"
    }
  }
  limit = 10
}

// List IPv6 Range Templates with resource details included
list "infoblox_ipv6_range_template" "list_ipv6_range_template_with_resource" {
  provider         = infoblox
  include_resource = true
}
