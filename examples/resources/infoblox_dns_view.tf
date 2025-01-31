# Creating DNS view resource with minimal set of parameters
resource "infoblox_dns_view" "view1" {
  name = "test_view"
}

# Creating DNS view resource with full set of parameters
resource "infoblox_dns_view" "view2" {
  name = "customview"
  network_view = "default"
  comment = "test dns view example"
  ext_attrs = jsonencode({
    "Site" = "Main test site"
  })
}

# Creating DNS View under non default network view
resource "infoblox_dns_view" "view3" {
  name = "custom_view"
  network_view = "non_defaultview"
  comment = "example under custom network view"
  ext_attrs = jsonencode({
    "Site" = "Cal Site"
  })
}
