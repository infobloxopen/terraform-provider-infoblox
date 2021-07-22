resource "infoblox_network_view" "nv1" {
  tenant_id = local.tenant_id
  network_view=local.net_view
}

resource "infoblox_ipv4_network_container" "v4nc_1" {
  network_view=infoblox_network_view.nv1.network_view
  cidr = "10.0.0.0/16"
  comment = "new network container"

  ext_attrs = jsonencode({
    "Location" = "Test loc."
    "Site" = "Test site"
    "Tenant ID" = local.tenant_id
  })
}

resource "infoblox_ipv4_network" "subnet1"{
  network_view=infoblox_network_view.nv1.network_view
  allocate_prefix_len = 24
  parent_cidr = infoblox_ipv4_network_container.v4nc_1.cidr
  reserve_ip=3

  ext_attrs = jsonencode({
    "Tenant ID" = local.tenant_id
    "Network Name" = "${local.res_prefix}_subnet1"
    "TestEA1" = "text3"
    "TestEA2" = 7
  })
}

resource "infoblox_ipv4_network" "subnet2"{
  network_view=infoblox_network_view.nv1.network_view
  allocate_prefix_len = 24
  parent_cidr = infoblox_ipv4_network_container.v4nc_1.cidr
  reserve_ip=3

  ext_attrs = jsonencode({
    "Tenant ID" = local.tenant_id
    "Network Name" = "${local.res_prefix}_subnet2"
    Location = "Test loc."
    Site = "Test site"
    TestEA1 = ["text1","text2"]
    TestEA2 = [4,5]
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
    "VM Name" = "${local.res_prefix}_vm1"
  })
}

resource "infoblox_ipv4_allocation" "alloc2" {
  network_view=infoblox_network_view.nv1.network_view
  cidr=infoblox_ipv4_network.subnet2.cidr

  #Create Host Record with DNS and DHCP flags
  #dns_view="default"
  #fqdn="testipv4.vmware.com"
  #enable_dns = "false"
  #enable_dhcp = "false"

  ext_attrs = jsonencode({
    "Tenant ID" = local.tenant_id
    "VM Name" = "${local.res_prefix}_vm1"
  })
}
