resource "infoblox_address" "static" {
  uddi = {
    address = "10.0.0.5"
    space   = "ipam/ip_space/00c2c546-6ad1-11f1-88c3-1e3cda826891"
    comment = "statically assigned address"
    tags = {
      Site = "location-1"
    }
  }
}

resource "infoblox_address" "dynamic" {
  uddi = {
    space   = "ipam/ip_space/00c2c546-6ad1-11f1-88c3-1e3cda826891"
    comment = "dynamically allocated address"
    dynamic_allocation = {
      next_available_id = "ipam/subnet/b7e24b63-752c-11f1-8869-1ae03fbde013"
    }
    tags = {
      Site = "location-2"
    }
  }
}
