resource "infoblox_network_view" "netview1234" {
  name    = "one_more_network_view"
  comment = "example network view"

  ext_attrs = jsonencode({
    "Location" = "the North pole"
  })
}
