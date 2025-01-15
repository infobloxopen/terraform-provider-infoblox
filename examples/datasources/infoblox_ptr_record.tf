resource "infoblox_ptr_record" "host1" {
  ptrdname = "host.example.org"
  ip_addr = "2a05:d014:275:cb00:ec0d:12e2:df27:aa60"
  comment = "workstation #3"
  ttl = 300 # 5 minutes
  ext_attrs = jsonencode({
    "Location" = "the main office"
  })
}

data "infoblox_ptr_record" "host1" {
  filters = {
    ptrdname="host.example.org"
    ipv6addr="2a05:d014:275:cb00:ec0d:12e2:df27:aa60"
  }

  # This is just to ensure that the record has been be created
  # using 'infoblox_ptr_record' resource block before the data source will be queried.
  depends_on = [infoblox_ptr_record.host1]
}

output "ptr_rec_res" {
  value = data.infoblox_ptr_record.host1
}
