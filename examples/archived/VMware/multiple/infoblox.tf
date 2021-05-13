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
  count="2"
  network_view_name="demo1"
  vm_name="terraform-demo3-${count.index + 1}"
  cidr="${infoblox_network.demo_network.cidr}"
  tenant_id="test"  
}

resource "infoblox_ip_association" "demo_associate"{
  count="2" 
  network_view_name="demo1"
  vm_name="${element(infoblox_ip_allocation.demo_allocation.*.vm_name,count.index)}" 
  cidr="${infoblox_network.demo_network.cidr}"
  mac_addr ="${element(vsphere_virtual_machine.vm.*.network_interface.0.mac_address, count.index)}"
  ip_addr="${infoblox_ip_allocation.demo_allocation.*.ip_addr[count.index]}"
  vm_id ="${element(vsphere_virtual_machine.vm.*.id,count.index)}"
  tenant_id="test"
}
resource "infoblox_a_record" "demo_record"{
  count="2"
  network_view_name="demo1"
  cidr="${infoblox_network.demo_network.cidr}"
  ip_addr="${infoblox_ip_allocation.demo_allocation.*.ip_addr[count.index]}"
  vm_name="${element(infoblox_ip_allocation.demo_allocation.*.vm_name,count.index)}"
  dns_view="default.demo1"
  zone="aa.com"
  tenant_id="test"
 vm_id ="${element(vsphere_virtual_machine.vm.*.id,count.index)}"
}

