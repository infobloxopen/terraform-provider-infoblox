---
layout: "infoblox"
page_title: "Infoblox: infoblox_network"
description: |-
  Fetches information on Network ID and Name from NIOS.
---


# infoblox\_network

Fetches information on Network ID(_ref) and Name from NIOS.

When applied, data such as _ref and Name will be returned.

## Example Usage

```terraform

resource "infoblox_network" "test" {
  network_name = "test"
  cidr         = "10.0.23.0/24"
  reserve_ip   = 2
  tenant_id    = "default"
}

#Example usage of Network Datasource
data "infoblox_network" "test" {
  cidr      = infoblox_network.test.cidr #add a CIDR for which the data is to be fetched
  tenant_id = "default"
}
```
## Argument Reference

The following arguments are supported:

* `network_view_name` - (Optional) Unless specified, the providers considers default network view.
* `network_name` - (Computed) A name that is fetched from the datasource.
* `cidr` - (Required) The network block in cidr format.
* `tenant_id` - (Required) The tenant in which the network exists.
