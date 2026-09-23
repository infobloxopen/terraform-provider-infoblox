# Terraform Provider for Infoblox

The Terraform Provider for Infoblox allows you to manage your Infoblox resources such as DNS records, networks, fixed addresses, and more using Terraform. It is a unified provider: the same provider works with both Infoblox backends, NIOS and UDDI.

| Backend | Description |
|---------|-------------|
| NIOS | Your on-premises Infoblox Grid, accessed over WAPI |
| UDDI | Infoblox Universal DDI, accessed through the Infoblox Portal |

This provider uses the [infoblox-nios-go-client](https://github.com/infobloxopen/infoblox-nios-go-client) for all API calls to the Infoblox NIOS WAPI, and the [universal-ddi-go-client](https://github.com/infobloxopen/universal-ddi-go-client) for all API calls to UDDI.

## Table of Contents

- [Requirements](#requirements)
- [How the Unified Provider Works](#how-the-unified-provider-works)
- [Getting Started](#getting-started)
  - [Configure the Provider](#configure-the-provider)
  - [Prerequisites](#prerequisites)
    - [Setting Up Terraform Internal ID](#setting-up-terraform-internal-id)
- [Managing a NIOS Grid Through the Infoblox Portal](#managing-a-nios-grid-through-the-infoblox-portal)
- [Default Tags for UDDI](#default-tags-for-uddi)
- [Usage Examples](#usage-examples)
- [Available Resources and DataSources](#available-resources-and-datasources)
- [Host Record Management](#host-record-management)
- [Listing Existing Objects](#listing-existing-objects)
- [Importing Existing Resources](#importing-existing-resources)
- [Documentation](#documentation)
- [Logging and Debugging](#logging-and-debugging)
- [Support](#support)

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.12.1
- [Go](https://golang.org/doc/install) >= 1.25.1
- One of:
  - [Infoblox NIOS](https://www.infoblox.com/products/nios/) (version 9.0.6 or higher)
  - An [Infoblox Portal](https://portal.infoblox.com) account

## How the Unified Provider Works

Each resource and data source has a nested block named after the backend. You fill in the block for the backend you configured:

```hcl
// Against NIOS
resource "infoblox_record_a" "example" {
  nios = {
    name     = "web.example.com"
    ipv4addr = "10.0.0.18"
    comment  = "Web server"
  }
}

// Against UDDI - same resource type, different block
resource "infoblox_record_a" "example" {
  uddi = {
    name_in_zone = "web"
    zone         = infoblox_zone_auth.example.id
    rdata = {
      address = "10.0.0.18"
    }
    comment = "Record A"
  }
}
```

The two backends are different products, so their fields are not the same. Keeping them in separate blocks lets you use one provider and one state file for both.

A configuration targets one backend at a time. The provider accepts either a `nios` block or a `uddi` block, not both. To manage both in the same run, declare two provider instances with aliases.

## Getting Started

### Configure the Provider

For NIOS:

```hcl
terraform {
  required_providers {
    infoblox = {
      source  = "infobloxopen/infoblox"
      version = ">= 3.0.0"
    }
  }
}

provider "infoblox" {
  nios = {
    host_url = "https://<NIOS_GRID_IP>"
    username = "<NIOS_USERNAME>"
    password = "<NIOS_PASSWORD>"
  }
}
```

For UDDI:

```hcl
provider "infoblox" {
  uddi = {
    portal_url = "https://csp.infoblox.com"
    portal_key = "<INFOBLOX_PORTAL_API_KEY>"
  }
}
```

The provider also accepts these optional settings:

| Setting | Backend | Description |
|---------|---------|-------------|
| `operation_timeout` | Both | Time in seconds allowed for one operation, including retries. Default `60`. |
| `manage_internal_id_ea` | NIOS | Whether the provider maintains the `Terraform Internal ID` extensible attribute. Default `true`. |
| `uddi.default_tags` | UDDI | Tags applied to every object the provider creates or updates. |

For detailed installation instructions, please refer to the [Quickstart Guide](guides/quickstart.md).

### Prerequisites

#### Setting Up Terraform Internal ID

> This applies to the NIOS backend only.

- A resource can manage its drift state by using the extensible attribute `Terraform Internal ID` when its Reference ID is changed by any manual intervention.
- To use the Terraform Provider for Infoblox NIOS, you must either define the following extensible attributes in NIOS or 
  install the Cloud Network Automation license in the NIOS Grid, which adds the extensible attributes by default:
  * `Tenant ID`: String Type 
  * `CMP Type`: String Type 
  * `Cloud API Owned`: List Type (Values: True, False)
- To use the NIOS Terraform Plugin, you must either define the extensible attribute `Terraform Internal ID`
  in NIOS or use `super user` to execute the below cmd. It will create the read only extensible attribute `Terraform Internal ID`.

  ```shell
  curl -k -u <SUPERUSER>:<PASSWORD> -H "Content-Type: application/json" -X POST https://<NIOS_GRID_IP>/wapi/<WAPI_VERSION>/extensibleattributedef -d '{"name": "Terraform Internal ID", "flags": "CR", "type": "STRING", "comment": "Internal ID for Terraform Resource"}'
  ``` 

  For more details refer to the prerequisites in [Terraform Internal ID](guides/tf_internal_id_management.md) page.

## Managing a NIOS Grid Through the Infoblox Portal

If your NIOS Grid is connected to the Infoblox Portal, you can manage it through the Portal instead of connecting to the Grid directly, by setting `enable_nios_passthru = true` in the `uddi` block.

For detailed information, refer to the [WAPI Passthrough](guides/wapi_passthrough.md) page.

## Default Tags for UDDI

Use `default_tags` in the `uddi` block to apply the same tags to every object the provider creates or updates. A tag set on a resource overrides the default of the same name.

```hcl
provider "infoblox" {
  uddi = {
    portal_url = var.portal_url
    portal_key = var.portal_key
    default_tags = {
      managed_by = "terraform"
    }
  }
}
```

## Usage Examples

Detailed examples for each resource and data source are available in the `examples` directory of the repository. Each resource and data source has its own directory with sample configurations, split by backend:

```
examples/resources/dns/infoblox_record_a/nios/resource.tf
examples/resources/dns/infoblox_record_a/uddi/resource.tf
```

For example:
- Resources examples: [`examples/resources/infoblox_*`](examples/resources/)
- Data sources examples: [`examples/data-sources/infoblox_*`](examples/data-sources/)
- List resources examples: [`examples/list-resources/infoblox_*`](examples/list-resources/)
- Provider configuration examples: [`examples/provider/`](examples/provider/)

Please refer to these examples for detailed usage patterns and configurations.

## Available Resources and DataSources

The object groups available in this provider are categorized as follows:
  - [DHCP](guides/resources_datasources.md#dhcp)
  - [DNS](guides/resources_datasources.md#dns)
  - [DTC](guides/resources_datasources.md#dtc)
  - [RPZ](guides/resources_datasources.md#rpz)
  - [IPAM](guides/resources_datasources.md#ipam)
  - [IPAM FEDERATION](guides/resources_datasources.md#ipam-federation)
  - [GRID](guides/resources_datasources.md#grid)
  - [SECURITY](guides/resources_datasources.md#security)
  - [ACL](guides/resources_datasources.md#acl)
  - [KEYS](guides/resources_datasources.md#keys)
  - [NOTIFICATION](guides/resources_datasources.md#notification)
  - [MISC](guides/resources_datasources.md#miscellaneous)

Not every object is available on both backends. The documentation page for each one shows whether it accepts `nios`, `uddi`, or both.

For a detailed list of available resources and data sources, refer to the [Resources and Data Sources](guides/resources_datasources.md) page.

## Host Record Management

- The `record_host` resource allocates a new IP address from an existing NIOS network and manages the corresponding DNS-related settings. It creates a Host Record in NIOS with either an IPv4 address, an IPv6 address, or both. The IP can be allocated statically (by specifying the address) or dynamically (as the next available address from a network). Once allocated, the address is marked as used in NIOS.

- The `ip_association` resource manages DHCP-related settings of the Host Record created via record_host. It updates the record with VM-specific details such as the MAC address for IPv4 and the DUID for IPv6, enabling full integration with cloud or virtualized environments.

Both resources are available on the NIOS backend only.

Detailed documentation for these resources can be found in [Host Record Documentation](guides/host_record_management.md) page.

## Listing Existing Objects

Every resource has a matching list resource, which finds objects that already exist without importing them into state. Run them with `terraform query`.

For detailed information, refer to the [Listing Existing Objects](guides/list_resources.md) page.

## Importing Existing Resources

Resources that already exist in Infoblox can be brought under Terraform management. Every resource in this provider supports import.

For detailed information, refer to the [Importing Existing Resources](guides/importing_resources.md) page.

## Documentation

For detailed documentation, refer to the [Documentation](guides/documentation_details.md) page.

## Logging and Debugging

For detailed information, refer to the Logging and Debugging page in the docs: [Debugging](guides/logging_debugging.md)

## Support

If you have any questions or issues, you can reach out to us using the following channels:

- Github Issues:
  - Submit your issues or requests for enhancements on the [Github Issues Page](https://github.com/infobloxopen/terraform-provider-infoblox/issues)
- Infoblox Support:
  - For any questions or issues, please contact [Infoblox Support](https://info.infoblox.com/contact-form/).
