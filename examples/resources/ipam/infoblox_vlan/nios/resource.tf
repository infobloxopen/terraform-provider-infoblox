// Create VLAN View (Required as Parent)
resource "infoblox_vlan_view" "ipam_vlanview_parent" {
  nios = {
    start_vlan_id = 5
    end_vlan_id   = 10
    name          = "example_vlan_view"
    ext_attrs = {
      Site = "us-east-1"
    }
  }
}

// Create VLAN with Basic Fields
resource "infoblox_vlan" "ipam_vlan_basic" {
  nios = {
    id     = 6
    name   = "example_vlan"
    parent = infoblox_vlan_view.ipam_vlanview_parent.id
  }
}

// Create VLAN with Additional Fields
resource "infoblox_vlan" "ipam_vlan_with_additional_fields" {
  nios = {
    id     = 7
    name   = "example_vlan_additional"
    parent = infoblox_vlan_view.ipam_vlanview_parent.id

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

// Create VLAN with Next Available VLAN ID
resource "infoblox_vlan" "example_dynamic_allocation" {
  nios = {
    name   = "example_vlan_2"
    parent = infoblox_vlan_view.ipam_vlanview_parent.id

    dynamic_allocation = {
      vlan_view = infoblox_vlan_view.ipam_vlanview_parent.nios.name
    }
  }
}

resource "infoblox_vlan" "example_dynamic_allocation_2" {
  nios = {
    name   = "example_vlan_3"
    parent = infoblox_vlan_view.ipam_vlanview_parent.id

    dynamic_allocation = {
      filter_params = {
        "*Site" : "us-east-1"
      }
    }
  }
}
