// List specific Reserved Blocks using filters
list "infoblox_reserved_block" "list_reserved_block_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "Example Reserved Block created through Terraform"
    }
  }
  limit = 10
}

// List specific Reserved Blocks using Tags
list "infoblox_reserved_block" "list_reserved_block_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Reserved Blocks with resource details included
list "infoblox_reserved_block" "list_reserved_block_with_resource" {
  provider         = infoblox
  include_resource = true
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}
