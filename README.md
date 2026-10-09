# Terraform Provider for Infoblox

> [!WARNING]
> **Version 2.x.x is deprecated**  
> As of October 2026, version 2.x.x of this provider is no longer officially supported or maintained. We strongly recommend upgrading to version 3.x.x. For more details, see [Migrating from Other Infoblox Providers](#migrating-from-other-infoblox-providers).

The Terraform Provider for Infoblox allows you to manage your DDI Infrastructure such as DNS records, Networks, Fixed Addresses, and more using Terraform. It is a unified provider: the same provider works with both Infoblox backends, NIOS and UDDI.

| Backend | Description |
|---------|-------------|
| NIOS | Your on-premises Infoblox Grid, accessed over WAPI |
| UDDI | Infoblox Universal DDI, accessed through the Infoblox Portal |

This provider uses the [infoblox-nios-go-client](https://github.com/infobloxopen/infoblox-nios-go-client) for all API calls to the Infoblox NIOS WAPI, and the [universal-ddi-go-client](https://github.com/infobloxopen/universal-ddi-go-client) for all API calls to UDDI.

## Table of Contents

- [Requirements](#requirements)
- [How the Provider Works](#how-the-provider-works)
- [Getting Started](#getting-started)
  - [Configure the Provider](#configure-the-provider)
  - [Proxy Settings](#proxy-settings)
  - [Retry Behavior](#retry-behavior)
  - [Prerequisites](#prerequisites)
    - [Setting Up Terraform Internal ID](#setting-up-terraform-internal-id)
- [Managing a NIOS Grid Through the Infoblox Portal](#managing-a-nios-grid-through-the-infoblox-portal)
- [Usage Examples](#usage-examples)
- [Available Resources and DataSources](#available-resources-and-datasources)
- [Migrating from Other Infoblox Providers](#migrating-from-other-infoblox-providers)
- [Host Record Management](#host-record-management)
- [Listing Existing Objects](#listing-existing-objects)
- [Importing Existing Resources](#importing-existing-resources)
- [Update Trigger](#update-trigger)
- [Roadmap](#roadmap)
- [Documentation](#documentation)
- [Logging and Debugging](#logging-and-debugging)
- [Contributing](#contributing)
- [Support](#support)

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.12.1
- [Go](https://golang.org/doc/install) >= 1.25.8
- One of:
  - Infoblox NIOS (version 9.0.6 or higher, WAPI v2.13.6)
  - An Infoblox Portal account

The objects you manage must be licensed on your backend. For example, RPZ objects need an RPZ license, and Threat Defense objects need a Threat Defense subscription etc.

## How the Provider Works

Each resource and data source has a nested block named after the backend. You fill in the block for the backend you configured:

```hcl
// NIOS Managed Record A
resource "infoblox_record_a" "example" {
  nios = {
    name     = "web.example.com"
    ipv4addr = "10.0.0.18"
    comment  = "Web server"
  }
}

// UDDI Managed Record A 
// Uses same resource type, but different block
resource "infoblox_record_a" "example" {
  uddi = {
    name_in_zone = "web"
    zone         = <dns/auth_zone/{id}>
    rdata = {
      address = "10.0.0.18"
    }
    comment = "Record A"
  }
}
```

A configuration targets one backend at a time. The provider accepts either a `nios` block or a `uddi` block, not both. To manage both in the same run, declare two provider instances with [aliases](https://developer.hashicorp.com/terraform/language/providers/configuration#alias-multiple-provider-configurations) and set `provider` on each resource:

```hcl
provider "infoblox" {
  nios = {
    host_url = "https://<NIOS_GRID_IP>"
    username = "<NIOS_USERNAME>"
    password = "<NIOS_PASSWORD>"
  }
}

provider "infoblox" {
  alias = "uddi"
  uddi = {
    portal_url = "<INFOBLOX_PORTAL_URL>"
    portal_key = "<INFOBLOX_PORTAL_API_KEY>"
  }
}

// Uses the default (NIOS) provider
resource "infoblox_record_a" "nios_example" {
  nios = {
    name     = "web.example.com"
    ipv4addr = "10.0.0.18"
  }
}

// Uses the aliased UDDI provider
resource "infoblox_record_a" "uddi_example" {
  provider = infoblox.uddi
  uddi = {
    name_in_zone = "web"
    zone         = <dns/auth_zone/{id}>
    rdata = {
      address = "10.0.0.18"
    }
  }
}
```

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

For detailed installation instructions, please refer to the [Quickstart Guide](docs/guides/quickstart.md).
### Proxy Settings

These settings apply to the NIOS backend only:

- `nios.proxy_url`: HTTP proxy URL to route NIOS WAPI calls through, for example `http://proxy.example.com:8080`.
- `proxy_search`: Where WAPI requests are processed. `LOCAL` (default) processes them on the member you connect to, and `GM` redirects them to the Grid Master.

For detailed installation instructions, please refer to the [Quickstart Guide](guides/quickstart.md).

### Retry Behavior

The provider retries API calls that fail with a temporary error. All attempts of one operation must finish within `operation_timeout`, which defaults to `60` seconds. To allow more time for retries, increase it:

```hcl
provider "infoblox" {
  operation_timeout = 120
  nios = { ... }
}
```

### Prerequisites

#### Setting Up Terraform Internal ID

> This applies to the NIOS backend only.

- A resource can manage its drift state by using the extensible attribute `Terraform Internal ID` when its Reference ID is changed by any manual intervention.
- To use the Terraform Provider for Infoblox NIOS, you must either define the following extensible attributes in NIOS or 
  install the Cloud Network Automation license in the NIOS Grid, which adds the extensible attributes by default:
  * `Tenant ID`: String Type 
  * `CMP Type`: String Type 
  * `Cloud API Owned`: List Type (Values: True, False)
- To manage resources on NIOS using the Provider, you must either define the extensible attribute `Terraform Internal ID`
  in NIOS or use `super user` to execute the below cmd. It will create the read only extensible attribute `Terraform Internal ID`.

  ```shell
  curl -k -u <SUPERUSER>:<PASSWORD> -H "Content-Type: application/json" -X POST https://<NIOS_GRID_IP>/wapi/<WAPI_VERSION>/extensibleattributedef -d '{"name": "Terraform Internal ID", "flags": "CR", "type": "STRING", "comment": "Internal ID for Terraform Resource"}'
  ``` 

  For more details refer to the prerequisites in [Terraform Internal ID](docs/guides/tf-internal-id-management.md) page.

## Managing a NIOS Grid Through the Infoblox Portal using WAPI Passthru

If your NIOS Grid is connected to the Infoblox Portal, you can manage it through the Portal instead of connecting to the Grid directly, by setting `enable_nios_passthru = true` in the `uddi` block.

For detailed information, refer to the [WAPI Passthrough](guides/wapi_passthrough.md) page.


## Usage Examples

Detailed examples for each resource and data source are available in the `examples` directory of the repository. Each resource and data source has its own directory with sample configurations, split by backend:

```
examples/resources/dns/infoblox_record_a/nios/resource.tf
examples/resources/dns/infoblox_record_a/uddi/resource.tf
```

Example Directory :
- Resources examples: [`examples/resources/infoblox_*`](examples/resources/)
- Data sources examples: [`examples/data-sources/infoblox_*`](examples/data-sources/)
- List resources examples: [`examples/list-resources/infoblox_*`](examples/list-resources/)
- Provider configuration examples: [`examples/provider/`](examples/provider/)

Please refer to these examples for detailed usage patterns and configurations.

## Available Resources and DataSources

The object groups available in this provider are categorized as follows:

  - [ACL](docs/guides/resources-datasources.md#acl)
  - [ANYCAST](docs/guides/resources-datasources.md#anycast)
  - [CLOUD](docs/guides/resources-datasources.md#cloud)
  - [CLOUD DISCOVERY](docs/guides/resources-datasources.md#cloud-discovery)
  - [DHCP](docs/guides/resources-datasources.md#dhcp)
  - [DISCOVERY](docs/guides/resources-datasources.md#discovery)
  - [DNS](docs/guides/resources-datasources.md#dns)
  - [DTC](docs/guides/resources-datasources.md#dtc)
  - [FW](docs/guides/resources-datasources.md#fw)
  - [GRID](docs/guides/resources-datasources.md#grid)
  - [INFRA](docs/guides/resources-datasources.md#infra)
  - [IPAM](docs/guides/resources-datasources.md#ipam)
  - [IPAM FEDERATION](docs/guides/resources-datasources.md#ipam-federation)
  - [KEYS](docs/guides/resources-datasources.md#keys)
  - [MISC](docs/guides/resources-datasources.md#miscellaneous)
  - [NOTIFICATION](docs/guides/resources-datasources.md#notification)
  - [REDIRECT](docs/guides/resources-datasources.md#redirect)
  - [RIR](docs/guides/resources-datasources.md#rir)
  - [RPZ](docs/guides/resources-datasources.md#rpz)
  - [SECURITY](docs/guides/resources-datasources.md#security)

Not every object is available on both backends. The guide shows whether each one supports NIOS, UDDI, or both.

For a detailed list of available resources and data sources, refer to the [Resources and Data Sources](docs/guides/resources-datasources.md) page.

## Migrating from Other Infoblox Providers

This provider replaces the legacy Infoblox provider (2.x), the NIOS provider, and the BloxOne provider. To find the unified name for each resource and data source, refer to the [Migration Object Mapping](docs/guides/object-mapping.md) page:

- [Legacy Infoblox Provider](docs/guides/object-mapping.md#legacy-infoblox-provider-to-unified-provider)
- [NIOS Provider](docs/guides/object-mapping.md#nios-provider-to-unified-provider)
- [BloxOne Provider](docs/guides/object-mapping.md#bloxone-provider-to-unified-provider)

## Host Record Management

- The `record_host` resource allocates a new IP address from an existing NIOS network and manages the corresponding DNS-related settings. It creates a Host Record in NIOS with either an IPv4 address, an IPv6 address, or both. The IP can be allocated statically (by specifying the address) or dynamically (as the next available address from a network). Once allocated, the address is marked as used in NIOS.

- The `ip_association` resource manages DHCP-related settings of the Host Record created via record_host. It updates the record with VM-specific details such as the MAC address for IPv4 and the DUID for IPv6, enabling full integration with cloud or virtualized environments.

Both resources are available on the NIOS backend only.

Detailed documentation for these resources can be found in [Host Record Documentation](docs/guides/host-record-management.md) page.

## Listing Existing Objects

Every resource has a corresponding list resource, which defines a structured query block that discovers existing cloud infrastructure so you can map or import those resources into your workspace. Run them with `terraform query`.

> [!NOTE]
> List resources and `terraform query` require Terraform v1.14.0 or later.

For detailed information, refer to the [Listing Existing Objects](guides/list-resources.md) page.

## Importing Existing Resources

Resources created externally in Infoblox NIOS or UDDI can be brought under Terraform management. Resources can be imported in Terraform in several ways.

For detailed information, refer to the [Importing Existing Resources](docs/guides/importing-resources.md) page.

## Update Trigger

Most resources have an optional `update_trigger` attribute. The provider never sends it to the API. Change its value when you want Terraform to run an update, even though Terraform reports no changes. The field accepts any string value, such as an incrementing counter, a timestamp, or another unique value:

```hcl
resource "infoblox_record_a" "example" {
  update_trigger = "v2"
  nios = {
    name     = "web.example.com"
    ipv4addr = "10.0.0.18"
  }
}
```

## Roadmap

Support for additional NIOS and Universal DDI objects, and for infrastructure deployment modules, is planned for later releases.

For the list of what is planned, refer to the [Roadmap](ROADMAP.md) page.

## Documentation

For detailed documentation, refer to the [Documentation](guides/documentation-details.md) page.

## Logging and Debugging

For detailed information, refer to the Logging and Debugging page in the docs: [Debugging](docs/guides/logging-debugging.md)

## Contributing

We encourage you to open an issue rather than a pull request for code changes, as every change has to be tested on both the NIOS and UDDI backends. For details, refer to the [Contributing Guide](CONTRIBUTING.md).

## Support

If you have any questions or issues, you can reach out to us using the following channels:

- Github Issues:
  - Submit your issues or requests for enhancements on the [Github Issues Page](https://github.com/infobloxopen/terraform-provider-infoblox/issues)
- Infoblox Support:
  - For any questions or issues, please contact [Infoblox Support](https://info.infoblox.com/contact-form/).
