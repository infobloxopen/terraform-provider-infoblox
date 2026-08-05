// Create an Auth Zone (Required as Parent)
resource "infoblox_zone_auth" "example" {
  nios = {
    fqdn = "example.com"
  }
}

resource "infoblox_record_a" "example_1" {
  nios = {
    name     = "rec-1.${infoblox_zone_auth.example.nios.fqdn}"
    ipv4addr = "10.0.0.18"
    comment  = "This is a test A record"
    creator  = "DYNAMIC"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

resource "infoblox_record_a" "example_dynamic_allocation" {
  nios = {
    name    = "rec-dynamic-1.${infoblox_zone_auth.example.nios.fqdn}"
    comment = "A record with a dynamically allocated address"
    dynamic_allocation = {
      network = "13.0.0.0/24"
    }
  }
}

resource "infoblox_record_a" "example_dynamic_allocation_2" {
  nios = {
    name    = "rec-dynamic-2.${infoblox_zone_auth.example.nios.fqdn}"
    comment = "A record with a dynamically allocated address"
    dynamic_allocation = {
      filter_params = {
        "*Site" : "location-1"
      }
    }
  }
}
