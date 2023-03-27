data "infoblox_a_record" "rec1" {
  dns_view = "default" // this is a required parameter
  fqdn = "static1.example1.org"
  ip_addr = "1.3.5.4"

  depends_on = [infoblox_a_record.rec1]
}
