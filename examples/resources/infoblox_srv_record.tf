// minimal set of parameters
resource "infoblox_srv_record" "rec1" {
    name = "_http._tcp.example.org"
    priority = 100
    weight = 75
    port = 8080
    target = "www.example.org"
} 

// all set of parameters for SRV record
resource "infoblox_srv_record" "rec2" {
    dns_view = "nondefault_dnsview1" // not 'default' thus must be specified
    name = "_sip._udp.example2.org"
    priority = 12
    weight = 10
    port = 5060
    target = "sip.example2.org"
    ttl = 3600
    comment = "example SRV record"
    ext_attrs = jsonencode({
        "Location" = "65.8665701230204, -37.00791763398113"
    })
}
