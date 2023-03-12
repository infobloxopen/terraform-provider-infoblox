data "infoblox_mx_record" "ds1" {
    // the arguments are taken from the examples for infoblox_mx_record resource
    dns_view = infoblox_mx_record.rec2.dns_view // required
    fqdn = infoblox_mx_record.rec2.fqdn
    mail_exchanger = infoblox_mx_record.rec2.mail_exchanger

    // preference, ttl, comment, ext_attrs arguments may be retrieved using this data source.
}