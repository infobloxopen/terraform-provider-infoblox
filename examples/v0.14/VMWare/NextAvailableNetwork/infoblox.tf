# Allocate a network in Infoblox Grid under provided parent CIDR
resource "infoblox_network" "demo_network"{
  network_view_name = "default"
  network_name = "test_network"
  tenant_id = "test"
  allocate_prefix_len = 24
  parent_cidr = "10.0.0.0/16"
  reserve_ip = 2  
}

# Allocate IP from network
resource "infoblox_ip_allocation" "demo_allocation"{
  network_view_name = "default"
  vm_name = "terraform-demo"
  cidr = infoblox_network.demo_network.cidr
  tenant_id = "test"  
}

# Update Grid with VM data
resource "infoblox_ip_association" "demo_associate"{
  network_view_name = "default"
  vm_name = infoblox_ip_allocation.demo_allocation.vm_name
  cidr = infoblox_network.demo_network.cidr
  mac_addr = vsphere_virtual_machine.vm.network_interface.0.mac_address
  ip_addr = infoblox_ip_allocation.demo_allocation.ip_addr
  vm_id = vsphere_virtual_machine.vm.id
  tenant_id ="test"
}


