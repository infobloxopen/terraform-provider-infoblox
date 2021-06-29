 <a href="https://www.infoblox.com">
    <img src="https://avatars.githubusercontent.com/u/8064882?s=400&u=3b245589302c409aff2ce2ba26d95e6df6cfe342&v=4" alt="Infoblox logo" title="Infoblox" align="right" height="50" />
</a> 
 
# Terraform Provider for Infoblox
Terraform provider plugin to integrate with Infoblox Network Identity Operating System [NIOS].
The plugin enables lifecycle management of Infoblox NIOS DDI resources.

The latest version of Infoblox NIOS provider is [v1.1.1](https://github.com/infobloxopen/terraform-provider-infoblox/releases/tag/v1.1.1)
The features in development are available at [`develop`](https://github.com/infobloxopen/terraform-provider-infoblox/tree/develop) branch.

## NIOS Requirements
CNA License need to be installed on NIOS. If CNA is not installed then following default EA's should be added in NIOS side manually:
   * Tenant ID :: String Type
   * CMP Type :: String Type
   * Cloud API Owned :: List Type (Values True, False)
   * Network Name :: String Type
   * VM Name :: String Type
   * VM ID :: String Type

## Quick Starts
- [Using the provider](docs/USING.md)
- [Developing the provider](docs/DEVELOPMENT.md)

## Documentation
The comprehensive documentation of plugin is available at Terraform registry.

https://registry.terraform.io/providers/infobloxopen/infoblox/latest/docs

## Provider features
Provider has NIOS DDI resources as Terraform resources and datasources. Below is the consolidated list of the same.
### Resource
* Network View
* Network
* Allocation & Deallocation of IP from a Network
* Association & Disassociation of IP Address for a VM
* A Record
* PTR Record
* CNAME Record

### Data Source
* Network
* A Record
* CNAME Record
