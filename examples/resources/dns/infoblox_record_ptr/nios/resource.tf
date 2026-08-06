// Create the parent forward zone (Required as Parent for ptrdname)
resource "infoblox_zone_auth" "parent_zone" {
  nios = {
    fqdn = "example.com"
    view = "default"
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

resource "infoblox_zone_auth" "reverse_zone2" {
  nios = {
    fqdn        = "22.0.0.0/24"
    view        = "default"
    zone_format = "IPV4"
    comment     = "Reverse zone for 22.0.0.0/24 network"
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
    ptrdname = "example_record1.${infoblox_zone_auth.parent_zone.nios.fqdn}"
    ipv4addr = "10.20.1.2"
    view     = "default"
    ext_attrs = {
      Site = "location-1"
    }
  }
  depends_on = [infoblox_zone_auth.parent_zone, infoblox_zone_auth.reverse_zone1]
}

// Create an IPv6 PTR record with Basic Fields
resource "infoblox_record_ptr" "create_ptr_record_with_ipv6addr" {
  nios = {
    ptrdname = "example_record2.${infoblox_zone_auth.parent_zone.nios.fqdn}"
    ipv6addr = "2001::123"
    view     = "default"
    ext_attrs = {
      Site = "location-2"
    }
  }
  depends_on = [infoblox_zone_auth.parent_zone, infoblox_zone_auth.reverse_zone_ipv6]
}

// Create an IPv4 PTR record by name with Basic Fields
resource "infoblox_record_ptr" "create_ptr_record_with_name" {
  nios = {
    ptrdname = "example_record3.${infoblox_zone_auth.parent_zone.nios.fqdn}"
    name     = "11.0.0.22.in-addr.arpa"
    view     = "default"
    ext_attrs = {
      Site = "location-3"
    }
  }
  depends_on = [infoblox_zone_auth.parent_zone, infoblox_zone_auth.reverse_zone2]
}

// Create an IPv4 PTR record with Additional Fields
resource "infoblox_record_ptr" "create_ptr_record_with_additional_fields" {
  nios = {
    ptrdname = "example_record4.${infoblox_zone_auth.parent_zone.nios.fqdn}"
    name     = "12.0.0.22.in-addr.arpa"

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
  depends_on = [infoblox_zone_auth.parent_zone, infoblox_zone_auth.reverse_zone2]
}

// Create an IPv4 reverse mapping zone (Required as Parent)
resource "infoblox_zone_auth" "create_zone1" {
  nios = {
    fqdn        = "60.0.0.0/24"
    view        = "default"
    zone_format = "IPV4"
  }
}

// Create an IPv4 PTR record by name with arpa notation
resource "infoblox_record_ptr" "create_ptr_record_with_ipv4_arpa" {
  nios = {
    name     = "5.0.0.60.in-addr.arpa"
    ptrdname = "host.example.com"
    view     = "default"
    ext_attrs = {
      Site = "location-3"
    }
  }
  depends_on = [infoblox_zone_auth.create_zone1]
}

// Create an IPv6 reverse mapping zone (Required as Parent)
resource "infoblox_zone_auth" "create_zone2" {
  nios = {
    fqdn        = "2002:1100::/64"
    view        = "default"
    zone_format = "IPV6"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create an IPv6 PTR record by name with arpa notation
resource "infoblox_record_ptr" "create_ptr_record_with_ipv6_arpa" {
  nios = {
    ptrdname = "example_record.example.com"
    name     = "7.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.1.1.2.0.0.2.ip6.arpa"
    view     = "default"
    ext_attrs = {
      Site = "location-2"
    }
  }
  depends_on = [infoblox_zone_auth.create_zone2]
}
