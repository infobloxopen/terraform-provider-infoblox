data "infoblox_srv_record" "srv_rec1" {
    name = "_sip._udp.example.test.com"
    dns_view = "nondefault_view"
}