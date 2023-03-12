data "infoblox_srv_record" "ds1" {
    // the arguments are taken from the examples for infoblox_srv_record resource
    dns_view = infoblox_srv_record.rec2.dns_view // required
    name = infoblox_srv_record.rec2.name
    target = infoblox_srv_record.rec2.target
    port = infoblox_srv_record.rec2.port

    // priority, weight, ttl, comment, ext_attrs arguments may be retrieved using this data source.
}
