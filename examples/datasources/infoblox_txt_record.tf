data "infoblox_txt_record" "txt_rec1" {
  dns_view = "default" //this is a required parameter
  fqdn = "first.test.com"
}
