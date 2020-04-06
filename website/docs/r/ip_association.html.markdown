---
layout: "infoblox"
page_title: "Infoblox: infoblox_ip_association "
description: |-
  Updates the properties of an IP address in NIOS.
---


# infoblox\_ip\_association

Reserves an IP from a network in NIOS .

When applied, the properties of an IP address in NIOS are updated. Additional constraints, such as zone,dns view , mac address can also be configured. The same resource can be used to update Fixed address or Host Record. To update Host record, the zone and dns view parameters need to be specified. 

## Example Usage

```hcl
# The example is given with integration of vmware provider
resource "infoblox_ip_association" "demo_associate"{
  network_view_name="demo1"
  vm_name="test"
  cidr="10.0.0.0/24"
  mac_addr ="${vsphere_virtual_machine.vm.network_interface.0.mac_address}"
  ip_addr="${infoblox_ip_allocation.demo_allocation.ip_addr}"
  vm_id ="${vsphere_virtual_machine.vm.0.id}"
  tenant_id="test"

}

```
## Argument Reference

The following arguments are supported:

* `network_view_name` - (Optional) Unless specified, the providers tries to update IP properties in default network view
* `vm_name` - (Required) A name you want to associate with the IP address.
* `vm_id` - (Required) Updates the VM id of the vm used to provision
* `cidr` - (Required) The network block in cidr format
* `tenant_id` - (Required) Links the network  to a tenant
* `dns_view` - (Optional) The view which contains the details of the zone. If not provided , record will be created under default view
* `zone` - (Optional) The zone in which you want to update a host record
* `ip_addr` - (Required) - The IP address you want to update in NIOS. Use the Same IP you have passed during IP allocation.
* `mac_addr` - (Optional) - Updates the actual mac adress when used with another provider

