// creating an alias record with minimum set of parameters
resource "infoblox_alias_record" "alias_record_minimum_params" {
  name        = "alias-record1.test.com"
  target_name = "aa.bb.com"
  target_type = "PTR"
}

// creating an alias record with full set of parameters
resource "infoblox_alias_record" "alias_record_full_params" {
  name        = "alias-record2.test.com"
  target_name = "kk.ll.com"
  target_type = "AAAA"
  comment     = "example alias record"
  dns_view    = "view2"
  disable     = false
  ttl         = 120

  ext_attrs = jsonencode({
    "Site" = "MOROCCO"
  })
}

// creating an alias record with ipv4 reverse mapping zone
resource "infoblox_alias_record" "alias_record_reverse_ipv4" {
  name        = "alias.0.0.192.in-addr.arpa"
  target_name = "test.info.com"
  target_type = "AAAA"
  comment     = "example alias record"
  disable     = false
  ttl         = 120

  ext_attrs = jsonencode({
    "Site" = "MOROCCO"
  })
}
