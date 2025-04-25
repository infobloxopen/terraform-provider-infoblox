resource "infoblox_ipv4_range" "range" {
  start_addr = "17.0.0.221"
  end_addr   = "17.0.0.240"
  options {
    name         = "dhcp-lease-time"
    value        = "43200"
    vendor_class = "DHCP"
    num          = 51
    use_option   = true
  }
  network              = "17.0.0.0/24"
  network_view = "default"
  comment              = "test comment"
  name                 = "test_range"
  disable              = false
  member = {
    name = "infoblox.localdomain"
  }
  server_association_type= "MEMBER"
  ext_attrs = jsonencode({
    "Site" = "Blr"
  })
  use_options = true
}