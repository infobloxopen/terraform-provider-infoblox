// Create a VLAN View with Basic Fields
resource "infoblox_vlan_view" "ipam_vlanview_basic" {
  nios = {
    start_vlan_id = 5
    end_vlan_id   = 10
    name          = "example_vlan_view"
  }
}

// Create a VLAN View with Additional Fields
resource "infoblox_vlan_view" "ipam_vlanview_with_additional_fields" {
  nios = {
    start_vlan_id = 50
    end_vlan_id   = 100
    name          = "example_vlan_view_2"

    // Additional Fields
    comment                 = "Example VLAN View"
    allow_range_overlapping = true

    //Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
}
