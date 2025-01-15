resource "infoblox_ipv4_network" "external_ipv4_pool" {
  network_view = local.external_netview_name
  comment = "the pool to allocate external IPv4 addresses"
  cidr = local.external_ipv4_pool
  reserve_ip = local.external_ipv4_pool_reserve_ip
  gateway = local.external_ipv4_pool_gateway
  ext_attrs = jsonencode({
    "Network Name" = "external network pool 1"
    "Location" = "DMZ perimeter"
  })
}

resource "infoblox_ipv6_network" "external_ipv6_pool" {
  network_view = local.external_netview_name
  comment = "the pool to allocate external IPv6 addresses"
  cidr = local.external_ipv6_pool
  reserve_ip = local.external_ipv4_pool_reserve_ip
  gateway = local.external_ipv6_pool_gateway
  ext_attrs = jsonencode({
    "Network Name" = "external network pool 2"
  })
}

resource "infoblox_ip_allocation" "external_ip_allocation_mailserver" {
  network_view = local.external_netview_name
  fqdn = "mail-server.${local.default_external_dns_zone}"
  dns_view = local.external_default_dnsview
  enable_dns = true # It is true by default so may be omitted as well.
  ipv4_addr = local.external_ipv4_pool_mailserver
  ipv6_addr = local.external_ipv6_pool_mailserver
  ttl = 1800 # TTL value for underlying host record on NIOS side
}

resource "infoblox_ip_association" "external_ip_association_mailserver" {
  internal_id = infoblox_ip_allocation.external_ip_allocation_mailserver.internal_id # This is how to set this field correctly.
  enable_dhcp = true # It is false by default so we should enable it explicitly.
  duid = local.external_vm_mail_nic1_duid
  mac_addr = local.external_vm_mail_nic1_mac
}
