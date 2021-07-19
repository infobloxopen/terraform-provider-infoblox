resource "infoblox_network_view" "nv1" {
  tenant_id = local.tenant_id
  network_view=local.net_view
}

resource "infoblox_ipv4_network_container" "nc_1" {
  network_view=infoblox_network_view.nv1.network_view
  cidr = "10.0.0.0/16"
  comment = "new network container"

  ext_attrs = jsonencode({
    "Tenant ID" = local.tenant_id
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}

resource "infoblox_ipv4_network" "subnet1"{
  network_view=infoblox_network_view.nv1.network_view
  allocate_prefix_len = 24
  parent_cidr = infoblox_ipv4_network_container.nc_1.cidr
  reserve_ip=3

  ext_attrs = jsonencode({
    "Tenant ID" = local.tenant_id
    "Network Name" = "${local.res_prefix}_subnet1"
  })
}

resource "infoblox_ipv4_allocation" "alloc1" {
  network_view=infoblox_network_view.nv1.network_view
  cidr=infoblox_ipv4_network.subnet1.cidr

  #Create Host Record with DNS and DHCP flags
  dns_view="default"
  fqdn="testipv4.vmware.com"
  enable_dns = "false"
  enable_dhcp = "false"

  ext_attrs = jsonencode({
    "Tenant ID" = local.tenant_id
  })
}

resource "infoblox_ipv4_association" "assoc1"{
  network_view=infoblox_network_view.nv1.network_view
  cidr = infoblox_ipv4_allocation.alloc1.cidr
  mac_addr = local.vm_mac_addr
  ip_addr = infoblox_ipv4_allocation.alloc1.ip_addr

  #Create Host Record with DNS and DHCP flags
  dns_view="default"
  fqdn="testipv4.vmware.com"
  enable_dns = "false"
  enable_dhcp = "false"  

  ext_attrs = jsonencode({
    "Tenant ID" = local.tenant_id
    "VM Name" = "${local.res_prefix}_vm1"
    "VM ID" = local.vm_id
  })
}

resource "infoblox_ipv6_network_container" "nc_2" {
  network_view=infoblox_network_view.nv1.network_view
  cidr = "fc00::/56"
  comment = "new network container"

  ext_attrs = jsonencode({
    "Tenant ID" = local.tenant_id
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}

resource "infoblox_ipv6_network" "subnet2"{
  network_view=infoblox_network_view.nv1.network_view
  allocate_prefix_len = 64
  parent_cidr = infoblox_ipv6_network_container.nc_2.cidr
  reserve_ip=10


  ext_attrs = jsonencode({
    "Tenant ID" = local.tenant_id
    "Network Name" = "${local.res_prefix}_subnet2"
  })
}

locals {
  ipv6_reserved_ips = 10
}

resource "infoblox_ipv6_allocation" "alloc_reserved" {
  count=local.ipv6_reserved_ips

  network_view=infoblox_network_view.nv1.network_view
  cidr=infoblox_ipv6_network.subnet2.cidr
  duid = format("00:%.2x", count.index)

  #Create Host Record with DNS and DHCP flags
  #dns_view="default"
  #fqdn="testipv6.vmware.com"
  #enable_dns = "false"
  #enable_dhcp = "false"

  ext_attrs = jsonencode({
    "Tenant ID" = local.tenant_id
  })
}

resource "infoblox_ipv6_allocation" "alloc2" {
  network_view=infoblox_network_view.nv1.network_view
  cidr=infoblox_ipv6_network.subnet2.cidr
  duid = format("00:%.2x", local.ipv6_reserved_ips)

  #Create Host Record with DNS and DHCP flags
  #dns_view="default"
  #fqdn="testipv6.vmware.com"
  #enable_dns = "false"
  #enable_dhcp = "false"

  ext_attrs = jsonencode({
    "Tenant ID" = local.tenant_id
  })

  depends_on = [infoblox_ipv6_allocation.alloc_reserved]
}

resource "infoblox_ipv6_association" "assoc2"{
  network_view=infoblox_network_view.nv1.network_view
  cidr = infoblox_ipv6_allocation.alloc2.cidr
  mac_addr = local.vm_mac_addr
  ip_addr = infoblox_ipv6_allocation.alloc2.ip_addr

  #Create Host Record with DNS and DHCP flags
  #dns_view="default"
  #fqdn="testipv6.vmware.com"
  #enable_dns = "false"
  #enable_dhcp = "false"
  
  ext_attrs = jsonencode({
    "Tenant ID" = local.tenant_id
    "VM Name" = "${local.res_prefix}_vm1"
    "VM ID" = local.vm_id
  })
}
