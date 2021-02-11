# Terraform Provider for Infoblox
 <img width="171" alt="capture" src="https://user-images.githubusercontent.com/36291746/39614422-6b653088-4f8d-11e8-83fd-05b18ca974a2.PNG">

## Requirements
* [Terraform](https://www.terraform.io/downloads.html) 0.11.x or greater
* [Go](https://golang.org/doc/install) 1.12.x (to build the provider plugin)
* CNA License need to be installed on NIOS. If CNA is not installed then following default EA's should be added in NIOS side:
   * VM Name :: String Type
   * VM ID :: String Type
   * Tenant ID :: String Type
   * CMP Type :: String Type
   * Cloud API Owned :: List Type (Values True, False)
   * Network Name :: String Type

## Building the Provider

```sh
$ git clone https://github.com/infobloxopen/terraform-provider-infoblox
$ cd terraform-provider-infoblox
$ make build
```

## Using the Provider
If you're building the provider, follow the instructions to [install it as a plugin](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin). After the build is complete, copy the `terraform-provider-infoblox` binary into the same path as your terraform binary. After placing it into your plugins directory, run `terraform init` to initialize it.

## Developing the Provider
If you wish to work on the provider, you'll first need Go installed on your machine (version 1.12+ is required).

To compile the provider, run the following steps:
```sh
$ make build
...
$ ./terraform-provider-infoblox
...
```
To test the provider, you can simply run `make test`.
```sh
$ make test
```

In order to run the full suite of acceptance tests `make testacc`.
```sh
$ make testacc
```
## Features of Provider
### Resource
* Creation of Network View in NIOS appliance
* Creation & Deletion of Network in NIOS appliance
* Allocation & Deallocation of IP from a Network
* Association & Disassociation of IP Address for a VM
* Creation and Deletion of A, CNAME, Host, and Ptr records
* Allocate Next Available Network from a parent network container

### Data Source
* Supports Data Source for Network

## Disclaimer
To use the provider for DNS purposes, a parent (i.e. zone) must already exist. The plugin does not support the creation of zones.
while running acceptance tests create a 10.0.0.0/24 network under default network view and create a reservation for 10.0.0.2 IP
