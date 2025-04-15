resource "infoblox_range_template" "range_template1" {
  name = "range-template3"
  number_of_addresses = 60
  offset = 76
  comment = "Temporary Range Template"
  cloud_api_compatible = false
  use_options = true
  ext_attrs = jsonencode({
    "Site" = "Kobe"
  })
  options {
    name = "domain-name-servers"
    value = "11.22.23.24"
    vendor_class = "DHCP"
    num = 6
    use_option = true
  }
  member = {
    ipv4addr = "10.197.81.146"
    ipv6addr = "2403:8600:80cf:e10c:3a00::1192"
    name = "infoblox.localdomain"
  }
  failover_association = "failover1"
  server_association_type = "FAILOVER"
}

data "infoblox_range_template" "range_template_read" {
  filters = {
    failover_association = "failover1"
  }

  // This is just to ensure that the range template has been be created
  // using 'infoblox_range_template' resource block before the data source will be queried.
  depends_on = [infoblox_range_template.range_template1]
}

output "range_template_res" {
  value = data.infoblox_range_template.range_template_read
}