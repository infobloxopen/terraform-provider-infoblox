// static AAAA-record, minimal set of parameters
resource "infoblox_aaaa_record" "rec1" {
  fqdn      = "static1.example1.org"
  ipv6_addr = "2002:1111::1401" // not necessarily from a network existing in NIOS DB
}

// all the parameters for a static AAAA-record
resource "infoblox_aaaa_record" "rec2" {
  fqdn      = "static2.example4.org"
  ipv6_addr = "2002:1111::1402"
  comment   = "example static AAAA-record rec2"
  dns_view  = "nondefault_dnsview2"
  ttl       = 120 // 120s

  ext_attrs = jsonencode({
    "Location" = "65.8665701230204, -37.00791763398113"
  })
}

// all the parameters for a dynamic AAAA-record
resource "infoblox_aaaa_record" "rec3" {
  fqdn         = "dynamic1.example2.org"
  cidr         = infoblox_ipv6_network.net2.cidr         // the network  must exist, you may use the example for infoblox_ipv6_network resource.
  network_view = infoblox_ipv6_network.net2.network_view // not necessarily in the same network view as the DNS view resides in.
  comment      = "example dynamic AAAA-record rec3"
  dns_view     = "nondefault_dnsview1"
  ttl          = 0 // 0 = disable caching

  ext_attrs = jsonencode({})
}

// dynamically created AAAA-record using next_available_ip
resource "infoblox_aaaa_record" "recordAAAA" {
  fqdn         = "aaa123.test.com"
  comment      = "example dynamic AAAA-record rec18"
  ttl          = 120
  network_view = "test"

  filter_params = jsonencode({
    "*Site" : "Turkey"
  })

  ext_attrs = jsonencode({
    "Location" = "65.8665701230204, -37.00791763398113"
  })
}

// dynamically created AAAA-record using next_available_ip with dns_view
resource "infoblox_aaaa_record" "recordaaaa" {
  fqdn         = "aaa123.test.com"
  comment      = "example dynamic AAAA-record rec18"
  ttl          = 120
  network_view = "custom"
  dns_view     = "default.custom"

  filter_params = jsonencode({
    "*Site" : "Turkey"
  })

  ext_attrs = jsonencode({
    "Location" = "65.8665701230204, -37.00791763398113"
  })

}
