data "infoblox_srv_record" "ds1" {
    // the arguments are taken from the examples for infoblox_srv_record resource

    // as we use a reference to a resource's field, we do not know if
    // it is 'default' (may be omitted) or not.
    dns_view = infoblox_srv_record.rec2.dns_view

    name = infoblox_srv_record.rec2.name
    target = infoblox_srv_record.rec2.target
    port = infoblox_srv_record.rec2.port

    // priority, weight, ttl, comment, ext_attrs arguments may be retrieved using this data source.

    depends_on = [infoblox_srv_record.rec1]
}
