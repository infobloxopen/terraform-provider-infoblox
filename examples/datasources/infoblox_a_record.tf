data "infoblox_a_record" "rec1" {
  // dns_view = "default" // optional, may be omitted
  fqdn = "static1.example1.org"
  ip_addr = "1.3.5.4"

  depends_on = [infoblox_a_record.rec1]
}
