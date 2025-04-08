// static A-record, minimal set of parameters
resource "infoblox_a_record" "rec1" {
  fqdn    = "static1.example1.org"
  ip_addr = "1.3.5.4" // not necessarily from a network existing in NIOS DB
}

// all the parameters for a static A-record
resource "infoblox_a_record" "rec2" {
  fqdn     = "static2.example4.org"
  ip_addr  = "1.3.5.1"
  comment  = "example static A-record rec2"
  dns_view = "nondefault_dnsview2"
  ttl      = 120 // 120s

  ext_attrs = jsonencode({
    "Location" = "65.8665701230204, -37.00791763398113"
  })
}

// all the parameters for a dynamic A-record
resource "infoblox_a_record" "rec3" {
  fqdn         = "dynamic1.example2.org"
  cidr         = infoblox_ipv4_network.net2.cidr         // the network  must exist, you may use the example for infoblox_ipv4_network resource.
  network_view = infoblox_ipv4_network.net2.network_view // not necessarily in the same network view as the DNS view resides in.
  comment      = "example dynamic A-record rec3"
  dns_view     = "nondefault_dnsview1"
  ttl          = 0 // 0 = disable caching
  ext_attrs    = jsonencode({})
}


// dynamically created A-record using next_available_ip
resource "infoblox_a_record" "recordA" {
  fqdn         = "gg11.test.com"
  comment      = "example dynamic A-record rec18"
  ttl          = 120
  network_view = "custom"

  filter_params = jsonencode({
    "*Site" : "Turkey"
  })

  ext_attrs = jsonencode({
    "Location" = "65.8665701230204, -37.00791763398113"
  })

}
