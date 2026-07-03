data "infoblox_next_available_address_blocks" "by_id" {
  id                  = "ipam/address_block/9f9675a2-6ad1-11f1-8248-6ad7b099fb40"
  cidr                = 24
  address_block_count = 3
}

data "infoblox_next_available_address_blocks" "by_tags" {
  cidr                = 24
  address_block_count = 3
  tag_filters = {
    environment = "production"
  }
}
