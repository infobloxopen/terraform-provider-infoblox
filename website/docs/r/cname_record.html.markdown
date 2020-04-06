---
layout: "infoblox"
page_title: "Infoblox: infoblox_cname_record"
description: |-
  Creates a cname record in NIOS.
---


# infoblox\_ptr\_allocation

Creates an cname record in NIOS .

When applied, cname Record will be created in NIOS

## Example Usage

```hcl
resource "infoblox_cname_record" "demo_cname"{

  canonical="demo"
  zone="aa.com"
  alias="demo1"
tenant_id="test"
}

```
## Argument Reference

The following arguments are supported:

* `canonical` - (Required) A name you want to associate with the IP address.
* `vm_id` - (Optional) Updates the VM id of the vm used to provision
* `tenant_id` - (Required) Links the network  to a tenant
* `dns_view` - (Optional) The view which contains the details of the zone. If not provided , record will be created under default view
* `zone` - (Required) The zone in which you want to update a host record
* `alias`- (Required) Alias for you cname record
