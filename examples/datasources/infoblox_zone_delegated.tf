//creating Zone delegated resource
resource "infoblox_zone_delegated" "zone_delegated" {
  fqdn    = "zone_delegate.test_fwzone"
  comment = "zone delegated IPV4"
  delegate_to {
    name    = "test22.dz.ex.com"
    address = "10.0.0.1"
  }
  delegate_to {
    name    = "test2.dz.ex.com"
    address = "10.0.0.2"
  }
}

// accessing Zone delegated by specifying fqdn, view and comment
data "infoblox_zone_delegated" "data_zone_delegated" {
  filters = {
    fqdn    = "zone_delegate.test_fwzone"
    view    = "default"
    comment = "zone delegated IPV4"
  }
  // This is just to ensure that the record has been be created
  depends_on = [infoblox_zone_delegated.zone_delegated]
}

// returns matching Zone delegated with fqdn and view, if any
output "zone_delegated_data3" {
  value = data.infoblox_zone_delegated.data_zone_delegated
}