// Create VLAN View (Required as Parent)
resource "infoblox_vlan_view" "ipam_vlanview_parent" {
  nios = {
    start_vlan_id = 1
    end_vlan_id   = 200
    name          = "example_vlan_view_for_range"
  }
}

// Create VLAN Range with Basic Fields
resource "infoblox_vlan_range" "ipam_vlanrange_basic" {
  nios = {
    start_vlan_id = 5
    end_vlan_id   = 10
    name          = "example_vlan_range"
    vlan_view     = infoblox_vlan_view.ipam_vlanview_parent.id
  }
}

// Create VLAN Range with Additional Fields
resource "infoblox_vlan_range" "ipam_vlanrange_with_additional_fields" {
  nios = {
    start_vlan_id = 50
    end_vlan_id   = 100
    name          = "example_vlan_range2"
    vlan_view     = infoblox_vlan_view.ipam_vlanview_parent.id

    // Additional Fields
    comment          = "Example VLAN Range"
    pre_create_vlan  = true
    vlan_name_prefix = "vlan_range_"

    // Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
}
