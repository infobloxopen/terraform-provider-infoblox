// Create NAT Group with Basic Fields
resource "infoblox_natgroup" "natgroup_basic_fields" {
  nios = {
    name = "natgroup-basic"
  }
}

// Create NAT Group with Additional Fields
resource "infoblox_natgroup" "natgroup_with_additional_config" {
  nios = {
    name    = "natgroup-example"
    comment = "Example NAT Group for Grid communication"
  }
}
