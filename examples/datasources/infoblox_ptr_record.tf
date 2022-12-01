// searching by an IP address
data "infoblox_ptr_record" "ptr1" {
  ptrdname = "rec1.example1.org"
  ip_addr = "10.0.0.1"
}

// searching by a record's name
data "infoblox_ptr_record" "ptr2" {
  ptrdname = "rec2.example1.org"
  record_name = "2.0.0.10.in-addr.arpa"
}

// non-default DNS view name
data "infoblox_ptr_record" "ptr3" {
  ptrdname = "rec3.example2.org"
  dns_view = "nondefault_dnsview1"
  ip_addr = "2002:1f93::3"
}
