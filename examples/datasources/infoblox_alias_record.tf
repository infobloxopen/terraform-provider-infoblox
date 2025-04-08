// creating an alias record resource
resource "infoblox_alias_record" "alias_record" {
  name        = "alias-record.test.com"
  target_name = "hh.ll.com"
  target_type = "NAPTR"
  comment     = "example alias record"
  dns_view    = "default"
  disable     = true
  ttl         = 1200
  ext_attrs = jsonencode({
    "Site" = "Ireland"
  })
}

// accessing alias record by specifying name and comment
data "infoblox_alias_record" "alias_read" {
  filters = {
    name    = infoblox_alias_record.alias_record.name
    comment = infoblox_alias_record.alias_record.comment
  }
}

// returns matching alias record with name and comment, if any
output "alias_record_res" {
  value = data.infoblox_alias_record.alias_read
}

// accessing alias record by specifying dns_view, zone, target_name and target_type
data "infoblox_alias_record" "alias_read1" {
  filters = {
    view        = "default"
    zone        = "test.com"
    target_name = "hh.ll.com"
    target_type = "NAPTR"
  }
}

// returns matching alias record with dns_view, zone, target_name and target_type, if any
output "alias_record_res1" {
  value = data.infoblox_alias_record.alias_read1
}