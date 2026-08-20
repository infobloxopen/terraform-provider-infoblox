// Create Extensible Attribute Definition with Basic Fields
resource "infoblox_extensibleattributedef" "create_ea_basic" {
  nios = {
    name = "example_ea_1"
    type = "STRING"
  }
}

// Create Extensible Attribute Definition with Additional Fields
resource "infoblox_extensibleattributedef" "create_ea_additional_fields" {
  nios = {
    name                 = "example_ea_2"
    type                 = "INTEGER"
    min                  = 1
    max                  = 4094
    default_value        = "1"
    comment              = "Extensible Attribute Definition created by Terraform"
    flags                = "I"
    allowed_object_types = ["Network", "IPv6Network", "NetworkContainer", "IPv6NetworkContainer"]
  }
}
