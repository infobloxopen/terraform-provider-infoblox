resource "infoblox_dns_view" "view1" {
  name = "test_blog"
  comment = "strange test dnsview"
  ext_attrs = jsonencode({
    "Site" = "New Location"
  })
}

data "infoblox_dns_view" "dview" {
  filters = {
    name = "test_blog"
  }
  depends_on = [infoblox_dns_view.view1]
}

output "DView" {
  value = data.infoblox_dns_view.dview
}
