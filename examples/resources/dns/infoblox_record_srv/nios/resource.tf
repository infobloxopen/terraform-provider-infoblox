resource "infoblox_record_srv" "example_1" {
  nios = {
    name     = "_sip._tcp.example.com"
    port     = 5060
    priority = 10
    target   = "sip.example.com"
    weight   = 5
    comment  = "This is a test SRV record"
    ttl      = 300
    view     = "default"
  }
}
