resource "infoblox_srv_record" "rec2" {
    dns_view = "nondefault_dnsview1"
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

data "infoblox_srv_record" "ds1" {
    filters = {
        dns_view = "nondefault_dnsview1"
        name = "_sip._udp.example2.org"
        port = 5060
        target = "sip.example2.org"
    }

    // This is just to ensure that the record has been be created
    // using 'infoblox_srv_record' resource block before the data source will be queried.
    depends_on = [infoblox_srv_record.rec2]
}

output "srv_rec_res" {
    value = data.infoblox_srv_record.ds1
}