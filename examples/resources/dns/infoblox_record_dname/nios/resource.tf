// Create an Auth Zone (Required as Parent)
resource "infoblox_zone_auth" "parent_zone1" {
  nios = {
    fqdn = "example.com"
  }
}

// Create a DNAME record with Basic Fields
resource "infoblox_record_dname" "create_record_dname_with_basic_fields" {
  nios = {
    target = "example-dname-1.com"
    name   = infoblox_zone_auth.parent_zone1.nios.fqdn
  }
}

// Create an Auth Zone (Required as Parent)
resource "infoblox_zone_auth" "parent_zone2" {
  nios = {
    fqdn = "example-1.com"
  }
}

// Create a DNAME record with Additional Fields
resource "infoblox_record_dname" "create_record_dname_with_additional_fields" {
  nios = {
    target = "example-dname-2.com"
    name   = infoblox_zone_auth.parent_zone2.nios.fqdn
    ext_attrs = {
      Site = "location-1"
    }
    comment = "DNAME record created by Terraform"
    ttl     = 10
  }
}

// Create an IPV4 reverse mapping zone (Required as Parent)
resource "infoblox_zone_auth" "parent_zone3" {
  nios = {
    fqdn        = "10.0.0.0/24"
    zone_format = "IPV4"
  }
}

// Create a DNAME record in IPV4 reverse mapping zone
resource "infoblox_record_dname" "create_record_dname1" {
  nios = {
    target = "example-dname-1.com"
    // We use display_domain for reverse mapping zones as arpa format is required for name
    name = infoblox_zone_auth.parent_zone3.nios.display_domain
  }
}

// Create an IPV6 reverse mapping zone (Required as Parent)
resource "infoblox_zone_auth" "parent_zone4" {
  nios = {
    fqdn        = "2002:1100::/64"
    zone_format = "IPV6"
  }
}

// Create a DNAME record in IPV6 reverse mapping zone
resource "infoblox_record_dname" "create_record_dname2" {
  nios = {
    target = "example-dname-1.com"
    // We use display_domain for reverse mapping zones as arpa format is required for name
    name = infoblox_zone_auth.parent_zone4.nios.display_domain
  }
}
