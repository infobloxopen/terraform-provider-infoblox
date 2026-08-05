// Create an Auth Zone (Required as Parent)
resource "infoblox_zone_auth" "parent_zone" {
  nios = {
    fqdn = "example.com"
  }
}

resource "infoblox_record_ns" "example_1" {
  nios = {
    name       = infoblox_zone_auth.parent_zone.nios.fqdn
    nameserver = "ns1.${infoblox_zone_auth.parent_zone.nios.fqdn}"
    addresses = [{
      address         = "192.168.1.10"
      auto_create_ptr = false
    }]
  }
}

// Create an IPV4 reverse mapping zone (Required as Parent for auto_create_ptr)
resource "infoblox_zone_auth" "reverse_zone" {
  nios = {
    fqdn        = "192.168.1.0/24"
    zone_format = "IPV4"
  }
}

resource "infoblox_record_ns" "example_2" {
  nios = {
    name       = infoblox_zone_auth.parent_zone.nios.fqdn
    nameserver = "ns2.${infoblox_zone_auth.parent_zone.nios.fqdn}"
    addresses = [{
      address         = "192.168.1.11"
      auto_create_ptr = true
    }]
  }
  depends_on = [infoblox_zone_auth.reverse_zone]
}
