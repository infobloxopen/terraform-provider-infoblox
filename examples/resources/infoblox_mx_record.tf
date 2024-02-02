// MX-record, minimal set of parameters
resource "infoblox_mx_record" "rec1" {
  fqdn = "rec1.example.org"
  mail_exchanger = "sample.zone2.com"
  preference = 30
}

// MX-record, full set of parameters
resource "infoblox_mx_record" "rec2" {
  dns_view = "nondefault_dnsview1"
  fqdn = "rec2.example2.org"
  mail_exchanger = "sample.test.com"
  preference = 40
  comment = "example MX-record"
  ttl = 120
  ext_attrs = jsonencode({
    "Location" = "Las Vegas"
  })
}
