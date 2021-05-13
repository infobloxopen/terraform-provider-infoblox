 <img width="171" alt="capture" src="https://user-images.githubusercontent.com/36291746/39614422-6b653088-4f8d-11e8-83fd-05b18ca974a2.PNG">
 
# Terraform Provider for Infoblox
Terraform provider plugin to integrate with Infoblox Network Identity Operating System [NIOS].

The latest version of Infoblox NIOS provider is [v1.1.1](https://github.com/infobloxopen/terraform-provider-infoblox/releases/tag/v1.1.1)

## Building the Provider
* Install and set apt environment variables [Golang](https://golang.org/doc/install) 1.16.x
* Clone the repo and build it
```sh
$ git clone https://github.com/infobloxopen/terraform-provider-infoblox
$ cd terraform-provider-infoblox
$ make build
```

## Developing the Provider
If you wish to work on the provider, follow the above steps to build it.

To test the provider and to run the full suite of acceptance tests run below commands accordingly,
```sh
$ make test
$ make testacc
```

## Using the Provider
* To use the plugin install v0.14.x [Terraform](https://www.terraform.io/downloads.html)
* If you're building the provider, follow the instructions to [install it as a plugin](https://www.terraform.io/docs/cli/config/config-file.html#development-overrides-for-provider-developers). 
* Once the build is completed, set the `terraform-provider-infoblox` binary file location appropriately in in `dev_overrides`.

## NIOS Requirements
* Plugin can be used without a CNA license and does not mandate to specify any EAs.

* If Cloud Network Automation[CNA] License is installed on NIOS and has a Cloud Platform[CP] member attached. Make sure to have below mandatory EAs in .tf file if non mandatory cloud EAs(like "VM Name", "VM ID", "Network Name") are used.
   * Tenant ID :: String Type
   * CMP Type :: String Type
   * Cloud API Owned :: List Type (Values True, False)

## Features of Provider
### Resource
Create, Update and Delete of below resources is supported along with comment and EAs fields.
* IPv4 & IPv6 Network Container
* IPv4 & IPv6 Network
* Allocation & Deallocation of IPv4 or IPv6 address.
* Association & Disassociation of IPv4 or IPv6 address for a VM.

Below resources are still under development and incomplete.
* Creation & Deletion of Network View in NIOS appliance
* Creation and Deletion of A, CNAME, Host, and PTR records

### Data Source
Data Sources for below records are supported.
* IPv4 Network
* A Record
* CNAME Record

## Disclaimer
To use the provider for DNS purposes, a parent (i.e. zone) must already exist. The plugin does not support the creation of zones.
While running acceptance tests create a 10.0.0.0/24 network under default network view and create a reservation for 10.0.0.2 IP
