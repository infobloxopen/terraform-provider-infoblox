<a href="https://www.infoblox.com">
    <img src="https://avatars.githubusercontent.com/u/8064882?s=400&u=3b245589302c409aff2ce2ba26d95e6df6cfe342&v=4" alt="Infoblox logo" title="Infoblox" align="right" height="50" />
</a> 
 
# Infoblox NIOS Terraform Provider
This is a provider plugin for Terraform to manage Infoblox NIOS (Network Identity Operating System) resources using Terraform infrastructure as code solutions.
The plugin enables lifecycle management of Infoblox NIOS DDI resources.

The latest version of Infoblox provider is [v2.0.0](https://github.com/infobloxopen/terraform-provider-infoblox/releases/tag/v2.0.0)

The features and bug fixes under development will be available in the [`develop`]((https://github.com/infobloxopen/terraform-provider-infoblox/tree/develop)) branch.

## Provider Features
The provider plugin has NIOS DDI resources represented as Terraform resources and data sources. The consolidated list of supported resources and data sources is as follows:

### Resources:
* Network view
* Network container
* Network
* A record
* AAAA record
* PTR record
* CNAME record
* Host record as a backend for the following operations:
    * Allocation and deallocation of IP address from a Network
    * Association and dissociation of IP address from a VM

All of the above resources are supported with `comment` and `ext_attr` fields.
DNS records have the `ttl` field support.

### Data Sources:
* IPv4 Network
* A record
* CNAME record

All of the above data sources are supported with `comment` and `ext_attr` fields.
DNS records have the `ttl` field support.

## Quick Start
- [Getting the provider plugin](docs/GETTING.md)
- [Developing on the provider plugin](docs/DEVELOP.md)

## Documentation
The comprehensive documentation of plugin is available at Terraform registry.

https://registry.terraform.io/providers/infobloxopen/infoblox/latest/docs

## NIOS Requirements
* Plugin (from v2.0.0 onwards) can be used without a Cloud Network Automation (CNA) license on NIOS Grid. 
* If a CNA license is installed, cloud objects can be created and tracked on the Cloud tab in NIOS Grid Master(GM).

## Terraform Configuration file requirements
* Users are not mandated to specify any EAs in tf file when there is no CNA license in the grid.
* User must have the following EAs in tf files to create cloud objects when a CNA license is installed in GM:
    * Tenant ID :: String Type
    * CMP Type :: String Type
    * Cloud API Owned :: List Type (Values True, False)