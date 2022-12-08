# TXT Record with minimumally required input parameters
resource "infoblox_txt_record" "txt_record_minimal" {
  name         = "txt-minimal.test.com"
  text         = "minimal"
}

# TXT Record with all possible input parameters
resource "infoblox_txt_record" "txt_record_all" {
  name         = "txt-all.test.com"
  text         = "all"
  dns_view     = "default"
  ttl          = 3600
  comment      = "txt record comment"
  ext_attrs    = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Location" = "Test location"
    "Site" = "Test site"
  })
}