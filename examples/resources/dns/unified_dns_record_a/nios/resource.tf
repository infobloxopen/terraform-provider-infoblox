resource "unified_dns_record_a" "test1" {
  nios = {
    name     = "test-rec-1.example.com"
    ipv4addr = "10.0.0.18"
    comment  = "This is a test A record"
    creator  = "DYNAMIC"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

resource "unified_dns_record_a" "test_dynamic1" {
  nios = {
    name    = "test-rec-dynamic-1.example.com"
    comment = "A record with a dynamically allocated address"
    dynamic_allocation = {
      network = "13.0.0.0/24"
    }
  }
}

resource "unified_dns_record_a" "test_dynamic2" {
  nios = {
    name    = "test-rec-dynamic-2.example.com"
    comment = "A record with a dynamically allocated address"
    dynamic_allocation = {
      filter_params = {
        "*Site" : "location-1"
      }
    }
  }
}
