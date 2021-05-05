---
layout: "infoblox"
page_title: "Infoblox: infoblox_network_view"
description: |-
  Creates a network view in NIOS.
---

# infoblox\_network\_view

Creates a network view in NIOS.

This resource allows you to create network view in NIOS. When applied the network view will be created.


## Example Usage
Creates a network view and links the network view to a tenant.
```terraform

resource "infoblox_network_view" "demo_network_view" {
  network_view_name = "demo1"
  tenant_id         = "test"
}
```
## Argument Reference

The following arguments are supported:


* `tenant_id` - (Required) Links the network view to a tenant
* `network_view_name` - (Required) Create a network view with a given name
