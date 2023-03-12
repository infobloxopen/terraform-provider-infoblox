// TXT-Record, minimal set of parameters
resource "infoblox_txt_record" "rec1" {
  fqdn = "sample1.example.org"
  text = "this is just a sample"
}

// some parameters for a TXT-Record
resource "infoblox_txt_record" "rec2" {
  dns_view = "default" // may be omitted
  fqdn = "sample2.example.org"
  text = "data for TXT-record #2"
  ttl = 120 // 120s
}

// all the parameters for a TXT-Record
resource "infoblox_txt_record" "rec3" {
  dns_view = "nondefault_dnsview1" // not 'default' thus must be specified
  fqdn = "example3.example2.org"
  text = "data for TXT-record #3"
  ttl = 300
  comment = "example TXT record #3"
  ext_attrs = jsonencode({
    "Location" = "65.8665701230204, -37.00791763398113"
  })
}
