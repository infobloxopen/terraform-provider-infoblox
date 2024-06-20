resource "infoblox_zone_auth" "zone1" {
  fqdn = "example.org"
  view = "default"
}

resource "infoblox_ip_allocation" "allocation1" {
  dns_view = "default"
  enable_dns = true
  fqdn = "host1.example.org"
  ipv4_addr = "10.10.0.7"
  ipv6_addr = "1::1"
  ext_attrs = jsonencode({"Location" = "USA"})

  depends_on = [infoblox_zone_auth.zone1]
}

resource "infoblox_ip_association" "association1" {
  internal_id = infoblox_ip_allocation.allocation1.id
  mac_addr = "12:00:43:fe:9a:8c"
  duid = "12:00:43:fe:9a:81"
  enable_dhcp = false
  depends_on = [infoblox_ip_allocation.allocation1]
}

data "infoblox_host_record" "host_rec_temp" {
  filters = {
    name = "host1.example.org"
  }
}

output "host_rec_res" {
  value = data.infoblox_host_record.host_rec_temp
}

// fetching Host-Records through EAs
data "infoblox_host_record" "host_rec_ea" {
  filters = {
    "*Location" = "USA"
  }
}

output "host_ea_out" {
  value = data.infoblox_host_record.host_rec_ea
}