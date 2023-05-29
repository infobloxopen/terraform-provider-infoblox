<a href="https://www.infoblox.com">
    <img src="https://avatars.githubusercontent.com/u/8064882?s=400&u=3b245589302c409aff2ce2ba26d95e6df6cfe342&v=4" alt="Infoblox logo" title="Infoblox" align="right" height="50" />
</a> 
 
# Infoblox NIOS Terraform Provider

This is a provider plugin for Terraform to manage Infoblox NIOS (Network Identity Operating System) resources using Terraform infrastructure as code solutions.
The plugin enables lifecycle management of Infoblox NIOS DDI resources.

The latest version of Infoblox provider is [v2.4.0](https://github.com/infobloxopen/terraform-provider-infoblox/releases/tag/v2.4.0)

## Provider Features

The provider plugin has NIOS DDI resources represented as Terraform resources and data sources. The consolidated list of supported resources and data sources is as follows:

### Resources:

* Network view (`infoblox_network_view`)
* Network container (`infoblox_ipv4_network_container`, `infoblox_ipv6_network_container`)
* Network (`infoblox_ipv4_network`, `infoblox_ipv6_network`)
* A-record (`infoblox_a_record`)
* AAAA-record (`infoblox_aaaa_record`)
* PTR-record (`infoblox_ptr_record`)
* CNAME-record (`infoblox_cname_record`)
* MX-record (`infoblox_mx_record`)
* TXT-record (`infoblox_txt_record`)
* SRV-record (`infoblox_srv_record`)
* Host record as a backend for the following operations:
    * Allocation and de-allocation of an IP address from a Network (`infoblox_ip_allocation`)
    * Association and de-association of an IP address from a VM (`infoblox_ip_association`)

All of the above resources are supported with `comment` and `ext_attrs` fields.
DNS records and `infoblox_ip_allocation` resource have the `ttl` field's support.

### Data Sources:

* Network View (`infoblox_network_view`)
* IPv4 Network (`infoblox_ipv4_network`)
* IPv4 Network Container (`infoblox_ipv4_network_container`)
* A-record (`infoblox_a_record`)
* AAAA-record (`infoblox_aaaa_record`)
* CNAME-record (`infoblox_cname_record`)
* PTR-record (`infoblox_ptr_record`)
* MX-record (`infoblox_mx_record`)
* TXT-record (`infoblox_txt_record`)
* SRV-record (`infoblox_srv_record`)

All of the above data sources are supported with `comment` and `ext_attr` fields.
DNS records have the `ttl` and `zone` fields' support.

## Quick Start

- [Getting the provider plugin](GETTING.md)
- [Developing on the provider plugin](DEVELOP.md)

## Documentation

The comprehensive documentation of plugin is available at [Terraform registry](https://registry.terraform.io/providers/infobloxopen/infoblox/latest/docs)
and on [Infoblox internet site](https://infoblox-docs.atlassian.net/wiki/spaces/ipamdriverterraform/pages/53055610/Overview+of+Infoblox+IPAM+Plug-In+for+Terraform) as well.

## Prerequisites

Whether you intend to use the published plug-in or the customized version that you have built yourself, you must
complete the following prerequisites:

* Install and set up a physical or virtual Infoblox NIOS appliance that is running on
  NIOS and has necessary licenses installed. To try out the plug-in, you can download and install the evaluation version
  of vNIOS from the [Infoblox Download Center](https://www.infoblox.com/infoblox-download-center).
  For more information, see sections Downloading NIOS and Setting Up NIOS.
* Download and install Terraform (as of now, only version 0.14 is supported).
* Configure the access permissions for Terraform to interact with NIOS Grid objects.
* If you plan to develop a plug-in that includes features that are not in the published version,
  then install the [infblox-go-client](https://github.com/infobloxopen/infoblox-go-client) and Go programming language.
* To use the Infoblox IPAM Plug-In for Terraform, you must either define the following extensible attributes in NIOS or 
  install the Cloud Network Automation license in the NIOS Grid, which adds the extensible attributes by default:
  * `Tenant ID`: String Type 
  * `CMP Type`: String Type 
  * `Cloud API Owned`: List Type (Values: True, False)
* For creation of host records using the `infoblox_ip_allocation` and `infoblox_ip_association` resources,
  you must create the extensible attribute `Terraform Internal ID` of String Type in Infoblox NIOS Grid Manager.
  For steps, refer to the [Infoblox NIOS Documentation](https://infoblox-docs.atlassian.net/wiki/spaces/ILP/pages/15433773).

## Limitations

The limitations of Infoblox IPAM Plug-In for Terraform version 2.3.0 are as follows:

* No support for creating a DNS zone. Therefore, to work with DNS
  records, you must ensure that appropriate DNS zones have been created in NIOS.
* Allocation and association through a fixed-address record are not supported.
* For `infoblox_ip_allocation` and `infoblox_ip_association` resources: creation of a host
  record with multiple IP addresses of the same type is not supported.
  But you can create a host record with a single IPv4 and IPv6 address (of both IP types at the same host record).
* For `infoblox_ipv4_allocation`, `infoblox_ipv6_allocation`, `infoblox_ipv4_association` and `infoblox_ipv6_association`
  resources: creation of a host record with multiple IP addresses of the same type or
  a combination of IPv4 and IPv6 types, is not supported.
* Authority delegation of IP addresses and DNS name spaces to a cloud platform appliance, is not supported.
* Inheritance of extensible attributes is not supported.
* Required extensible attributes specified in NIOS Grid Manager are not validated by the plug-in.
* In NIOS, the gateway IP addresses of networks created using the `infoblox_ipv4_network` and
  `infoblox_ipv6_network` resources display as "IPv4 Reservation" and "IPv6 Fixed Address" respectively.
* Use of capital letters in the domain name of a Terraform resource may lead to unexpected results. For example,
  when you use a Terraform data source to search for a DNS record that has capital letters in its name, no results
  are returned if you specify the name in the same text case. You must specify the name in lower case.
* The import functionality is not supported by the following resources (they are deprecated and not supported anymore):
  * `infoblox_ipv4_allocation`
  * `infoblox_ipv6_allocation`
  * `infoblox_ipv4_association`
  * `infoblox_ipv6_association`
* The Update functionality is currently not working for the CIDR field in A and AAAA records.
* The fetch functionality in data sources returns output for only one matching object even if it finds multiple objects matching the search criteria.
