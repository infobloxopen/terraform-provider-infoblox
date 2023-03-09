data "infoblox_mx_record" "mx_rec" {
    dns_view = "nondefault_view" //required parameter
    fqdn = "static1.example.org"
    mail_exchanger = "sample.test1.com" 
}