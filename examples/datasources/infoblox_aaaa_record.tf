resource "infoblox_aaaa_record" "vip_host" {
  fqdn = "very-interesting-host.example.com"
  ipv6_addr = "2a05:d014:275:cb00:ec0d:12e2:df27:aa60"
  comment = "some comment"
  ttl = 120 # 120s
  ext_attrs = jsonencode({
    "Location" = "65.8665701230204, -37.00791763398113"
  })
}

data "infoblox_aaaa_record" "qa_rec_temp" {
  filters = {
    name ="very-interesting-host.example.com"
    ipv6addr ="2a05:d014:275:cb00:ec0d:12e2:df27:aa60"
  }

  # This is just to ensure that the record has been be created
  # using 'infoblox_aaaa_record' resource block before the data source will be queried.
  depends_on = [infoblox_aaaa_record.vip_host]
}

output "qa_rec_res" {
  value = data.infoblox_aaaa_record.qa_rec_temp
}
