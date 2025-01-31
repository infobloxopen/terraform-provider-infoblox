resource "infoblox_cname_record" "foo" {
  dns_view = "default.nondefault_netview"
  canonical = "strange-place.somewhere.in.the.net"
  alias = "foo.test.com"
  comment = "we need to keep an eye on this strange host"
  ttl = 0 # disable caching
  ext_attrs = jsonencode({
    Site = "unknown"
    Location = "TBD"
  })
}

data "infoblox_cname_record" "cname_rec"{
  filters = {
    name = "foo.test.com"
    canonical = "strange-place.somewhere.in.the.net"
    view = "default.nondefault_netview"
  }

  # This is just to ensure that the record has been be created
  # using 'infoblox_cname_record' resource block before the data source will be queried.
  depends_on = [infoblox_cname_record.foo]
}

output "cname_rec_out" {
  value = data.infoblox_cname_record.cname_rec
}
