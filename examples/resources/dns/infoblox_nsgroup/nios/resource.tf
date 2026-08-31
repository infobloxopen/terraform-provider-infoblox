// Create NS group with Basic Fields
resource "infoblox_nsgroup" "create_ns_group" {
  nios = {
    name = "example_ns_group"
    grid_primary = [
      {
        name = "infoblox.localdomain"
      }
    ]
  }
}

// Create NS Group with Additional Fields
resource "infoblox_nsgroup" "create_ns_group_with_additional_fields" {
  nios = {
    name    = "example_ns_group_1"
    comment = "Example NS Group"

    grid_secondaries = [
      {
        name = "infoblox.localdomain",
      },
    ]
    external_primaries = [
      {
        name    = "ns1.example.com",
        address = "2.3.4.5",
      },
    ]
    use_external_primary = true
    ext_attrs = {
      Site = "location-1"
    }
  }
}
