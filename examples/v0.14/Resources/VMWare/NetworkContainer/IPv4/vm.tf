provider "vsphere" {
   allow_unverified_ssl = true
}

data"vsphere_datacenter" "dc"{
 name = "vRA-DC"
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
