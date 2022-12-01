data "infoblox_a_record" "a_rec1" {
  dns_view = "default" // this is a required parameter
  fqdn = "static1.example1.org"
  ip_addr = "1.3.5.4"
}
