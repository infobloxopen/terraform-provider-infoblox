data "unified_dns_record_a" "get_record_using_filters" {
  filters = {
    name = "test-rec-1.example.com"
  }
}

data "unified_dns_record_a" "get_record_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

data "unified_dns_record_a" "get_all_a_records" {}
