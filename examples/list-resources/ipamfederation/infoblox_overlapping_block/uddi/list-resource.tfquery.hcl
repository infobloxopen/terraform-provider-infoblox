// List specific Overlapping Blocks using filters
list "infoblox_overlapping_block" "list_overlapping_block_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "Created by Terraform"
    }
  }
  limit = 10
}

// List specific Overlapping Blocks using Tags
list "infoblox_overlapping_block" "list_overlapping_block_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Overlapping Blocks with resource details included
list "infoblox_overlapping_block" "list_overlapping_block_with_resource" {
  provider         = infoblox
  include_resource = true
}
