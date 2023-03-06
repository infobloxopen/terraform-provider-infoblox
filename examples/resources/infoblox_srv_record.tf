//some set of parameters of SRV Record
resource "infoblox_srv_record" "srv_rec1" {
    name = "_http._udp.test.com"
    dns_view = "default"
    priority = 100
    weight = 75
    port = 8080
    target = "sample.test2.com"
} 

// all set of parameters for SRV record
resource "infoblox_srv_record" "srv_rec2" {
    name = "_http._udp.example.test.com"
    priority = 120
    weight = 80
    port = 3060
    target = "sample.test.com"
    ttl = 140 //140s
    comment = "test comment for srv"
    extattrs = jsonencode({
        "Location" = "65.8665701230204, -37.00791763398113"
    })
}

//dns_view with nondefault_view
resource "infoblox_srv_record" "srv_rec3" {
    name = "_http._udp.test.com"
    dns_view = "nondefault_view"
    priority = 80
    weight = 60
    port = 88
    target = "sample.test3.com"
    comment = "test comment"
}
