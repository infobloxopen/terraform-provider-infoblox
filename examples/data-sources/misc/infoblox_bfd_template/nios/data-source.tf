// Retrieve a specific BFD Template by filters
data "infoblox_bfd_template" "bfd_template_with_filters" {
  filters = {
    name = "example_bfdtemplate"
  }
}

// Retrieve all BFD Templates
data "infoblox_bfd_template" "get_all_bfd_templates" {}
