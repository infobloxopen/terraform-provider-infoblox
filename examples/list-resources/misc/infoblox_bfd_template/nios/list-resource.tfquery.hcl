// List specific BFD Templates using filters
list "infoblox_bfd_template" "list_bfdtemplates_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_bfdtemplate"
    }
  }
  limit = 10
}

// List BFD Templates with resource details included
list "infoblox_bfd_template" "list_bfdtemplates_with_resource" {
  provider         = infoblox
  include_resource = true
}
