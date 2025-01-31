# IP address allocation, minimal set of parameters
resource "infoblox_ip_allocation" "allocation1" {
  dns_view = "default" # this time the parameter is not optional, because...
  # enable_dns = true # ... this is 'true', by default
  fqdn      = "host1.example1.org" # this resource type is implemented as a host record on NIOS side
  ipv4_addr = "1.2.3.4"
}

# Wide set of parameters
resource "infoblox_ip_allocation" "allocation2" {
  dns_view  = "nondefault_dnsview1"
  fqdn      = "host2.example2.org"
  ipv6_addr = "2002:1f93::1234"
  ttl       = 120
  comment   = "another host record, statically allocated"
  ext_attrs = jsonencode({
    "Tenant ID" = "tenant_3261798"
  })
}

# IPv4 and IPv6 at the same time
resource "infoblox_ip_allocation" "allocation3" {
  dns_view     = "nondefault_dnsview2"
  network_view = "nondefault_netview" # we have to mention the exact network view which the DNS view belongs to
  fqdn         = "host3.example4.org"
  ipv6_addr    = "2002:1f93::12:40"
  ipv4_addr    = "1.2.3.40"
  ttl          = 0
  comment      = "still statically allocated, but IPv4 and IPv6 together"
  ext_attrs = jsonencode({
    "Tenant ID" = "tenant_3261798"
  })
}

resource "infoblox_ip_allocation" "allocation4" {
  enable_dns = false
  # dns_view = "nondefault_dnsview2" # dns_view makes no sense when enable_dns = false and
  # you MUST remove it or comment out
  network_view = "nondefault_netview" # we have to mention the exact network view which the DNS view belongs to
  fqdn         = "host4"              # just one name component, because there is no host-to-DNS-zone assignment

  # Either of the IP addresses must belong to an existing
  # network (not a network container) in the GIVEN network view,
  # because of enable_dns = false
  ipv6_addr = "2002:1f93::12:40"
  ipv4_addr = "10.1.0.60"

  # We must ensure that appropriate network exists by the time this resource is being created
  depends_on = [infoblox_ipv4_network.net1]
}

# Dynamic allocation
resource "infoblox_ip_allocation" "allocation5" {
  dns_view     = "nondefault_dnsview2"
  network_view = "nondefault_netview"
  fqdn         = "host5.example4.org"
  ipv6_cidr    = infoblox_ipv6_network.net2.cidr
  ipv4_cidr    = infoblox_ipv4_network.net2.cidr
}

# Dynamic allocation of IPV6 Host record using next-available
resource "infoblox_ip_allocation" "allocation6" {
  fqdn = "host5.test.com"
  filter_params = jsonencode({
    "*Site" : "Turkey"
  })
  ip_address_type = "IPV6"
  enable_dns      = true
  ttl             = 60
}

# Dynamic allocation of IPV4 Host record using next-available
resource "infoblox_ip_allocation" "allocation7" {
  fqdn = "host6.test.com"
  filter_params = jsonencode({
    "*Site" : "Turkey"
  })
  ip_address_type = "IPV4"
  enable_dns      = true
  ttl             = 60
}

# Dynamic allocation of both IPV4 and IPV6 Host record using next-available
resource "infoblox_ip_allocation" "allocation8" {
  fqdn = "host7.test.com"
  filter_params = jsonencode({
    "*Site" : "Turkey"
  })
  ip_address_type = "Both"
  enable_dns      = true
  ttl             = 60
}

# Static allocation of both IPV4 and IPV6 Host record with aliases
resource "infoblox_ip_allocation" "allocation9" {
  network_view = "custom"
  enable_dns   = true
  fqdn         = "host8.test.com"
  aliases      = ["alias1"]
  ipv6_addr    = "2002:93::1234"
  ipv4_addr    = "10.1.1.0"
  ttl          = 120
  comment      = "another host record, statically allocated"
  ext_attrs = jsonencode({
    "Tenant ID" = "tenant_3261798"
  })
}
