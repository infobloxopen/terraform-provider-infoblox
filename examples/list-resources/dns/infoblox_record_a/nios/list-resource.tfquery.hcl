list "infoblox_record_a" "list_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      view = "default"
    }
  }
  limit = 10
}
