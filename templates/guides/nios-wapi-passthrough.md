---
page_title: "Managing NIOS through the Infoblox Portal (WAPI Passthrough)"
subcategory: "Guides"
description: |-
  Configure the Infoblox provider to manage an on-prem NIOS Grid through the Infoblox Portal, without direct network access to the Grid.
---

# Managing NIOS through the Infoblox Portal (WAPI Passthrough)

The WAPI Passthrough feature enables secure, remote management of on-prem Infoblox NIOS environments by exposing the existing Web API (WAPI) interface through the Infoblox Portal. Traditionally, administrators and automation tools could access NIOS WAPI endpoints only from within the local network, which limited the flexibility of automation workflows. WAPI Passthrough overcomes this limitation by allowing WAPI operations to be invoked through the Infoblox Portal, which securely proxies each request to the appropriate on-prem NIOS Grid.

By default, the Infoblox provider connects to a Grid's WAPI endpoint directly, which requires the host running Terraform to have network access to the Grid and to authenticate with NIOS administrator credentials. When passthrough is enabled, the provider issues the same requests to the Portal instead. The API itself is unchanged - the same objects, the same fields, and full CRUD support - so only the route and the credentials differ:

| | Direct Grid | Through the Portal |
|---|---|---|
| Credentials | NIOS username + password | Portal service API key + Grid license UID |
| Network reachability | Terraform must reach the Grid | Terraform needs access to Infoblox Portal |

## How the Provider Supports It

Passthrough is a **transport option, not a separate backend**. When it is enabled, the provider builds a NIOS client that targets the Portal endpoint and authenticates with the Portal API key and the Grid's license UID. Everything above the transport layer is unchanged:

- The entire NIOS object set is available over passthrough: every **resource**, **data source**, and **list resource** the provider supports. No object requires additional configuration, and none are excluded.
- Resource configuration remains in the **`nios = { ... }` block**, exactly as for a direct Grid connection.
- Attributes and filters documented as NIOS-only (`ext_attrs`, `ext_attr_filters`, `max_results`, and so on) apply. The UDDI-only ones do not.
- A configuration written for a direct Grid connection works unchanged over passthrough. Only the `provider` block differs.

To use passthrough, set `enable_nios_passthru = true` in the provider block. The provider does not infer the mode from any other value.

## Configuring the Provider

Passthrough credentials are supplied in the `uddi` block, because the credentials are Portal credentials. The `nios` block is not used in this mode, and the provider rejects a configuration that sets both blocks.

````terraform
provider "infoblox" {
  uddi = {
    enable_nios_passthru = true
    portal_url           = "<INFOBLOX_PORTAL_WAPI_URL>"
    portal_key           = var.infoblox_portal_key
    nios_license_uid     = var.nios_license_uid
  }
}
````

| Attribute | Required for passthrough | Notes |
|---|---|---|
| `enable_nios_passthru` | Yes | Must be `true`, and must be known at plan time. The provider selects the transport during planning, so it cannot depend on a value computed during apply. |
| `portal_url` | Yes | The Portal **WAPI** endpoint, not the Portal UI or the CSP API endpoint. Specify the host only, the provider appends the WAPI base path and version. |
| `portal_key` | Yes | Portal service API key. Sensitive. |
| `nios_license_uid` | Yes | License UID of the target Grid. Sensitive. It is also shown as a tag on the equivalent NIOS Infra Host in the Infoblox Portal. |

> **Note:** The WAPI endpoint is a different host from the Portal CSP API endpoint used for Universal DDI objects, and in non-production or regional environments its name does not follow from the Portal URL. Copy the exact host from the Infoblox Portal rather than deriving it. TLS certificates are always verified on this route, so an approximate host name fails with a certificate error instead of connecting.

### Using Environment Variables

Any of the three values may be omitted from the configuration and supplied through an environment variable instead:

| Attribute | Environment variable |
|---|---|
| `portal_url` | `INFOBLOX_PORTAL_URL` |
| `portal_key` | `INFOBLOX_PORTAL_KEY` |
| `nios_license_uid` | `NIOS_LICENSE_UID` |

`enable_nios_passthru` has no environment variable equivalent and must always be set in the provider block.

## Example

The following configuration creates an authoritative zone and an A record on the Grid. Apart from the `provider` block, it is identical to a configuration written for a direct Grid connection:

````terraform
provider "infoblox" {
  uddi = {
    enable_nios_passthru = true
    portal_url           = "<INFOBLOX_PORTAL_WAPI_URL>"
    portal_key           = var.infoblox_portal_key
    nios_license_uid     = var.nios_license_uid
  }
}

resource "infoblox_zone_auth" "example" {
  nios = {
    fqdn    = "example.com"
    comment = "Managed by Terraform through the Infoblox Portal"
  }
}

resource "infoblox_record_a" "web" {
  nios = {
    name     = "web.${infoblox_zone_auth.example.nios.fqdn}"
    ipv4addr = "10.0.0.20"
    comment  = "Web server"
    ext_attrs = {
      Site = "location-1"
    }
  }
}
````

### Multiple Grids and Universal DDI

Each provider configuration targets a single Grid through its license UID. Use a provider alias for each Grid, and a further alias without the passthrough flag for Universal DDI objects. The same Portal API key authenticates all of them:

````terraform
// NIOS Grid, through the Portal
provider "infoblox" {
  alias = "grid_emea"
  uddi = {
    enable_nios_passthru = true
    portal_url           = "<INFOBLOX_PORTAL_WAPI_URL>"
    portal_key           = var.infoblox_portal_key
    nios_license_uid     = var.grid_emea_license_uid
  }
}

// Universal DDI objects, through the CSP API
provider "infoblox" {
  alias = "uddi"
  uddi = {
    portal_url = "<INFOBLOX_PORTAL_URL>"
    portal_key = var.infoblox_portal_key
  }
}

// A NIOS object on the Grid, using the nios block
resource "infoblox_record_a" "emea" {
  provider = infoblox.grid_emea
  nios = {
    name     = "web.emea.example.com"
    ipv4addr = "10.10.0.20"
  }
}

// A Universal DDI object, using the uddi block
resource "infoblox_zone_auth" "cloud" {
  provider = infoblox.uddi
  uddi = {
    fqdn         = "example.com."
    primary_type = "cloud"
  }
}
````

Each resource selects its provider with the `provider` argument, and the block it populates follows from that provider: resources on a passthrough provider use `nios`, and resources on a Universal DDI provider use `uddi`.

## Reference

To obtain the Portal API key and the Grid license UID, and to review the permissions and roles that govern passthrough access, refer to [Configuring WAPI Passthrough in the Infoblox Portal](https://docs.infoblox.com/space/BloxOneDDI/1505493401/Configuring+WAPI+Passthrough+in+the+Infoblox+Portal).
