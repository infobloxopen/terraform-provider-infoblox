resource "infoblox_network_view" "demo_network_view"{
  network_view_name="demo1"
  tenant_id="test"

}
resource "infoblox_network" "demo_network"{
  network_view_name="demo1"
  network_name="ex1"
  cidr="10.10.20.0/24"
  tenant_id="test"
}

resource "infoblox_ip_allocation" "demo_allocation"{
  network_view_name="demo1"
  host_name="terraform-demo3"
  cidr="${infoblox_network.demo_network.cidr}"
  tenant_id="test"  
}

resource "infoblox_ip_association" "demo_associate"{
  network_view_name="demo1"
  host_name="${infoblox_ip_allocation.demo_allocation.host_name}"
  cidr="${infoblox_network.demo_network.cidr}"
  mac_addr ="${vsphere_virtual_machine.vm.network_interface.0.mac_address}"
  ip_addr="${infoblox_ip_allocation.demo_allocation.ip_addr}"
  vm_id ="test_id" 
  tenant_id="test"
}
