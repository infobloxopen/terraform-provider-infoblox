// Create an Auth Zone (Required as Parent)
resource "infoblox_zone_auth" "parent_zone" {
  nios = {
    fqdn = "example.com"
  }
}

// Create an IPv6 Network (Required for Dynamic Allocation)
resource "infoblox_ipv6_network" "example_network" {
  nios = {
    network = "2001:db8:abcd:12::/64"
  }
}

// Create Record AAAA with Basic Fields
resource "infoblox_record_aaaa" "create_record_aaaa_with_basic_fields" {
  nios = {
    name     = "example_record.${infoblox_zone_auth.parent_zone.nios.fqdn}"
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
    name     = "example_record_with_ttl.${infoblox_zone_auth.parent_zone.nios.fqdn}"
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
    name    = "example_record_with_func_call.${infoblox_zone_auth.parent_zone.nios.fqdn}"
    comment = "AAAA record with a dynamically allocated address"
    dynamic_allocation = {
      network      = infoblox_ipv6_network.example_network.nios.network
      network_view = "default"
    }
  }
}
