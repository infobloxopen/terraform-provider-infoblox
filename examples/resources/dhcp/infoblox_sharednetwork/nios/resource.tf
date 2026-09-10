// Create IPV4 Networks (Required as Parent)
resource "infoblox_network" "parent_network1" {
  nios = {
    network      = "21.21.14.0/24"
    network_view = "default"
    comment      = "Parent network for shared network 1"
  }
}

resource "infoblox_network" "parent_network2" {
  nios = {
    network      = "21.21.13.0/24"
    network_view = "default"
    comment      = "Parent network for shared network 1"
  }
}

resource "infoblox_network" "parent_network3" {
  nios = {
    network      = "15.14.1.0/24"
    network_view = "default"
    comment      = "Parent network for shared network 2"
  }
}

resource "infoblox_network" "parent_network4" {
  nios = {
    network      = "16.0.0.0/24"
    network_view = "default"
    comment      = "Parent network for shared network 2"
  }
}

// Create Shared Network with Basic Fields
resource "infoblox_sharednetwork" "shared_network_basic_fields" {
  nios = {
    name = "example_shared_network1"
    networks = [
      {
        ref = infoblox_network.parent_network1.id
      },
      {
        ref = infoblox_network.parent_network2.id
      }
    ]
    network_view = "default"
    ext_attrs = {
      Site = "location-1"
    }
  }
  depends_on = [
    infoblox_network.parent_network1,
    infoblox_network.parent_network2
  ]
}

// Create Shared Network with Additional Fields
resource "infoblox_sharednetwork" "shared_network_additional_fields" {
  nios = {
    name = "example_shared_network2"
    networks = [
      {
        ref = infoblox_network.parent_network3.id
      },
      {
        ref = infoblox_network.parent_network4.id
      }
    ]
    ignore_mac_addresses = ["66:77:88:99:aa:bb", "00:11:22:33:44:55"]
    options = [
      {
        name  = "domain-name-servers"
        num   = "6"
        value = "11.22.1.2,11.22.1.3"
      },
      {
        name  = "time-offset"
        num   = "2"
        value = "1000"
      },
      {
        name  = "domain-name"
        num   = "15"
        value = "aa.bb.com"
      },
    ]
    comment                    = "Shared network with additional fields"
    ddns_server_always_updates = true
    ddns_use_option81          = true
  }
  depends_on = [
    infoblox_network.parent_network3,
    infoblox_network.parent_network4
  ]
}
