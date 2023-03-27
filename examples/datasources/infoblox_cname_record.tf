data "infoblox_cname_record" "rec1" {
  dns_view = "default" // required parameter here
  canonical = "bla-bla-bla.somewhere.in.the.net"
  alias = "hq-server.example1.org"

  depends_on = [infoblox_cname_record.rec1]
}
