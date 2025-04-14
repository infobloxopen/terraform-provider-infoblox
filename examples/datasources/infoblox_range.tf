resource "infoblox_range" "range" {
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
  member = jsonencode({
    name = "infoblox.localdomain"
  })
  server_association_type= "MEMBER"
  ext_attrs = jsonencode({
    "Site" = "Blr"
  })
  use_options = true
}

data "infoblox_range" "range_rec_temp" {
  filters = {
    start_addr = "17.0.0.221"
  }
  depends_on = [infoblox_range.range]
}

output "range_rec_res" {
  value = data.infoblox_range.range_rec_temp
}

//accessing range through EA
data "infoblox_range" "range_rec_temp_ea" {
  filters = {
    "*Site" = "Blr"
  }
}

output "range_rec_res1" {
  value = data.infoblox_range.range_rec_temp_ea
}