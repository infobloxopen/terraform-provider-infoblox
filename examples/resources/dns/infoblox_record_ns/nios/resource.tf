resource "infoblox_record_ns" "example_1" {
  nios = {
    name       = "example.com"
    nameserver = "ns1.example.com"
    addresses = [{
      address         = "192.168.1.10"
      auto_create_ptr = false
    }]
    view = "default"
  }
}

resource "infoblox_record_ns" "example_2" {
  nios = {
    name       = "example.com"
    nameserver = "ns2.example.com"
    addresses = [{
      address         = "192.168.1.11"
      auto_create_ptr = true
    }]
    view = "default"
  }
}
