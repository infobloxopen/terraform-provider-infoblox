// Create an RP Zone (Required as Parent)
resource "infoblox_zone_rp" "example" {
  nios = {
    fqdn = "rpz.example.com"
  }
}

// Create Record RPZ PTR with IPv4 Address
resource "infoblox_record_rpz_ptr" "create_record_rpz_ptr_ipv4" {
  nios = {
    ptrdname = "record1.rpz.example.com"
    ipv4addr = "10.10.0.1"
    rp_zone  = infoblox_zone_rp.example.nios.fqdn
  }
}

// Create Record RPZ PTR with IPv6 Address
resource "infoblox_record_rpz_ptr" "create_record_rpz_ptr_ipv6" {
  nios = {
    ptrdname = "record2.rpz.example.com"
    ipv6addr = "2001:db8::1"
    rp_zone  = infoblox_zone_rp.example.nios.fqdn
  }
}

// Create Record RPZ PTR using ARPA Name
resource "infoblox_record_rpz_ptr" "create_record_rpz_ptr_name" {
  nios = {
    ptrdname = "record3.rpz.example.com"
    name     = "3.0.10.10.in-addr.arpa.rpz.example.com"
    rp_zone  = infoblox_zone_rp.example.nios.fqdn
  }
}

// Create Record RPZ PTR with Additional Fields
resource "infoblox_record_rpz_ptr" "create_record_rpz_ptr_additional" {
  nios = {
    ptrdname = "record4.rpz.example.com"
    ipv4addr = "10.10.0.4"
    rp_zone  = infoblox_zone_rp.example.nios.fqdn
    ttl      = 10
    comment  = "Example RPZ PTR record"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create a Substitute (PTR Record) Rule in a Custom View
resource "infoblox_view" "parent_view" {
  nios = {
    name = "custom-view"
  }
}

resource "infoblox_zone_rp" "parent_zone" {
  nios = {
    fqdn = "rpz-custom.example.com"
    view = infoblox_view.parent_view.nios.name
  }
}

resource "infoblox_record_rpz_ptr" "create_record_rpz_ptr_custom_view" {
  nios = {
    ptrdname = "record5.rpz-custom.example.com"
    ipv4addr = "10.10.0.5"
    rp_zone  = infoblox_zone_rp.parent_zone.nios.fqdn
    view     = infoblox_view.parent_view.nios.name
  }
}
