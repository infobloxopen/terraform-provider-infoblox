resource "infoblox_ipv4_network" "net2" {
  cidr = "18.0.0.0/24"
}
resource "infoblox_ipv4_fixed_address" "fix4"{
  client_identifier_prepend_zero=true
  comment= "fixed address"
  dhcp_client_identifier="23"
  disable= true
  ext_attrs = jsonencode({
    "Site": "Blr"
  })
  match_client = "CLIENT_ID"
  name = "fixed_address_1"
  network = "18.0.0.0/24"
  network_view = "default"
  options {
    name         = "dhcp-lease-time"
    value        = "43200"
    vendor_class = "DHCP"
    num          = 51
    use_option   = true
  }
  options {
    name = "routers"
    num = "3"
    use_option = true
    value = "18.0.0.2"
    vendor_class = "DHCP"
  }
  use_option = true
  depends_on=[infoblox_ipv4_network.net2]
}
data "infoblox_ipv4_fixed_address" "testFixedAddress_read1" {
  filters = {
    "*Site" = "Blr"
  }
  depends_on = [infoblox_ipv4_fixed_address.fix4]
}
output "fa_rec_temp1" {
  value = data.infoblox_ipv4_fixed_address.testFixedAddress_read1
}
