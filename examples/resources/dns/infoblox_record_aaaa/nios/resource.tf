// Create Record AAAA with Basic Fields
resource "infoblox_record_aaaa" "create_record_aaaa_with_basic_fields" {
  nios = {
    name     = "example_record.example.com"
    ipv6addr = "2002:1111::1401"
    comment  = "This is a test AAAA record"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create Record AAAA with Additional Fields
resource "infoblox_record_aaaa" "create_record_aaaa_with_additional_fields" {
  nios = {
    name     = "example_record_with_ttl.example.com"
    ipv6addr = "2002:1111::1401"
    ttl      = 10
    comment  = "Example AAAA record"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create Record AAAA using dynamic allocation (next_available_ip via func_call)
resource "infoblox_record_aaaa" "create_record_aaaa_with_dynamic_allocation" {
  nios = {
    name    = "example_record_with_func_call.example.com"
    comment = "AAAA record with a dynamically allocated address"
    dynamic_allocation = {
      network      = "2001:db8:abcd:12::/64"
      network_view = "default"
    }
  }
}
