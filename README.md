<img width="171" alt="Infoblox logo" src="https://user-images.githubusercontent.com/36291746/39614422-6b653088-4f8d-11e8-83fd-05b18ca974a2.PNG">

# Infoblox Terraform Provider
This is a provider plugin for Terraform,
to manage Infoblox NIOS (Network Identity Operating System) resources using
Terraform IaaC solutions.

The latest version of Infoblox provider is [v1.1.1](https://github.com/infobloxopen/terraform-provider-infoblox/releases/tag/v1.1.1)

## Provider features
The provider plugin has NIOS DDI resources represented as Terraform resources and
data sources. Below is the consolidated list of them.
### Resource
* Network View
* Network container
* Network
* A-record
* AAAA-record
* PTR-record
* CNAME-record
* Fixed address and host record as a backend for the following operations:
  * Allocation & de-allocation of IP from a Network
  * Association & de-association of IP Address for a VM
  
For all the resources above, 'comment' and 'ext_attr' properties
updating is supported, except Network View.

### Data Source
Data sources for below records are supported.
* IPv4 Network
* A-record
* CNAME-record

## Disclaimer
To use the provider for DNS purposes, a parent zone must exist in advance.
The plugin does not support the creation of zones.
While running acceptance tests, create a 10.0.0.0/24 network
under the default network view, and create a reservation for the IP address 10.0.0.2.

## NIOS Requirements
* Plugin can be used without a CNA license and does not mandate to specify any extensible attributes (EAs).
* If a Cloud Network Automation (CNA) License is installed on NIOS and
  has a Cloud Platform (CP) member attached, make sure to have the following mandatory
  EAs in TF-files for creating a cloud object (otherwise, the object is not treated as a cloud one):
  * Tenant ID :: String Type
  * CMP Type :: String Type
  * Cloud API Owned :: List Type (Values True, False)

## Using the Provider
There are two ways of using the plugin:

1. Pre-built binary from the Terraform registry.
   You do not need to do anything special,
   the plugin will be installed by Terraform automatically once
   it is mentioned as a requirement in your TF files.
   
2. Build from the source code. This way is suitable for those
   who wants to introduce some customizations,
   which are not in the official plugin yet.
   
## Building a binary from the source code
* Install and set up Go-lang [Golang](https://golang.org/doc/install) of version 1.16 or higher.
* Install Terraform CLI v0.14.x [Terraform](https://www.terraform.io/downloads.html)
* Clone the repo and build it

```sh
$ cd `go env GOPATH`/src
$ mkdir -p github.com/infobloxopen
$ cd github.com/infobloxopen
$ git clone https://github.com/infobloxopen/terraform-provider-infoblox
$ cd terraform-provider-infoblox
$ make build
```

* follow the instructions to [install the resulting binary as a plugin](https://www.terraform.io/docs/cli/config/config-file.html#development-overrides-for-provider-developers).

## Developing the Provider
If you wish to develop the plugin, follow the above steps
to build it after making your changes.

To run the full suite of acceptance tests run these commands:
```sh
$ export INFOBLOX_SERVER=some_ip-addr_or_hostname
$ export INFOBLOX_USERNAME=some_username_on_the_server
$ export INFOBLOX_PASSWORD=appropriate_password
$ export TF_ACC=true # without this only unit tests (not acceptance tests) run
$ make test
$ make testacc
```

Please read comments for the tests in the code and
ensure the conditions mentioned there are met.
For example, you will have to create some objects on
NIOS server before running the tests: DNS zones, views, etc. 
