// PTR-record, minimal set of parameters
// Actually, either way may be used in reverse-mapping
// zones to specify an IP address:
//   1) 'ip_addr' (yes, literally) and
//   2) 'record_name' - in the form of a domain name (ex. 1.0.0.10.in-addr.arpa)
resource "infoblox_ptr_record" "rec1" {
  ptrdname = "rec1.example1.org"
  ip_addr = "10.0.0.1"
}

resource "infoblox_ptr_record" "rec2" {
  ptrdname = "rec2.example1.org"
  record_name = "2.0.0.10.in-addr.arpa"
}

// statically allocated PTR-record, full set of parameters
resource "infoblox_ptr_record" "rec3" {
  ptrdname = "rec3.example2.org"
  dns_view = "nondefault_dnsview1"
  ip_addr = "2002:1f93::3"
  comment = "workstation #3"
  ttl = 300 # 5 minutes
  ext_attrs = jsonencode({
    "Location" = "the main office"
  })
}

// dynamically allocated PTR-record, minimal set of parameters
resource "infoblox_ptr_record" "rec4" {
  ptrdname = "rec4.example2.org"
  cidr = infoblox_ipv4_network.net1.cidr
}

// statically allocated PTR-record, full set of parameters, non-default network view
resource "infoblox_ptr_record" "rec5" {
  ptrdname = "rec5.example2.org"
  dns_view = "nondefault_dnsview2"
  network_view = "nondefault_netview"
  ip_addr = "2002:1f93::5"
  comment = "workstation #5"
  ttl = 300 # 5 minutes
  ext_attrs = jsonencode({
    "Location" = "the main office"
  })
}

// PTR-record in a forward-mapping zone
resource "infoblox_ptr_record" "rec6_forward" {
  ptrdname = "example1.org"
  record_name = "www.example1.org"
}
