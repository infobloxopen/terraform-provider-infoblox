package infoblox

const testCasePtrRecordTestData01 = `
resource "infoblox_ipv4_network" "net1" {
  cidr = "10.0.0.0/8"
  reserve_ip = 200
}

resource "infoblox_ipv4_network" "net2" {
  cidr = "172.16.0.0/16"
  reserve_ip = 200
}

resource "infoblox_ipv6_network" "net1" {
  cidr = "2002:1f93::/64"
  reserve_ipv6 = 200
}

resource "infoblox_ipv6_network" "net2" {
  cidr = "2002:1f94::/64"
  reserve_ipv6 = 200
}

resource "infoblox_ipv6_network" "net3" {
  network_view = "nondefault_netview"
  cidr = "2002:1f93::/64"
  reserve_ipv6 = 200
}

resource "infoblox_ipv6_network" "net4" {
  network_view = "nondefault_netview"
  cidr = "2002:1f94::/64"
  reserve_ipv6 = 200
}

resource "infoblox_ptr_record" "rec1" {
  ptrdname = "ptr-target1.example1.org"
  ip_addr = "10.0.0.1"
  comment = "non-empty comment"
  ttl = 5
  ext_attrs = jsonencode({
    Location = "Test location"
  })
}

resource "infoblox_ptr_record" "rec2" {
  ptrdname = "ptr-target2.example1.org"
  record_name = "32.0.0.10.in-addr.arpa"

  depends_on = [infoblox_ptr_record.rec1]
}

resource "infoblox_ptr_record" "rec3" {
  ptrdname = "ptr-target3.example1.org"
  record_name = "ptr-rec3-2.example1.org"

  depends_on = [infoblox_ptr_record.rec2]
}

resource "infoblox_ptr_record" "rec4" {
  ptrdname = "ptr-target4.example1.org"
  cidr = "172.16.0.0/16"

  depends_on = [infoblox_ipv4_network.net2, infoblox_ptr_record.rec3]
}

resource "infoblox_ptr_record" "rec5" {
  dns_view = "default"
  ptrdname = "ptr-target5.example1.org"
  cidr = "2002:1f93::0/64"
  comment = "workstation #5-2"
  ttl = 302
  ext_attrs = jsonencode({
    "Location" = "the new office"
  })

  depends_on = [infoblox_ptr_record.rec4]
}

resource "infoblox_ptr_record" "rec6" {
  ptrdname = "ptr-target6.example1.org"
  cidr = "2002:1f93::/64"

  depends_on = [infoblox_ptr_record.rec5]
}

resource "infoblox_ptr_record" "rec7" {
  ptrdname = "ptr-target7.example1.org"
  cidr = "2002:1f93::/64"

  depends_on = [infoblox_ipv6_network.net1, infoblox_ptr_record.rec6]
}

resource "infoblox_ptr_record" "rec8" {
  dns_view = "default"
  ptrdname = "ptr-target8.example1.org"
  record_name = "ptr-rec8.example1.org"
  comment = "workstation #8"
  ttl = 301
  ext_attrs = jsonencode({
    "Location" = "the main office"
  })

  depends_on = [infoblox_ptr_record.rec7]
}

////////////////////////////////////////////////

resource "infoblox_ptr_record" "rec9" {
  dns_view = "nondefault_dnsview1"
  ptrdname = "ptr-target9.example2.org"
  cidr = "2002:1f93::/64"
  comment = "workstation #9"
  ttl = 300
  ext_attrs = jsonencode({
    "Location" = "the main office"
  })

  depends_on = [infoblox_ptr_record.rec8]
}

resource "infoblox_ptr_record" "rec10" {
  dns_view = "nondefault_dnsview1"
  ptrdname = "ptr-target10.example2.org"
  ip_addr = "2002:1f93::b"
  comment = "workstation #10"
  ttl = 30
  ext_attrs = jsonencode({
    "Location" = "the main office"
  })

  depends_on = [infoblox_ptr_record.rec9]
}

resource "infoblox_ptr_record" "rec11" {
  dns_view = "nondefault_dnsview1"
  ptrdname = "ptr-target11.example2.org"
  record_name = "ptr-rec11.example2.org"
  comment = "workstation #11"
  ttl = 301
  ext_attrs = jsonencode({
    "Location" = "the main office"
  })

  depends_on = [infoblox_ptr_record.rec10]
}

resource "infoblox_ptr_record" "rec12" {
  dns_view = "nondefault_dnsview1"
  network_view = "default"
  ptrdname = "ptr-target12.example2.org"
  record_name = "32.0.0.10.in-addr.arpa"
  comment = "workstation #12"
  ttl = 30
  ext_attrs = jsonencode({
    "Location" = "the main office"
  })

  depends_on = [infoblox_ipv4_network.net1, infoblox_ptr_record.rec11]
}

resource "infoblox_ptr_record" "rec13" {
  dns_view = "nondefault_dnsview2"
  network_view = "nondefault_netview"
  ptrdname = "ptr-target13.example4.org"

  // must be within the same network
  ip_addr = "2002:1f93::13"

  comment = "workstation #13"
  ttl = 30
  ext_attrs = jsonencode({
    "Location" = "the main office"
  })

  depends_on = [infoblox_ipv6_network.net3, infoblox_ptr_record.rec12]
}

////////////////////////////////////////////////

resource "infoblox_ptr_record" "rec14" {
  ptrdname = "ptr-target14.example1.org"
  record_name = "44.0.0.10.in-addr.arpa"

  depends_on = [infoblox_ptr_record.rec13]
}
`
