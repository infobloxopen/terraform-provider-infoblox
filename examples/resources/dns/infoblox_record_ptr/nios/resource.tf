// Create the parent zone (Required as Parent for ptrdname)
resource "infoblox_zone_auth" "parent_zone" {
  nios = {
    fqdn = "example.com"
    view = "default"
  }
}

resource "infoblox_record_ptr" "example_ptr" {
  nios = {
    ptrdname = "webserver.com"
    name     = "ptrrecord.${infoblox_zone_auth.parent_zone.nios.fqdn}"
    view     = "default"
  }
}

// Create IPv4 Reverse Mapping Zones (Required as Parent)
resource "infoblox_zone_auth" "reverse_zone1" {
  nios = {
    fqdn        = "10.20.1.0/24"
    view        = "default"
    zone_format = "IPV4"
    comment     = "Reverse zone for 10.20.1.0/24 network"
  }
}


// Create an IPv6 Reverse Mapping Zone (Required as Parent)
resource "infoblox_zone_auth" "reverse_zone_ipv6" {
  nios = {
    fqdn        = "2001::/64"
    view        = "default"
    zone_format = "IPV6"
    comment     = "Reverse zone for 2001::/64 network"
  }
}

// Create an IPv4 PTR record with Basic Fields
resource "infoblox_record_ptr" "create_ptr_record_with_ipv4addr" {
  nios = {
    ptrdname = "example_record1.com"
    ipv4addr = "10.20.1.2"
    view     = "default"
    ext_attrs = {
      Site = "location-1"
    }
  }
  depends_on = [infoblox_zone_auth.reverse_zone1]
}

// Create an IPv4 PTR record by name with Basic Fields
resource "infoblox_record_ptr" "create_ptr_record_with_name" {
  nios = {
    ptrdname = "abc.com"
    name     = "11.${infoblox_zone_auth.reverse_zone1.nios.display_domain}"
    view     = "default"
    ext_attrs = {
      Site = "location-3"
    }
  }
}

// Create an IPv6 PTR record with Basic Fields
resource "infoblox_record_ptr" "create_ptr_record_with_ipv6addr" {
  nios = {
    ptrdname = "abc.com"
    ipv6addr = "2001::123"
    view     = "default"
    ext_attrs = {
      Site = "location-2"
    }
  }
  depends_on = [infoblox_zone_auth.reverse_zone_ipv6]
}

// Create an IPv6 PTR record by name with arpa notation
resource "infoblox_record_ptr" "create_ptr_record_with_ipv6_arpa" {
  nios = {
    ptrdname = "abc2.com"
    name     = "7.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.${infoblox_zone_auth.reverse_zone_ipv6.nios.display_domain}"
    view     = "default"
    ext_attrs = {
      Site = "location-2"
    }
  }
}

// Create an IPv4 PTR record with Additional Fields
resource "infoblox_record_ptr" "create_ptr_record_with_additional_fields" {
  nios = {
    ptrdname = "abc3.com"
    name     = "12.${infoblox_zone_auth.reverse_zone1.nios.display_domain}"

    // Additional Fields
    view    = "default"
    use_ttl = true
    ttl     = 10
    creator = "DYNAMIC"
    comment = "Example PTR record"

    // Extensible Attributes
    ext_attrs = {
      Site = "location-4"
    }
  }
}
