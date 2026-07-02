data "unified_dns_record_a" "example_by_attribute" {
  filters = {
    "name" = "test-rec-1"
  }
}

data "unified_dns_record_a" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

data "unified_dns_record_a" "example_all" {}
