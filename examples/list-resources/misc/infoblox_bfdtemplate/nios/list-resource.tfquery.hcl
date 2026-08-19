// List specific BFD Templates using filters
list "infoblox_bfdtemplate" "list_bfdtemplates_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "bfd-template-1"
    }
  }
  limit = 10
}

// List BFD Templates with resource details included
list "infoblox_bfdtemplate" "list_bfdtemplates_with_resource" {
  provider         = infoblox
  include_resource = true
}
