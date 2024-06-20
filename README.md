<a href="https://www.infoblox.com">
    <img src="https://avatars.githubusercontent.com/u/8064882?s=400&u=3b245589302c409aff2ce2ba26d95e6df6cfe342&v=4" alt="Infoblox logo" title="Infoblox" align="right" height="50" />
</a> 
 
# Infoblox NIOS Terraform Provider

This is a provider plug-in for Terraform to manage Infoblox NIOS (Network Identity Operating System) resources using Terraform infrastructure as code solutions.
The plug-in enables lifecycle management of Infoblox NIOS DDI resources.

The latest version of Infoblox provider is [v2.7.0](https://github.com/infobloxopen/terraform-provider-infoblox/releases/tag/v2.7.0)

## Provider Features

The provider plug-in has NIOS DDI resources represented as Terraform resources and data sources. The consolidated list of supported resources and data sources is as follows:

### Resources:

* Network view (`infoblox_network_view`)
* Network container (`infoblox_ipv4_network_container`, `infoblox_ipv6_network_container`)
* Network (`infoblox_ipv4_network`, `infoblox_ipv6_network`)
* A-record (`infoblox_a_record`)
* AAAA-record (`infoblox_aaaa_record`)
* DNS View (`infoblox_dns_view`)
* PTR-record (`infoblox_ptr_record`)
* CNAME-record (`infoblox_cname_record`)
* MX-record (`infoblox_mx_record`)
* TXT-record (`infoblox_txt_record`)
* SRV-record (`infoblox_srv_record`)
* Zone Auth (`infoblox_zone_auth`)
* Host record as a backend for the following operations:
    * Allocation and deallocation of an IP address from a Network (`infoblox_ip_allocation`)
    * Association and disassociation of an IP address from a VM (`infoblox_ip_association`)
* Zone Forward (`infoblox_zone_forward`)

All of the above resources are supported with `comment` and `ext_attrs` fields.
DNS records and the `infoblox_ip_allocation` resources are supported with `ttl` field.
<br> A resource can manage its drift state by using the extensible attribute `Terraform Internal ID` when its Reference ID is changed by any manual intervention.


### Data Sources:

* Network View (`infoblox_network_view`)
* IPv4 Network (`infoblox_ipv4_network`)
* IPv4 Network Container (`infoblox_ipv4_network_container`)
* A-record (`infoblox_a_record`)
* AAAA-record (`infoblox_aaaa_record`)
* DNS View (`infoblox_dns_view`)
* CNAME-record (`infoblox_cname_record`)
* PTR-record (`infoblox_ptr_record`)
* MX-record (`infoblox_mx_record`)
* TXT-record (`infoblox_txt_record`)
* SRV-record (`infoblox_srv_record`)
* Zone Auth (`infoblox_zone_auth`)
* Zone Forward (`infoblox_zone_forward`)
* IPv6 Network (`infoblox_ipv6_network`)
* IPv6 Network Container (`infoblox_ipv6_network_container`)
* Host-record (`infoblox_host_record`)

All of the above data sources are supported with `comment` and `ext_attr` fields.
Data source of DNS records are supported with `ttl` and `zone` fields.

## Quick Start

- [Getting the provider plugin](GETTING.md)
- [Developing on the provider plugin](DEVELOP.md)

## Documentation

The comprehensive documentation of the plug-in is available at [Terraform registry](https://registry.terraform.io/providers/infobloxopen/infoblox/latest/docs)
and on [Infoblox internet site](https://docs.infoblox.com/space/ipamdriverterraform/17006594/Infoblox+IPAM+Driver+for+Terraform) as well.

## Prerequisites

Whether you intend to use the published plug-in or the customized version that you have built by yourself, you must
complete the following prerequisites:

* Install and set up a physical or virtual Infoblox NIOS appliance that is running on
  NIOS and has necessary licenses installed. To try out the plug-in, you can download and install the evaluation version
  of vNIOS from the [Infoblox Download Center](https://www.infoblox.com/infoblox-download-center).
  For more information, see sections Downloading NIOS and Setting Up NIOS.
* Download and install Terraform (as of now, only version 0.14 is supported).
* Configure the access permissions for Terraform to interact with NIOS Grid objects.
* If you plan to develop a plug-in that includes features that are not in the published version,
  then install the [infoblox-go-client](https://github.com/infobloxopen/infoblox-go-client) and Go programming language.
* To use the Infoblox IPAM Plug-In for Terraform, you must either define the following extensible attributes in NIOS or 
  install the Cloud Network Automation license in the NIOS Grid, which adds the extensible attributes by default:
  * `Tenant ID`: String Type 
  * `CMP Type`: String Type 
  * `Cloud API Owned`: List Type (Values: True, False)
* To use the Infoblox IPAM Plug-In for Terraform, you must either define the extensible attribute `Terraform Internal ID`
  in NIOS or use `super user` to execute the below cmd. It will create the read only extensible attribute `Terraform Internal ID`. 
  for more details refer to the [Infoblox NIOS Documentation](https://docs.infoblox.com/space/NIOS/35400616/NIOS).
  ```shell
  curl -k -u <SUPERUSER>:<PASSWORD> -H "Content-Type: application/json" -X POST https://<NIOS_GRID_IP>/wapi/<WAPI_VERSION>/extensibleattributedef -d '{"name": "Terraform Internal ID", "flags": "CR", "type": "STRING", "comment": "Internal ID for Terraform Resource"}'
  ```
## Limitations

The limitations of Infoblox IPAM Plug-In for Terraform are as follows:

* Allocation and association through a fixed-address record are not supported.
* For `infoblox_ip_allocation` and `infoblox_ip_association` resources: creation of a host
  record with multiple IP addresses of the same type is not supported.
  But you can create a host record with a single IPv4 and IPv6 address (of both IP types at the same host record).
* Authority delegation of IP addresses and DNS name spaces to a cloud platform appliance, is not supported.
* Inheritance of extensible attributes is not fully functional in this release. Infoblox supports only the retaining of
  inherited extensible attributes values in NIOS. The values are no longer deleted from NIOS as a result of any
  operation performed in Terraform.
* Configuring an A, AAAA, and a host record resource with both cidr and ip_addr parameters, or
  configuring a PTR record with a combination of cidr , ip_addr , and record_name parameters, may
  lead to unexpected behavior. For a consistent behavior, configure any one of the input parameters.
* Required extensible attributes specified in NIOS Grid Manager are not validated by the plug-in.
* In NIOS, the gateway IP addresses of networks created using the `infoblox_ipv4_network` and
  `infoblox_ipv6_network` resources display as "IPv4 Reservation" and "IPv6 Fixed Address" respectively.
* Use of capital letters in the domain name of a Terraform resource may lead to unexpected results. For example,
  when you use a Terraform data source to search for a DNS record that has capital letters in its name, no results
  are returned if you specify the name in the same text case. You must specify the name in lower case.
* In plug-in versions prior to `v2.5.0`, the fetch functionality in data sources returns output for only one matching 
  object even if it finds multiple objects matching the search criteria.
* When using the Terraform `import` block for a resource, a new Terraform internal ID is assigned to the resource when 
  the `terraform plan` command is run for the first time. If a subsequent `terraform apply` is aborted, the record will 
  still retain the `Terraform Internal ID` though the resource is not managed by Terraform.

## Best Practices

* Infoblox recommends that you manage all resources supported by IPAM Plug-In for Terraform from Terraform only. 
  Modifying a resource outside of Terraform may result in unexpected behavior.
* If you need to manage a large number of resources, Infoblox recommends that you manage them across multiple workspaces
  instead of using a single state file to manage all resources. For more information, see [Managing Workspaces](https://developer.hashicorp.com/terraform/cli/workspaces) 
  and [Structuring Terraform Configuration](https://www.hashicorp.com/blog/structuring-hashicorp-terraform-configuration-for-production).
