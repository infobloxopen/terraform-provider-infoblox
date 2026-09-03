// Create an IPAM Bulk Hostname Template with Basic Feilds
resource "infoblox_bulk_hostname_template" "bulk_hostname_template1" {
  nios = {
    template_name   = "one-octet"
    template_format = "-$4"
  }
}

// Create a Bulk Hostname Template with Two Octets
resource "infoblox_bulk_hostname_template" "bulk_hostname_template2" {
  nios = {
    template_name   = "two-octet"
    template_format = "host-$3-$4"
  }
}
