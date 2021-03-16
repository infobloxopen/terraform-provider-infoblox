resource "infoblox_network" "subnet1"{
  network_view_name=local.net_view
  tenant_id=local.tenant_id

  network_name="${local.res_prefix}_subnet1"
  allocate_prefix_len = 24
  parent_cidr = local.parent_cidr
  reserve_ip=3

  extensible_attributes = jsonencode({
    TestEA1 = "text3"
    TestEA2 = 7
  })
}

resource "infoblox_network" "subnet2"{
  network_view_name=local.net_view
  tenant_id=local.tenant_id

  network_name="${local.res_prefix}_subnet2"
  allocate_prefix_len = 24
  parent_cidr = local.parent_cidr
  reserve_ip=3

  extensible_attributes = jsonencode({
    Location = "Test loc."
    Site = "Test site"
    TestEA1 = ["text1","text2"]
    TestEA2 = [4,5]
  })
}

resource "infoblox_ip_allocation" "alloc1" {
  network_view_name=local.net_view
  tenant_id=local.tenant_id

  vm_name="${local.res_prefix}_vm1"
  cidr=infoblox_network.subnet1.cidr
}

resource "infoblox_ip_allocation" "alloc2" {
  network_view_name=local.net_view
  tenant_id=local.tenant_id

  vm_name="${local.res_prefix}_vm1"
  cidr=infoblox_network.subnet2.cidr
}
