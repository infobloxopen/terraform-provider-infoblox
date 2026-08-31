// Retrieve a specific BFD Template by filters
data "infoblox_bfdtemplate" "bfd_template_with_filters" {
  filters = {
    name = "example_bfdtemplate"
  }
}

// Retrieve all BFD Templates
data "infoblox_bfdtemplate" "get_all_bfd_templates" {}
