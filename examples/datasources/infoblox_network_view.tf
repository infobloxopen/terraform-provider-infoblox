resource "infoblox_network_view" "inet_nv" {
  name    = "inet_visible_nv"
  comment = "Internet-facing networks"

  ext_attrs = jsonencode({
    "Location" = "the North pole"
  })
}

data "infoblox_network_view" "inet_nv" {
  filters = {
    name = "inet_visible_nv"
  }

  // This is just to ensure that the network view has been be created
  // using 'infoblox_network_view' resource block before the data source will be queried.
  depends_on = [infoblox_network_view.inet_nv]
}

output "nview_res" {
  value = data.infoblox_network_view.inet_nv
}
