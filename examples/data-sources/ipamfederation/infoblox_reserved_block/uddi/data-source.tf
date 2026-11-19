// Retrieve Reserved Blocks filtered by an attribute
data "infoblox_reserved_block" "get_reserved_block_using_filters" {
  filters = {
    comment = "Example Reserved Block created through Terraform"
  }
}

// Retrieve Reserved Blocks filtered by tag
data "infoblox_reserved_block" "get_reserved_block_using_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all Reserved Blocks
data "infoblox_reserved_block" "get_all_reserved_blocks" {}
