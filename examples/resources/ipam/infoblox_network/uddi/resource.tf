// Static address
resource "infoblox_network" "example" {
  uddi = {
    address = "10.0.0.0"
    cidr    = 24
    space   = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"

    // Other optional fields
    name    = "example_subnet"
    comment = "Subnet for Site A"
    tags = {
      Site = "location-1"
    }
  }
}

// Next available subnet — the address is allocated from the parent address block
resource "infoblox_network" "example_na_s" {
  uddi = {
    cidr  = 24
    space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
    dynamic_allocation = {
      next_available_id = "ipam/address_block/0acbbbed-94a4-11f1-8e35-aee0083f614b"
    }

    // Other optional fields
    name    = "example_subnet"
    comment = "Subnet for Site A"
    tags = {
      Site = "location-1"
    }
  }
}
