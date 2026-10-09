// List specific NS Records using filters
list "infoblox_record_ns" "list_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      view = "default"
    }
  }
  limit = 10
}

// List specific NS Records using nameserver filter
list "infoblox_record_ns" "list_records_using_nameserver" {
  provider = infoblox
  config {
    filters = {
      nameserver = "ns1.example.com"
    }
  }
}

// List NS Records with resource details included
list "infoblox_record_ns" "list_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
