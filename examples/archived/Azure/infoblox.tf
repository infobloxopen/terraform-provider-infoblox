resource "infoblox_network_view" "demo_network_view"{
  network_view_name="demo1"
  tenant_id="test"

}
resource "infoblox_network" "demo_network"{
  network_view_name="demo1"
  network_name="ex1"
  cidr="10.0.0.0/16"
  tenant_id="test"
  reserve_ip=3
}
resource "infoblox_ip_allocation" "demo_allocation"{
  network_view_name="demo1"
  vm_name="terraform-demo3"
  cidr="${infoblox_network.demo_network.cidr}"
  tenant_id="test"  
}

resource "infoblox_ip_association" "demo_associate"{
  network_view_name="demo1"
  vm_name="${infoblox_ip_allocation.demo_allocation.vm_name}"
  cidr="${infoblox_network.demo_network.cidr}"
  mac_addr ="${azurerm_network_interface.ni.mac_address}"
  ip_addr="${infoblox_ip_allocation.demo_allocation.ip_addr}"
  vm_id ="${azurerm_virtual_machine.vm.id}"
  tenant_id="test"
}

