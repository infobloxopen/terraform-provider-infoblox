list "infoblox_record_a" "list_records_using_filters" {
  provider = unified
  config {
    filters = {
      view = "default"
    }
  }
  limit = 10
}
