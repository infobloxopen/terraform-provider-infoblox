terraform {
  required_providers {
    infoblox = {
      source = "infobloxopen/infoblox"
    }
  }
}

provider "infoblox" {
  server   = "100.68.0.149"
  username = "admin"
  password = "infoblox"
}

resource "infoblox_vlan" "test_vlan" {
    parent  = "vlanview/ZG5zLnZsYW5fdmlldyRJQUMtQkdQLUFTLU5VTS42MDEuOTk5:IAC-BGP-AS-NUM/601/999"
    name    = "auto-vlan"
    comment = "VLAN with auto-allocated ID"
    # vlan_id is omitted - will be auto-allocated
  }
  resource "infoblox_vlan" "test_vlanx" {
    parent  = "vlanview/ZG5zLnZsYW5fdmlldyRJQUMtQkdQLUFTLU5VTS42MDEuOTk5:IAC-BGP-AS-NUM/601/999"
    name    = "auto-vlanx"
    comment = "VLAN with auto-allocated ID"
    # vlan_id is omitted - will be auto-allocated
  }