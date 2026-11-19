// Retrieve Overlapping Blocks filtered by an attribute
data "infoblox_overlapping_block" "get_overlapping_block_using_filters" {
  filters = {
    name = "example_overlapping_block"
  }
}

// Retrieve Overlapping Blocks filtered by tag
data "infoblox_overlapping_block" "get_overlapping_block_using_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all Overlapping Blocks
data "infoblox_overlapping_block" "get_all_overlapping_blocks" {}
