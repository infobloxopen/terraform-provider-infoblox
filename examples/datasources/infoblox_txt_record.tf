resource "infoblox_txt_record" "rec3" {
  dns_view = "nondefault_dnsview1"
  fqdn = "example3.example2.org"
  text = "\"data for TXT-record #3\""
  ttl = 300
  comment = "example TXT record #3"
  ext_attrs = jsonencode({
    "Location" = "65.8665701230204, -37.00791763398113"
  })
}

data "infoblox_txt_record" "ds3" {
  filters =  {
    dns_view = "nondefault_dnsview1"
    name = "example3.example2.org"
  }

  // This is just to ensure that the record has been be created
  // using 'infoblox_txt_record' resource block before the data source will be queried.
  depends_on = [infoblox_txt_record.rec3]
}

output "txt_rec_res" {
  value = data.infoblox_txt_record.ds3
}