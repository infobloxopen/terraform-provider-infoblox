# CNAME-record, minimal set of parameters
resource "infoblox_cname_record" "rec1" {
  canonical = "bla-bla-bla.somewhere.in.the.net"
  alias = "hq-server.example1.org"
}

# CNAME-record, full set of parameters
resource "infoblox_cname_record" "rec2" {
  dns_view = "default.nondefault_netview"
  canonical = "strange-place.somewhere.in.the.net"
  alias = "alarm-server.example3.org"
  comment = "we need to keep an eye on this strange host"
  ttl = 0 # disable caching
  ext_attrs = jsonencode({
    Site = "unknown"
    Location = "TBD"
  })
}
