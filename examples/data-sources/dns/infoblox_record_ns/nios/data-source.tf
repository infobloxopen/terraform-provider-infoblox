data "infoblox_record_ns" "get_ns_record_using_filters" {
  filters = {
    name = "example.com"
  }
}

data "infoblox_record_ns" "get_ns_record_using_nameserver_filter" {
  filters = {
    nameserver = "ns1.example.com"
  }
}

data "infoblox_record_ns" "get_all_ns_records" {}
