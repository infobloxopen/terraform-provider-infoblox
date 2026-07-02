list "unified_dns_record_a" "list_records_using_filters" {
  provider = unified
  config {
    filters = {
      comment = "Created by Terraform"
    }
  }
  limit = 10
}
