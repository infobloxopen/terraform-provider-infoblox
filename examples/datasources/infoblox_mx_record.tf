data "infoblox_mx_record" "ds1" {
    // the arguments are taken from the examples for infoblox_mx_record resource

    // as we use a reference to a resource's field, we do not know if
    // it is 'default' (may be omitted) or not.
    dns_view = infoblox_mx_record.rec2.dns_view

    fqdn = infoblox_mx_record.rec2.fqdn
    mail_exchanger = infoblox_mx_record.rec2.mail_exchanger
    preference = infoblox_mx_record.rec2.preference

    // preference, ttl, comment, ext_attrs arguments may be retrieved using this data source.

    depends_on = [infoblox_mx_record.rec1]
}