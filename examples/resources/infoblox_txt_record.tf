// static TXT-Record, minimal set of parameters
resource "infoblox_txt_record" "txt_rec1" {
  fqdn = "sample.test.com"
  text = "this is just sample"
}

// some parameters for a TXT-Record
resource "infoblox_txt_record" "txt_rec2" {
  fqdn = "sample.test1.com"
  text = "data for the TXT-Record"
  dns_view = "default"
  ttl = 120 // 120s
}

// all the parameters for a TXT-Record
resource "infoblox_txt_record" "txt_rec3" {
  fqdn = "example.test2.com"
  text = "data given for TXT Record"
  dns_view = "default"
  ttl = 300 //300s
  comment = "example TXT record txt_rec3"
  extattrs = jsonencode({
    "Location" = "65.8665701230204, -37.00791763398113"
  })
}