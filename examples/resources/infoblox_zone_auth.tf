//forward mapping zone, with full set of parameters
resource "infoblox_zone_auth" "zone1" {
  fqdn = "test3.com"
  view = "default"
  zone_format = "FORWARD"
  ns_group = ""
  restart_if_needed = true
  soa_default_ttl = 36000
  soa_expire = 72000
  soa_negative_ttl = 600
  soa_refresh = 1800
  soa_retry = 900
  comment = "Zone Auth created newly"
  ext_attrs = jsonencode({
    Location = "AcceptanceTerraform"
  })
}

//IPV4 reverse mapping zone, with full set of parameters
resource "infoblox_zone_auth" "zone2" {
  fqdn = "10.0.0.0/24"
  view = "default"
  zone_format = "IPV4"
  ns_group = "nsgroup1"
  restart_if_needed = true
  soa_default_ttl = 37000
  soa_expire = 92000
  soa_negative_ttl = 900
  soa_refresh = 2100
  soa_retry = 800
  comment = "IPV4 reverse zone auth created"
  ext_attrs = jsonencode({
    Location = "TestTerraform"
  })
}

//IPV6 reverse mapping zone, with minimal set of parameters
resource "infoblox_zone_auth" "zone2" {
  fqdn = "2002:1100::/64"
  view = "non_defaultview"
  zone_format = "IPV6"
  ns_group = "nsgroup2"
  comment = "IPV6 reverse zone auth created"
  ext_attrs = jsonencode({
    Location = "Random TF location"
  })
}
