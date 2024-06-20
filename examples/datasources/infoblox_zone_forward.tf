resource "infoblox_zone_forward" "forwardzone_forwardTo" {
  fqdn = "zone_forward.ex.org"
  forward_to {
    name = "test22.dz.ex.com"
    address = "10.0.0.1"
  }
  forward_to {
    name = "test2.dz.ex.com"
    address = "10.0.0.2"
  }
}

// accessing Zone Forward by specifying fqdn and view
data "infoblox_zone_forward" "data_zone_forward" {
  filters = {
    fqdn = "zone_forward.ex.org"
    view = "default"
  }
  // This is just to ensure that the record has been be created
  depends_on = [infoblox_zone_forward.forwardzone_forwardTo]
}

// returns matching Zone Forward with fqdn and view, if any
output "zone_forward_data3" {
  value = data.infoblox_zone_forward.data_zone_forward
}


resource "infoblox_zone_forward" "forwardzone_IPV4_nsGroup_externalNsGroup" {
  fqdn = "195.1.0.0/24"
  comment = "Forward zone IPV4"
  external_ns_group = "stub server"
  zone_format = "IPV4"
  ns_group = "test"
}

// accessing Zone Forward by specifying fqdn, view and comment
data "infoblox_zone_forward" "datazone_foward_fqdn_view_comment" {
  filters = {
    fqdn = "195.1.0.0/24"
    view = "default"
    comment = "Forward zone IPV4"
  }
  // This is just to ensure that the record has been be created
  depends_on = [infoblox_zone_forward.forwardzone_IPV4_nsGroup_externalNsGroup]
}

// returns matching Zone Forward with fqdn, view and comment, if any
output "zone_forward_data4" {
  value = data.infoblox_zone_forward.datazone_foward_fqdn_view_comment
}
