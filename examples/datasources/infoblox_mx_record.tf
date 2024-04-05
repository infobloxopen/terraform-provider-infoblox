resource "infoblox_mx_record" "rec2" {
    dns_view = "nondefault_dnsview1"
    fqdn = "rec2.example2.org"
    mail_exchanger = "sample.test.com"
    preference = 40
    comment = "example MX-record"
    ttl = 120
    ext_attrs = jsonencode({
        "Location" = "Las Vegas"
    })
}

data "infoblox_mx_record" "ds2" {
    filters = {
        view = "nondefault_dnsview1"
        name = "rec2.example2.org"
        mail_exchanger = "sample.test.com"
    }

    // This is just to ensure that the record has been be created
    // using 'infoblox_mx_record' resource block before the data source will be queried.
    depends_on = [infoblox_mx_record.rec2]
}

output "mx_rec_res" {
    value = data.infoblox_mx_record.ds2
}