// Retrieve a specific Bulk Hostname Template by filters
data "infoblox_bulk_hostname_template" "bulk_hostname_template" {
  filters = {
    template_name = "one-octet"
  }
}

// Retrieve all Bulk Hostname Templates
data "infoblox_bulk_hostname_template" "all_templates" {}
