data "unified_next_available_subnets" "by_id" {
  id           = "ipam/address_block/9f9675a2-6ad1-11f1-8248-6ad7b099fb40"
  cidr         = 26
  subnet_count = 5
}

data "unified_next_available_subnets" "by_tags" {
  cidr         = 26
  subnet_count = 5
  tag_filters = {
    environment = "test"
  }
}
