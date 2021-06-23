provider "vsphere" {
   allow_unverified_ssl = true
}

data "vsphere_datacenter" "dc"{
  name = "Blr-Devlab"
}

data "vsphere_datastore" "datastore" {
  name          = "datastore_44"
  datacenter_id = data.vsphere_datacenter.dc.id
}

data "vsphere_resource_pool" "pool" {
  name          = "Blr-Cloud/Resources"
  datacenter_id = data.vsphere_datacenter.dc.id

}

data "vsphere_network" "network" {
  name          = "VM Network"
  datacenter_id = data.vsphere_datacenter.dc.id
}

data "vsphere_virtual_machine" "template" {
  name          = "WebTinyCentOS65x86-tcpdump"
  datacenter_id = data.vsphere_datacenter.dc.id
}

# Create VM in vsphere
resource "vsphere_virtual_machine" "vm_ipv4" {
  name             = infoblox_ipv4_allocation.ipv4_allocation.host_name
  resource_pool_id = data.vsphere_resource_pool.pool.id
  datastore_id     = data.vsphere_datastore.datastore.id
  num_cpus = 2
  memory   = 1024
  guest_id = data.vsphere_virtual_machine.template.guest_id
  scsi_type = data.vsphere_virtual_machine.template.scsi_type

  network_interface {
    network_id   = data.vsphere_network.network.id
    adapter_type = data.vsphere_virtual_machine.template.network_interface_types[0]
  }

  disk {
    label            = "disk0"
    size             = data.vsphere_virtual_machine.template.disks.0.size
    eagerly_scrub    = data.vsphere_virtual_machine.template.disks.0.eagerly_scrub
    thin_provisioned = data.vsphere_virtual_machine.template.disks.0.thin_provisioned
  }

  clone {
    template_uuid = data.vsphere_virtual_machine.template.id

    customize {
      linux_options {
        host_name = "terraform-test1"
        domain    = "test.internal"
      }

      network_interface {
        ipv4_address = infoblox_ipv4_allocation.ipv4_allocation.ip_addr
        ipv4_netmask = 24
      }

      ipv4_gateway = infoblox_ipv4_network.ipv4_network.gateway
    }
  }
}

resource "vsphere_virtual_machine" "vm_ipv6" {
  name             = infoblox_ipv6_allocation.ipv6_allocation.host_name
  resource_pool_id = data.vsphere_resource_pool.pool.id
  datastore_id     = data.vsphere_datastore.datastore.id
  wait_for_guest_net_timeout = 0
  #wait_for_guest_ip_timeout  = 5
  num_cpus = 2
  memory   = 1024
  guest_id = data.vsphere_virtual_machine.template.guest_id
  scsi_type = data.vsphere_virtual_machine.template.scsi_type

  network_interface {
    network_id   = data.vsphere_network.network.id
    adapter_type = data.vsphere_virtual_machine.template.network_interface_types[0]
  }

  disk {
    label            = "disk0"
    size             = data.vsphere_virtual_machine.template.disks.0.size
    eagerly_scrub    = data.vsphere_virtual_machine.template.disks.0.eagerly_scrub
    thin_provisioned = data.vsphere_virtual_machine.template.disks.0.thin_provisioned
  }

  clone {
    template_uuid = data.vsphere_virtual_machine.template.id

    customize {
      linux_options {
        host_name = "terraform-test1"
        domain    = "test.internal"
      }

      network_interface {
        ipv6_address = infoblox_ipv6_allocation.ipv6_allocation.ip_addr
        ipv6_netmask = 64
      }

      ipv6_gateway = infoblox_ipv6_network.ipv6_network.gateway
    }
  }
}
