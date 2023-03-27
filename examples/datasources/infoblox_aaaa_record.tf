data "infoblox_aaaa_record" "rec2" {
  fqdn = "static2.example4.org"
  ipv6_addr = "2002:1111::1402"
  dns_view = "nondefault_dnsview2" // optional but differs from the default value

  depends_on = [infoblox_aaaa_record.rec2]
}
