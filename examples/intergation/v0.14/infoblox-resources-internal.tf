resource "infoblox_ipv4_network" "internal_ipv4_pool" {
  network_view = local.internal_netview_name
  comment = "the pool to allocate internal IPv4 addresses"
  cidr = local.internal_ipv4_pool
  reserve_ip = local.internal_ipv4_pool_reserve_ip
  gateway = local.internal_ipv4_pool_gateway
  ext_attrs = jsonencode({
    "Network Name" = "internal network pool 1"
  })
}

resource "infoblox_ipv6_network" "internal_ipv6_pool" {
  network_view = local.internal_netview_name
  comment = "the pool to allocate internal IPv6 addresses"
  cidr = local.internal_ipv6_pool
  reserve_ip = local.internal_ipv4_pool_reserve_ip
  gateway = local.internal_ipv6_pool_gateway
  ext_attrs = jsonencode({
    "Network Name" = "internal network pool 2"
  })
}

resource "infoblox_ip_allocation" "internal_ip_allocation_reception_1" {
  network_view = local.internal_netview_name
  fqdn = "reception-1.${local.default_internal_dns_zone}"
  dns_view = local.internal_default_dnsview
  enable_dns = true # It is true by default so may be omitted as well.
  ipv4_cidr = infoblox_ipv4_network.internal_ipv4_pool.cidr
  ipv6_cidr = infoblox_ipv6_network.internal_ipv6_pool.cidr
}

resource "infoblox_ip_association" "internal_ip_association_reception_1" {
  internal_id = infoblox_ip_allocation.internal_ip_allocation_reception_1.internal_id # This is how to set this field correctly.
  enable_dhcp = true # It is false by default so we should enable it explicitly.
  duid = local.internal_reception_1_nic1_duid
  mac_addr = local.internal_reception_1_nic1_mac
}

resource "infoblox_ip_allocation" "internal_ip_allocation_reception_2" {
  network_view = local.internal_netview_name
  fqdn = "reception-2.${local.default_internal_dns_zone}"
  dns_view = local.internal_default_dnsview
  enable_dns = true # It is true by default so may be omitted as well.
  ipv4_cidr = infoblox_ipv4_network.internal_ipv4_pool.cidr
  ipv6_cidr = infoblox_ipv6_network.internal_ipv6_pool.cidr
}

resource "infoblox_ip_association" "internal_ip_association_reception_2" {
  internal_id = infoblox_ip_allocation.internal_ip_allocation_reception_2.internal_id # This is how to set this field correctly.
  enable_dhcp = true # It is false by default so we should enable it explicitly.
  duid = local.internal_reception_2_nic1_duid
  mac_addr = local.internal_reception_2_nic1_mac
}

resource "infoblox_ip_allocation" "internal_ip_allocation_user_1" {
  network_view = local.internal_netview_name
  fqdn = "user-1.${local.default_internal_dns_zone}"
  dns_view = local.internal_default_dnsview
  enable_dns = true # It is true by default so may be omitted as well.
  ipv4_cidr = infoblox_ipv4_network.internal_ipv4_pool.cidr
  ipv6_cidr = infoblox_ipv6_network.internal_ipv6_pool.cidr
}

resource "infoblox_ip_association" "internal_ip_association_user_1" {
  internal_id = infoblox_ip_allocation.internal_ip_allocation_user_1.internal_id # This is how to set this field correctly.
  enable_dhcp = true # It is false by default so we should enable it explicitly.
  duid = local.internal_user_1_nic1_duid
  mac_addr = local.internal_user_1_nic1_mac
}

resource "infoblox_ip_allocation" "internal_ip_allocation_user_2" {
  network_view = local.internal_netview_name
  fqdn = "user-2.${local.default_internal_dns_zone}"
  dns_view = local.internal_default_dnsview
  enable_dns = true # It is true by default so may be omitted as well.
  ipv4_cidr = infoblox_ipv4_network.internal_ipv4_pool.cidr
  ipv6_cidr = infoblox_ipv6_network.internal_ipv6_pool.cidr
}

resource "infoblox_ip_association" "internal_ip_association_user_2" {
  internal_id = infoblox_ip_allocation.internal_ip_allocation_user_2.internal_id # This is how to set this field correctly.
  enable_dhcp = true # It is false by default so we should enable it explicitly.
  duid = local.internal_user_2_nic1_duid
  mac_addr = local.internal_user_2_nic1_mac
}

resource "infoblox_a_record" "websrv_addr_rec" {
  dns_view = local.internal_default_dnsview
  fqdn = local.internal_web_server
  ip_addr = local.internal_web_server_ipv4_addr
}
