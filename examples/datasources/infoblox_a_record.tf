resource "infoblox_a_record" "vip_host" {
  fqdn = "very-interesting-host.example.com"
  ip_addr = "10.3.1.65"
  comment = "special host"
  dns_view = "nondefault_dnsview2"
  ttl = 120 // 120s
  ext_attrs = jsonencode({
    "Location" = "65.8665701230204, -37.00791763398113"
  })
}


data "infoblox_a_record" "a_rec_temp" {
  filters = {
    name = "very-interesting-host.example.com"
    ipv4addr = "10.3.1.65" //alias is ip_addr
    view = "nondefault_dnsview2"
  }

  // This is just to ensure that the record has been be created
  // using 'infoblox_a_record' resource block before the data source will be queried.
  depends_on = [infoblox_a_record.vip_host]
}

output "a_rec_res" {
  value = data.infoblox_a_record.a_rec_temp
}