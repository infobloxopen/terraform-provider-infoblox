---
layout: "infoblox"
page_title: "Provider: Infoblox"
description: |-
  The Infoblox provider is used to interact with Infoblox Grid objects.
---

# Infoblox Provider

The Infoblox provider is used to interact with Infoblox Grid objects.

## Authentication

The provider allows you to manage Infoblox credentials. The following methods for supplying credentials are supported:
- Static credentials
- Environmental variables

### Static Credentials

Static credentials can be provided in the Infoblox provider block. A typical provider configuration will look like this:

```hcl
provider "infoblox"{
  version="~> 1.0"
  username="infoblox_user"
  password="infoblox"
  server="10.0.0.1"
}
```

#### Argument Reference

The following arguments are supported in the `provider` block:

* `version` - (Optional) Specify the provider version
* `username` - (Required) User login details
* `password` - (Required) Infoblox user password
* `server` - (Required) Grid master's or CP members IP

### Environmental Variables

Credentials can be provided using the `INFOBLOX_USENAME`, `INFOBLOX_PASSWORD`, and `INFOBLOX_SERVER` environmental variables which correspond to your username, password, and server respectively.

Example Usage:

```sh
$ export INFOBLOX_USERNAME="infoblox_user"
$ export INFOBLOX_PASSWORD="infoblox"
$ export INFOBLOX_SERVER="10.0.0.1"
```

## Supported Functionality

* The provider only allows creation of network views. Deletion of network views is not supported.
* The provider supports only Create , Read and Delete for networks/CIDRs . Updating a network is not supported.
* If the provider is used to allocate IPs to VMs using other providers, please use the 2 resource blocks `ip_allocation` and `ip_association`. [examples](https://github.com/terraform-providers/terraform-provider-infoblox/tree/master/examples) for using the Infoblox provider are shown.
* Using the `ip_allocation` block , you can create either a Reservation, Fixed address, or Host Record. To create a host record please look at the `ip_allocation` resource documentation for detailed instructions.
* If the provider is not used with any other providers, just use the `ip_allocation` block to allocate IPs. `ip_allocation` supports complete CRUD operations.
* `ip_association` block is used to update the properties of VMs. If you are not using the provider with other providers to deploy VMs and allocate IPs from NIOS, ignore this block.
* The provider supports create, Read and Delete for A,PTR,CNAME Records. Update functionality is not supported.

## Additional Note

The provider is designed keeping in mind the cloud network automation aspects of NIOS. If you don't have a cloud license installed in NIOS please add the below EAs manually.

In the Grid Manager, go to Administration > Extensible Attributes and add the following EA's:

* `VM Name` as string 
* `VM ID` as string
* `Cloud API Owned` as List(Values : True, False)
* `CMP Type` as string
* `Tenant ID` as string
* `Network Name` as string
