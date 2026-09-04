// List specific Bulk Hostname Template using filters
list "infoblox_bulk_hostname_template" "list_bulk_hostname_templates_using_filters" {
  provider = infoblox
  config {
    filters = {
      template_name = "one_octet"
    }
  }
  limit = 10
}

// List Bulk Hostname Templates with resource details included
list "infoblox_bulk_hostname_template" "list_bulk_hostname_templates_with_resource" {
  provider         = infoblox
  include_resource = true
}

