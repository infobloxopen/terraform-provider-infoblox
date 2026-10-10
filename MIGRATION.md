# Migrating to the Infoblox Terraform Provider

The Infoblox provider (`infobloxopen/infoblox` v3.0.0 and later) manages both NIOS and UDDI from a single provider. It replaces the following providers:

| Provider | Registry source | Backend |
|----------|-----------------|---------|
| BloxOne provider | `infobloxopen/bloxone` | UDDI |
| Legacy Infoblox provider (v2.x) | `infobloxopen/infoblox` | NIOS |
| NIOS provider | `infobloxopen/nios` | NIOS |

 Resource names, attributes, and state format are different, so you need to update your configuration. Your objects in NIOS and UDDI are not affected. The migration only moves them from the old provider's state to the unified provider's state, without deleting or recreating them.

This guide describes the common migration steps, followed by the changes specific to each provider:

* [Migrating from the BloxOne Provider](#migrating-from-the-bloxone-provider)
* [Migrating from the Legacy Infoblox Provider](#migrating-from-the-legacy-infoblox-provider)
* [Migrating from the NIOS Provider](#migrating-from-the-nios-provider)

## Prerequisites

* Terraform v1.12.1 or later
* For NIOS: Infoblox NIOS 9.0.6 (WAPI v2.13.6)
* For UDDI: an Infoblox Portal account and API key
* A backup of your current Terraform state


## Migration Steps

### Back Up Your State

```bash
terraform state pull > terraform.tfstate.backup
```

Keep this backup until the migration is complete and `terraform plan` shows no changes.

## Replace Resources in State

### Get Resource IDs

First, get the IDs of all existing resources:

```bash
terraform show -json | jq -c '.values.root_module.resources[] | {"resource":.address, "id":.values.id}'
```

### Remove Old Resource from State

Remove the old resource from state:

```bash
terraform state rm infoblox_a_record.example
```

> [!WARNING]
> Always run `terraform state rm` before you remove or rewrite the old resource blocks. If you delete a resource block while the resource is still in the state, Terraform plans to destroy the object.

### Update the Provider Configuration

Replace the old provider in the `terraform` block and the `provider` block with the unified provider, then run `terraform init -upgrade`:

```terraform
terraform {
  required_providers {
    infoblox = {
      source  = "infobloxopen/infoblox"
      version = ">= 3.0.0"
    }
  }
}

// NIOS backend
provider "infoblox" {
  nios = {
    host_url = "https://<NIOS_GRID_IP>"
    username = "<NIOS_USERNAME>"
    password = "<NIOS_PASSWORD>"
  }
}

// UDDI backend
provider "infoblox" {
  alias = "uddi"
  uddi = {
    portal_url = "<INFOBLOX_PORTAL_URL>"
    portal_key = "<INFOBLOX_PORTAL_KEY>"
  }
}
```

A provider block configures either `nios` or `uddi`, not both. To manage both backends in one configuration, use two provider blocks with an alias, as shown above.

> [!NOTE]
> Credentials can also be set with environment variables: `NIOS_HOST_URL`, `NIOS_USERNAME`, and `NIOS_PASSWORD` for NIOS, or `INFOBLOX_PORTAL_URL` and `INFOBLOX_PORTAL_KEY` for UDDI. You still need an empty `nios = {}` or `uddi = {}` block to select the backend.

### Rewrite the Resource Blocks

Rename each resource to its unified provider type and move its attributes into a `nios` or `uddi` block.

### Import New Resource into State

**Recommended approach:** Add an `import` block with the ID you recorded earlier, next to the rewritten resource block:

```terraform
import {
  to = infoblox_record_a.example
  id = "record:a/ZG5zLmEkLl9kZWZhdWx0LmNvbS5pbmZvYmxveC50ZXN0:a-record.example.com/default"
}

resource "infoblox_record_a" "example" {
  nios = {
    name     = "a-record.example.com"
    ipv4addr = "10.0.0.1"
    view     = "default"
  }
}
```

Run terraform plan to check the import and then apply it.

```bash 
terraform plan
terraform apply
``` 

After the apply, run `terraform plan` it should shows no changes.

> [!NOTE]
> **NIOS backend:** When you import a NIOS object, the provider adds a new `Terraform Internal ID` extensible attribute to it, replacing any value set by the old provider. The plan shows each resource as 1 to import and 1 to update. For more information, see [Management of Provider Operations with Terraform Internal ID](tf-internal-id-management.md).

You can also import with the CLI instead of an `import` block:

```bash
terraform import infoblox_record_a.example "record:a/ZG5zLmEkLl9kZWZhdWx0LmNvbS5pbmZvYmxveC50ZXN0:a-record.example.com/default"
```

Unlike import blocks, this command changes the state immediately and does not show a plan first.

> For **NIOS backend** The first `terraform apply` after the import attaches the `Terraform Internal ID` extensible attribute to the imported object. Plans after that show no changes.

**Generating configuration:** To avoid writing the resource blocks by hand, add only the `import` blocks and let Terraform generate the resource blocks:

```bash
terraform plan -generate-config-out=generated.tf
```

> [!IMPORTANT]
> Generated configuration is only a starting point. It can include conflicting or read-only attributes that make `terraform plan` fail, so review and fix `generated.tf` before you apply. For more information, see [Generating Configuration](https://developer.hashicorp.com/terraform/language/import/generating-configuration#limitations).

## Migrating from the BloxOne Provider

### Provider Configuration

**BloxOne provider:**

```terraform
terraform {
  required_providers {
    bloxone = {
      source = "infobloxopen/bloxone"
    }
  }
}

provider "bloxone" {
  csp_url = "<INFOBLOX_PORTAL_URL>"
  api_key = "<API_KEY>"

  default_tags = {
    managed_by = "terraform"
  }
}
```

**Unified provider:**

```terraform
terraform {
  required_providers {
    infoblox = {
      source  = "infobloxopen/infoblox"
      version = ">= 3.0.0"
    }
  }
}

provider "infoblox" {
  uddi = {
    portal_url = "<INFOBLOX_PORTAL_URL>"
    portal_key = "<INFOBLOX_PORTAL_KEY>"

    default_tags = {
      managed_by = "terraform"
    }
  }
}
```

| BloxOne provider | Unified provider |
|------------------|------------------|
| `csp_url` | `uddi.portal_url` |
| `api_key` | `uddi.portal_key` |
| `default_tags` | `uddi.default_tags` |

### Attribute Changes

Attribute names are the same as in the BloxOne provider. They move into the `uddi` block.

**BloxOne provider:**

```terraform
resource "bloxone_dns_a_record" "example" {
  zone         = bloxone_dns_auth_zone.example.id
  name_in_zone = "web"
  rdata = {
    address = "10.0.0.18"
  }
  tags = {
    Site = "location-1"
  }
}
```

**Unified provider:**

```terraform
resource "infoblox_record_a" "example" {
  uddi = {
    zone         = infoblox_zone_auth.example.id
    name_in_zone = "web"
    rdata = {
      address = "10.0.0.18"
    }
    tags = {
      Site = "location-1"
    }
  }
}
```

* References to other resources change from `bloxone_<type>.<name>.id` to `infoblox_<type>.<name>.id`. The `id` stays at the top level and is not part of the `uddi` block.
* Some read-only attributes of the BloxOne provider are not available. Check the documentation for each resource before you reference a computed attribute.



## Migrating from the Legacy Infoblox Provider

The legacy provider (v2.x) and the unified provider (v3.0.0 and later) use the same registry source, `infobloxopen/infoblox`. 

> [!NOTE]
> Some resource types, such as `infoblox_zone_auth` and `infoblox_network_view`, have the same name in both versions but a different schema. Remove **all** legacy resources from the state before you change the provider version. Otherwise, Terraform cannot read the existing state after the upgrade.

### Provider Configuration

**Legacy provider:**

```terraform
terraform {
  required_providers {
    infoblox = {
      source  = "infobloxopen/infoblox"
      version = "~> 2.0"
    }
  }
}

provider "infoblox" {
  server       = "<NIOS_GRID_IP>"
  username     = "<NIOS_USERNAME>"
  password     = "<NIOS_PASSWORD>"
  wapi_version = "2.12.3"
}
```

**Unified provider:**

```terraform
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

| Legacy provider | Unified provider |
|-----------------|------------------|
| `server` | `nios.host_url`, as a full URL such as `https://<NIOS_GRID_IP>` |
| `username` | `nios.username` |
| `password` | `nios.password` |

### Attribute Changes

The unified provider uses attribute names that match the NIOS WAPI field names, and the attributes move into the `nios` block.

**Legacy provider:**

```terraform
resource "infoblox_a_record" "example" {
  fqdn     = "a-record.example.com"
  ip_addr  = "10.0.0.1"
  ttl      = 300
  dns_view = "default"
  comment  = "A record created by Terraform"
  ext_attrs = jsonencode({
    "Site" = "location-1"
  })
}
```

**Unified provider:**

```terraform
resource "infoblox_record_a" "example" {
  nios = {
    name     = "a-record.example.com"
    ipv4addr = "10.0.0.1"
    ttl      = 300
    view     = "default"
    comment  = "A record created by Terraform"
    ext_attrs = {
      Site = "location-1"
    }
  }
}
```

Key changes:

* `fqdn` → `name`
* `ip_addr` → `ipv4addr`
* `dns_view` → `view`
* `ext_attrs` is a map instead of a JSON-encoded string, so `jsonencode()` is no longer needed.
* Repeated nested blocks are replaced by list attributes. For example, `auth_zones { ... }` blocks in `infoblox_dtc_lbdn` become `auth_zones = [ ... ]`.
* The `internal_id` and `ref` attributes are removed. The object reference is available as `id`.


## Migrating from the NIOS Provider

### Provider Configuration

**NIOS provider:**

```terraform
terraform {
  required_providers {
    nios = {
      source = "infobloxopen/nios"
    }
  }
}

provider "nios" {
  nios_host_url = "https://<NIOS_GRID_IP>"
  nios_username = "<NIOS_USERNAME>"
  nios_password = "<NIOS_PASSWORD>"
}
```

**Unified provider:**

```terraform
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

| NIOS provider | Unified provider |
|---------------|------------------|
| `nios_host_url` | `nios.host_url` |
| `nios_username` | `nios.username` |
| `nios_password` | `nios.password` |
| `manage_internal_id_ea` | `manage_internal_id_ea` |
| `retry_timeout` | `operation_timeout` |

### Attribute Changes

Attribute names already match the NIOS WAPI field names, so most of them stay the same. They move into the `nios` block.

**NIOS provider:**

```terraform
resource "nios_dns_record_a" "example" {
  name     = "a-record.example.com"
  ipv4addr = "10.0.0.1"
  view     = "default"
  extattrs = {
    Site = "location-1"
  }
}
```

**Unified provider:**

```terraform
resource "infoblox_record_a" "example" {
  nios = {
    name     = "a-record.example.com"
    ipv4addr = "10.0.0.1"
    view     = "default"
    ext_attrs = {
      Site = "location-1"
    }
  }
}
```

Key changes:

* `extattrs` → `ext_attrs`
* `func_call` → `dynamic_allocation`, which allocates the next available IP address from a network
* `ref` → `id`. References such as `nios_dns_zone_auth.example.ref` become `infoblox_zone_auth.example.id`.
* Some read-only attributes of the NIOS provider are not available. Check the documentation for each resource before you reference a computed attribute.
