---
layout: "infoblox"
page_title: "Infoblox: infoblox_ip_allocation"
description: |-
  Reserves an IP from a network in NIOS.
---


# infoblox\_ip\_allocation

Reserves an IP from a network in NIOS .

When applied, a next free availabe ip will be Reserved. Additional constraints, such as zone,dns view , mac address can also be configured. The same resource can be used to create Fixed address or Host Record. To create Host record, the zone and dns view parameters need to be specified. 

## Example Usage

```hcl
resource "infoblox_ip_allocation" "demo_allocation"{
  vm_name="terraform-demo1"
  cidr="10.0.0.0/24"
  tenant_id="test"
}
```
## Argument Reference

The following arguments are supported:

* `network_view_name` - (Optional) Unless specified the resource Reserves the IP under default network view
* `vm_name` - (Required) A name you want to associate with the IP address.
* `cidr` - (Required) The network block in cidr format
* `tenant_id` - (Required) Links the network  to a tenant
* `dns_view` - (Optional) The view which contains the details of the zone.If not provided , record will be created under default view
* `zone` - (Optional) The zone in which you want to create a host record
* `enable_dns` - (optional) A boolean value which either creates or not creates for DNS purposes
* `ip_addr` - (Optional) If set , a record will be created in NIOS using a passed IP address value. Takes in a string. If no value is given, a next available IP address will be allocated in NIOS
* `mac_addr` - (Optional) If not set , a reservation will be created in NIOS.

## Additional Note

Dont set the mac address if you are integrating with cloud providers to deploy a Vm and use Infoblox to give the IP address.
