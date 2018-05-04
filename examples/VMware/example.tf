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

data"vsphere_datacenter" "dc"{
 name = "Blr-Devlab"
}


data "vsphere_datastore" "datastore" {
  name          = "datastore_44"
  datacenter_id = "${data.vsphere_datacenter.dc.id}"
}

data "vsphere_resource_pool" "pool" {
  name          = "Blr-Cloud/Resources"
  datacenter_id = "${data.vsphere_datacenter.dc.id}"

}

data "vsphere_network" "network" {
  name          = "VM Network"
  datacenter_id = "${data.vsphere_datacenter.dc.id}"
}

data "vsphere_virtual_machine" "template" {
  name          = "WebTinyCentOS65x86-tcpdump"
  datacenter_id = "${data.vsphere_datacenter.dc.id}"
}

resource "vsphere_virtual_machine" "vm" {
  name             = "${infoblox_ip_allocation.demo_allocation.host_name}"
  resource_pool_id = "${data.vsphere_resource_pool.pool.id}"
  datastore_id     = "${data.vsphere_datastore.datastore.id}"
  num_cpus = 2
  memory   = 1024
  guest_id = "${data.vsphere_virtual_machine.template.guest_id}"
  scsi_type = "${data.vsphere_virtual_machine.template.scsi_type}"

  network_interface {
    network_id   = "${data.vsphere_network.network.id}"
    adapter_type = "${data.vsphere_virtual_machine.template.network_interface_types[0]}"
  }

  disk {
    label            = "disk0"
    size             = "${data.vsphere_virtual_machine.template.disks.0.size}"
    eagerly_scrub    = "${data.vsphere_virtual_machine.template.disks.0.eagerly_scrub}"
    thin_provisioned = "${data.vsphere_virtual_machine.template.disks.0.thin_provisioned}"
  }

  clone {
    template_uuid = "${data.vsphere_virtual_machine.template.id}"

    customize {
      linux_options {
        host_name = "terraform-test1"
        domain    = "test.internal"
      }

      #This is where we are injecting/allocating ip

      network_interface {
        ipv4_address = "${infoblox_ip_allocation.demo_allocation.ip_addr}"
        ipv4_netmask = 24
      }

      ipv4_gateway = "10.10.100.1"
    }
  }
}


output "mac_op" "op"{
value="${vsphere_virtual_machine.vm.network_interface.0.mac_address}"
}
  #we are updating mac address properties here

resource "infoblox_ip_association" "demo_associate"{
  network_view_name="demo1"
  host_name="${infoblox_ip_allocation.demo_allocation.host_name}"
  cidr="${infoblox_network.demo_network.cidr}"
  mac_addr ="${vsphere_virtual_machine.vm.network_interface.0.mac_address}"
  ip_addr="${infoblox_ip_allocation.demo_allocation.ip_addr}"
  vm_id ="test_id" 
  tenant_id="test"
}
