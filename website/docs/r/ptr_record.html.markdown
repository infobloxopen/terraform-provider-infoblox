---
layout: "infoblox"
page_title: "Infoblox: infoblox_ptr_record"
description: |-
  Creates a PTR record in NIOS.
---


# infoblox\_ptr\_allocation

Creates an PTR record in NIOS .

When applied, a PTR Record will be created in NIOS

## Example Usage

```terraform
resource "infoblox_ptr_record" "demo_record"{
  network_view_name = "demo1"
  vm_name           = "test"
  cidr              = "10.0.0.0/24"
  ip_addr           = "10.0.0.1"
  dns_view          = "default"
  zone              = "$reverse_mappinf_zone"
  tenant_id         = "test"
}
```
## Argument Reference

The following arguments are supported:

* `network_view_name` - (Optional) Unless specified, the providers tries to update IP properties in default network view
* `vm_name` - (Required) A name you want to associate with the IP address.
* `vm_id` - (Optional) Updates the VM id of the vm used to provision
* `cidr` - (Required) The network block in cidr format
* `tenant_id` - (Required) Links the network  to a tenant. For on-premise solutions, this can be any value.
* `dns_view` - (Optional) The view which contains the details of the zone. If not provided , record will be created under default view
* `zone` - (Required) The zone in which you want to update a host record
* `ip_addr` - (Required) - The IP address you want to update in NIOS. Use the Same IP you have passed during IP allocation.

