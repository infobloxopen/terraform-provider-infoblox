// Create NAT Group with Basic Fields
resource "infoblox_nat_group" "natgroup_basic_fields" {
  nios = {
    name = "natgroup-basic"
  }
}

// Create NAT Group with Additional Fields
resource "infoblox_nat_group" "natgroup_with_additional_config" {
  nios = {
    name    = "natgroup-example"
    comment = "Example NAT Group for Grid communication"
  }
}
