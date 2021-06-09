---
layout: "infoblox"
page_title: "Infoblox: infoblox_network"
description: |-
  Creates a network on NIOS.
---

# infoblox\_network

Creates a network on NIOS.

When applied,The network will be created on NIOS and the first IP will be reserved as gateway. Additional constraints such as reserve ip, network name can be configured.


## Example Usage

```hcl
# Protect the master branch of the foo repository. Additionally, require that
# the "ci/travis" context to be passing and only allow the engineers team merge
# to the branch.
resource "infoblox_network" "demo_network"{
  cidr="10.0.0.0/24"
  tenant_id="test"
}
```


## Argument Reference

The following arguments are supported:

* `network_view_name` - (Optional) Unless specified the resource creates network under default network view
* `network_name` - (optional) Unless specified the resource does not associate any name to the network
* `cidr` - (Required) The network block in cidr format
* `tenant_id` - (Required) Links the network  to a tenant
* `reserve_ip` - (optional) reserves the number of Ip's for later use. Takes an `int` value
* `gateway` - (Optional) give the IP you want to reserve for gateway, by default the first IP gets reserved for gateway
* `allocate_prefix_len` - (Optional) Set parameter value>0 to allocate next available network with prefix=value from network container defined by parent_cidr
* `parent_cidr` (Optional) The parent network container block in cidr format to allocate from

## Note

While linking the provider with azure , give `reserve_ip =3` because azure reserves first 4 IP's in it's cloud 
