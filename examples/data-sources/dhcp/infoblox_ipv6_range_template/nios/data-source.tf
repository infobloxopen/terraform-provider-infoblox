// Retrieve a specific IPV6 Range Template by filters
data "infoblox_ipv6_range_template" "get_ipv6_range_template_using_filters" {
  filters = {
    name = "example_ipv6_range_template"
  }
}

// Retrieve all IPV6 Range Templates
data "infoblox_ipv6_range_template" "get_all_ipv6_range_templates" {}
