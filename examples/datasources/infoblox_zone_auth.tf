resource "infoblox_zone_auth" "zone1"{
  fqdn = "17.1.0.0/16"
  zone_format = "IPV4"
  view = "default"
  comment = "test sample reverse zone"
  ext_attrs = jsonencode({
    Location =  "Test Zone Location"
  })
}
# EA search example for Zone Auth datasource
data "infoblox_zone_auth" "dzone1" {
  filters = {
    "*Location" = "Test Zone Location"
  }
  depends_on = [infoblox_zone_auth.zone1]
}

# Generic example using filters, Zone Auth datasource
data "infoblox_zone_auth" "acctest" {
  filters = {
    view = "default"
    fqdn = infoblox_zone_auth.zone1.fqdn
    zone_format = "IPV4"
  }
}

output "ZoneA" {
  value = data.infoblox_zone_auth.dzone1
}

# Example for specific value fetching in output
output "ZoneB" {
  value = data.infoblox_zone_auth.acctest.results.0.fqdn
}
