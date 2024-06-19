// accessing Zone Forward by specifying fqdn and view
data "infoblox_zone_forward" "data_zone_foward" {
  filters = {
    fqdn = "test.ex.com"
    view = "nondefault_view"
  }
}

// returns matching Zone Forward with fqdn and view, if any
output "zone_forward_data3" {
  value = data.infoblox_zone_forward.data_zone_foward
}
