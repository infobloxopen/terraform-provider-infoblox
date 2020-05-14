---
layout: "infoblox"
page_title: "Provider: Infoblox"
description: |-
  The Infoblox provider is used to interact with Infoblox organization resources.
---

# Infoblox Provider

The Infoblox provider is used to interact with Infoblox organization resources.

The provider allows you to manage infoblox credentials.

A typical provider configuration will look something like:

```hcl
provider "infoblox"{
INFOBLOX_USERNAME=infoblox
INFOBLOX_SERVER=10.0.0.1
INFOBLOX_PASSWORD=infoblox
}
```

## Argument Reference

The following arguments are supported in the `provider` block:
* `INFOBLOX_USERNAME` - (Required) User login details
* `INFOBLOX_SERVER` - (Required) Grid master's or CP members IP
* `INFOBLOX_PASSWORD` - (Required) Grid master's or CP members password

## Supported Functionality

* The provider only allows to create network views. Deletion of network view's is not supported
* The provider supports only Create , Read and delete for networks/CIDR's . Updating a network is not supported
* If the provider is used to allocate IP's to VM's using other provider, please use the 2 resource's `ip_allocation` and `ip_association` blocks. [examples](https://github.com/infobloxopen/terraform-provider-infoblox/tree/master/examples) for using Infoblox provider are shown.
* Using the `ip_allocation` block , you can either create a Reservation/Fixed address/Host Record. To create host record please look at the ip_allocation resource as to how to create a Host record
* If the provider is not used with any other provider's, just use the `ip_allocation` block. `ip_allocation` supports complete CRUD operations
* `ip_association` block is used to update the properties of VM's , If you are not using the provider with other providers to deploy VM and allocate IP from NIOS, ignore this block
* The provider supports create, Read and delete for A,PTR,CNAME Records. Update functionality is not supported.

## Additional Note

The provider is designed to keeping in mind the cloud aspective of NIOS. So if you don't have a cloud license installed in NIOS please add the below EA's in NIOS manually
Go to Administration > Extensible Attributes and the following EA's 
* `VM Name` as string 
* `VM ID` as string
* `Cloud API Owned` as List(Vlaues : True,False)
* `CMP Type` as string
* `Tenant ID` as string
* `Network Name` as string
