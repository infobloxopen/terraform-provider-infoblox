resource "infoblox_a_record" "a_record_static1"{
  fqdn = "static1.test.com" # the zone 'test.com' MUST exist in the DNS view
  ip_addr = "192.168.31.31"
  ttl = 10
  comment = "static A-record #1"
  ext_attrs = jsonencode({
    "Location" = "New York"
    "Site" = "HQ"
  })
}

resource "infoblox_a_record" "a_record_static2"{
  fqdn = "static2.test.com"
  ip_addr = "192.168.31.32"
  ttl = 0 # ttl=0 means 'do not cache'
  dns_view = "non_default_dnsview" # corresponding DNS view MUST exist
}

resource "infoblox_a_record" "a_record_dynamic1"{
  fqdn = "dynamic1.test.com"
  # ip_addr = "192.168.31.32" # commented out, CIDR is used for dynamic allocation
  # ttl = 0 # not mentioning TTL value means using parent's zone TTL value.

  # In case of non-default network view,
  # you should specify DNS view as well.
  network_view = "non_default" # corresponding network view MUST exist
  dns_view = "nondefault_view" # corresponding DNS view MUST exist
  cidr = "10.20.30.0/24" # appropriate network or network container MUST exist in the network view
}

resource "infoblox_a_record" "a_record_dynamic2"{
  fqdn = "dynamic2.test.com"
  # network_view = "default" # not specifying explicitly means using the default network view
  # dns_view = "default" # not specifying explicitly means using the default DNS view
  cidr = "10.20.30.0/24" # appropriate network or network container MUST exist in the network view
}
