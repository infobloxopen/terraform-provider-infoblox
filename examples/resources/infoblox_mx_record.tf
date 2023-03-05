// static MX-record, minimal set of parameters
resource "infoblox_mx_record" "mx_rec" {
  fqdn = "demo.test.com"
  mail_exchanger = "sample.zone2.com"
  preference = 30
  dns_view = ""
  comment = "test comment"
}

//static MX-record, with all set of parameters
resource "infoblox_mx_record" "mx_rec2" {
    fqdn = "demo.test1.com"
    mail_exchanger = "sample.test.com"
    preference = 40
    dns_view = "default"
    comment = "for the mx record"
    ttl = 120 //120s
    extattrs = jsonencode({
        "Location" = "Las Vegas"
    })
}

//without extattrs parameter
resource "infoblox_mx_record" "mx_rec3" {
    fqdn = "demo.test2.com"
    mail_exchanger = "sample.test2.com"
    preference = 100
    dns_view = "nondefault_view"
    comment = "data about mx record"
    ttl = 150 //150s
}

