---
page_title: "Quickstart"
subcategory: "Guides"
description: |-
  Configure the Infoblox provider against the NIOS or UDDI backend and create your first DNS and IPAM resources.
---

# Managing DDI Services with the Infoblox Terraform Provider

This guide provides step-by-step instructions for configuring the Infoblox provider and managing DDI resources with it.

The provider is unified: the same provider works with both Infoblox backends. A configuration targets one backend at a time, and each resource is configured in the nested block named after that backend.

| Backend | Description | Resource block |
|---|---|---|
| NIOS | An on-premise Infoblox Grid, accessed over WAPI | `nios = { ... }` |
| UDDI | Infoblox Universal DDI, accessed through the Infoblox Portal | `uddi = { ... }` |

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 1.12.1 or later
- One of:
  - Infoblox NIOS 9.0.6 (WAPI v2.13.6)
  - An Infoblox Portal account

## Configuring the Provider

Before getting started, ensure you have completed the [prerequisites](../../README.md#prerequisites).

Create a directory for the Terraform configuration and add a file named `main.tf`.

### NIOS Backend

````terraform
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
    host_url = "<NIOS_HOST_URL>"
    username = "<NIOS_USERNAME>"
    password = "<NIOS_PASSWORD>"
  }
}
````

> ⚠️ **Warning** : Hard-coded credentials are not recommended in any configuration file. It is recommended to use variables or environment variables.

These values may also be supplied through the `NIOS_HOST_URL`, `NIOS_USERNAME`, and `NIOS_PASSWORD` environment variables.

### UDDI Backend

````terraform
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
  }
}
````

These values may also be supplied through the `INFOBLOX_PORTAL_URL` and `INFOBLOX_PORTAL_KEY` environment variables.

> **Note:** The provider accepts either a `nios` block or a `uddi` block, not both. To manage both backends in the same configuration, declare two provider instances with aliases and set the `provider` argument on each resource.

Initialize the working directory. This downloads the provider:

```shell
terraform init
```

## Creating Resources

### DNS Resources

Create an authoritative zone, and an A record and a CNAME record within it.

````terraform
// Create an Auth Zone
resource "infoblox_zone_auth" "example_zone" {
  nios = {
    fqdn = "example.com"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create an A Record
resource "infoblox_record_a" "example_a" {
  nios = {
    name     = "a-record.${infoblox_zone_auth.example_zone.nios.fqdn}"
    ipv4addr = "10.0.0.10"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create a CNAME Record
resource "infoblox_record_cname" "example_cname" {
  nios = {
    name      = "cname-record.${infoblox_zone_auth.example_zone.nios.fqdn}"
    canonical = "a-record.${infoblox_zone_auth.example_zone.nios.fqdn}"
  }
}
````

### IPAM Resources

Create a Network View and a Network within it.

````terraform
// Create a Network View
resource "infoblox_network_view" "example_view" {
  nios = {
    name = "example_network_view"
  }
}

// Create an IPv4 Network
resource "infoblox_network" "example_network" {
  nios = {
    network      = "15.0.0.0/24"
    network_view = infoblox_network_view.example_view.nios.name
    comment      = "Created by Terraform"
    ext_attrs = {
      Site = "location-1"
    }
  }
}
````

### The Same Objects on UDDI

The resource types are the same, the fields are the ones UDDI defines, and they go in the `uddi` block.

Create a Network View and a Network within it.

````terraform
// Create a Network View
resource "infoblox_network_view" "example_network_view" {
  uddi = {
    name = "example_network_view"
  }
}

// Create an IPv4 Network
resource "infoblox_network" "example_network" {
  uddi = {
    address = "15.0.0.0"
    cidr    = 24
    space   = infoblox_network_view.example_network_view.id
    comment = "Created by Terraform"
    tags = {
      Site = "location-1"
    }
  }
}

// Create an Auth Zone
resource "infoblox_zone_auth" "example_zone" {
  uddi = {
    fqdn         = "example.com."
    primary_type = "cloud"
  }
}

// Create an A Record
resource "infoblox_record_a" "example_record_a" {
  uddi = {
    name_in_zone = "a-record"
    zone         = infoblox_zone_auth.example_zone.id
    rdata = {
      address = "10.0.0.10"
    }
  }
}
````

You can now run `terraform plan` to see what resources will be created.

```shell
terraform plan
```

## Applying the Configuration

To create the resources, run the following command:

```shell
terraform apply
```

## Destroying the Configuration

To destroy all the resources, run the following command:

```shell
terraform destroy
```

## Reading Existing Objects

Data sources retrieve objects that already exist.

````terraform
// Retrieve an Auth Zone
data "infoblox_zone_auth" "example" {
  filters = {
    view = "default"
    fqdn = "example.com"
  }
}

output "zone_info" {
  value = data.infoblox_zone_auth.example
}
````

Use `filters` for field matches, `ext_attr_filters` for extensible attributes on NIOS, and `tag_filters` for tags on UDDI.

Every resource also has a matching list resource, which finds existing objects and is run with `terraform query` rather than `terraform plan`.

## Next Steps

You can also use the Infoblox Terraform Provider to manage other resources. For more information, see the [Infoblox Terraform Provider documentation](https://registry.terraform.io/providers/infobloxopen/infoblox/latest/docs).