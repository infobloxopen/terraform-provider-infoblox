// Manage IPAM Vlan Views (Required as Parent)
resource "infoblox_vlanview" "ipam_vlanview_parent" {
  nios = {
    start_vlan_id = 5
    end_vlan_id   = 10
    name          = "example_vlan_view"
  }
}

// Manage IPAM Vlan with Basic Fields
resource "infoblox_vlan" "ipam_vlan_basic" {
  nios = {
    id     = 6
    name   = "example_vlan"
    parent = infoblox_vlanview.ipam_vlanview_parent.id
  }
}

// Manage IPAM Vlan with Additional Fields
resource "infoblox_vlan" "ipam_vlan_with_additional_fields" {
  nios = {
    id     = 7
    name   = "example_vlan_additional"
    parent = infoblox_vlanview.ipam_vlanview_parent.id

    // Additional Fields
    comment     = "Example VLAN"
    contact     = "Infoblox"
    department  = "Engineering"
    description = "This is an example VLAN"
    reserved    = false

    // Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
}
