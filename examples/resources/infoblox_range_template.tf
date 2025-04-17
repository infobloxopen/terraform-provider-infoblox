resource "infoblox_ipv4_range_template" "range_template_full_params" {
  name = "range-template1"
  number_of_addresses = 10
  offset = 20
}

resource "infoblox_ipv4_range_template" "range_template_full_set_parameters" {
  name = "range-template223"
  number_of_addresses = 40
  offset = 30
  comment = "Temporary Range Template"
  cloud_api_compatible = true
  use_options = true
  ext_attrs = jsonencode({
    "Site" = "Kobe"
  })
  options {
    name = "domain-name-servers"
    value = "11.22.33.44"
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
