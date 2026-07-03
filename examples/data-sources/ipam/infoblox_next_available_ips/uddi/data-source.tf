data "infoblox_next_available_ips" "by_id" {
  id         = "ipam/subnet/b7e24b63-752c-11f1-8869-1ae03fbde013"
  ip_count   = 5
  contiguous = false
}

data "infoblox_next_available_ips" "by_tags" {
  ip_count      = 5
  resource_type = "subnet"
  tag_filters = {
    environment = "production"
  }
}
