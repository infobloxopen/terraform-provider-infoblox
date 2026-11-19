---
page_title: "NIOS Service Account Permissions"
subcategory: "Guides"
description: |-
  The NIOS permissions a dedicated, non-superuser admin account needs to run the Infoblox provider. Applies to NIOS only.
---

# NIOS Service Account Permissions

> **Note:** This guide applies to **NIOS only**. Access through the `uddi` block is controlled by the Infoblox Portal API key.

You can run the provider with a dedicated NIOS admin account that is not a superuser. The account belongs to an admin group, and the group's permissions decide which objects Terraform can read and change. This guide lists the permissions each common resource and data source needs, and shows how to set up the account.

NIOS checks permissions on every request. When one is missing, the error names it, for example:

```
Write permission for the network '10.0.0.0/24' in the network view 'default' is required for this operation.
```

## Prerequisites

* **Terraform Internal ID extensible attribute.** A superuser must create the `Terraform Internal ID` extensible attribute definition once, before the service account is used. After that, the service account needs no permission on extensible attribute definitions. For the steps, see [Management of Provider Operations with Terraform Internal ID](tf-internal-id-management.md#creating-the-terraform-internal-id-extensible-attribute).
* **API access.** The admin group must allow API access (`access_method` includes `API`). GUI access is not needed.

## Permissions by Resource

Each permission is a resource type (such as `NETWORK` or `HOST`), an access level (`READ` or `WRITE`), and the object it applies to (such as a network view, DNS view or zone). A permission on a parent object applies to the matching objects inside it. `WRITE` includes `READ`.

Grant the narrowest scope that covers the objects Terraform manages.

| Resource | Permission | Scope |
|---|---|---|
| `infoblox_network` | `NETWORK` `WRITE` | Network view |
| `infoblox_network_container` | `NETWORK` `WRITE` | Network view |
| `infoblox_range` | `RANGE` `WRITE` | Parent network, or `NETWORK` `WRITE` on the network view |
| `infoblox_fixed_address` | `FIXED_ADDRESS` `WRITE` | Parent network, or `NETWORK` `WRITE` on the network view |
| `infoblox_zone_auth` | `ZONE` `WRITE` | DNS view |
| `infoblox_record_a` | `A` `WRITE` | Zone, or DNS view |
| `infoblox_record_aaaa` | `AAAA` `WRITE` | Zone, or DNS view |
| `infoblox_record_cname` | `CNAME` `WRITE` | Zone, or DNS view |
| `infoblox_record_ptr` | `PTR` `WRITE` | Reverse zone, or DNS view |
| `infoblox_record_host` | `HOST` `WRITE` | Forward zone **and** reverse zone of its address, or DNS view |
| `infoblox_record_host` with `dynamic_allocation` | Also `READ` on the network it allocates from | The network, or `NETWORK` `READ` on the network view |

Notes:
* In NIOS, `NETWORK` is the only resource type that can be granted on a network view. It covers networks, network containers, ranges and fixed addresses in that view.
* A host record also writes the matching reverse-mapping (PTR) entry, so it needs `HOST` `WRITE` on the reverse zone when one exists for its address. A permission on the DNS view covers both zones.
* Next-available allocation needs `READ` on the network. Without it, NIOS reports that the network was not found.

## Data Sources and Lookups

Data sources need `READ` on the resource type, at the same scope as in the table above. For example, `data "infoblox_record_host"` needs `HOST` `READ` on the zone or DNS view.

If the account cannot read an object, NIOS leaves it out of the results instead of returning an error. A lookup then returns no results, as if the object did not exist. If a configuration uses a lookup to decide whether to create an object, missing `READ` permission can lead to duplicates. Grant `READ` on everything the configuration looks up.

## Deleting Zones and Network Containers

Deleting an authoritative zone or a network container is a recursive operation in NIOS, even when the zone or container is empty. The Grid setting **Allow recursive deletion** (`allow_recursive_deletion` on the WAPI `grid` object) controls who can do it:

| Value | Who can delete zones and network containers |
|---|---|
| `ALL` | Any admin with `WRITE` permission on them |
| `SUPERUSERS` | Superusers only |
| `NOBODY` | Nobody |

With `SUPERUSERS`, `terraform destroy` of these resources fails for the service account with:

```
You have no permission to delete this object recursively
```

To let the service account delete them, a Grid administrator sets **Allow recursive deletion** to `ALL`. Otherwise, delete zones and network containers with a superuser account.

## Setting Up the Service Account

Run these steps as a superuser. The example creates a group named `terraform-sa` with access to the `default` network view and the `default` DNS view.

### 1. Create the admin group

```bash
curl -k -u <SUPERUSER>:<PASSWORD> -H "Content-Type: application/json" -X POST https://<NIOS_GRID_IP>/wapi/<WAPI_VERSION>/admingroup -d '{"name": "terraform-sa", "access_method": ["API"], "comment": "Terraform service account"}'
```

### 2. Add the permissions

Look up the reference of each object you want to grant permissions on:

```bash
curl -k -u <SUPERUSER>:<PASSWORD> "https://<NIOS_GRID_IP>/wapi/<WAPI_VERSION>/networkview?name=default"
curl -k -u <SUPERUSER>:<PASSWORD> "https://<NIOS_GRID_IP>/wapi/<WAPI_VERSION>/view?name=default"
```

Then add one permission for each resource type. For example, `NETWORK` on the network view, and `ZONE` and the record types on the DNS view:

```bash
curl -k -u <SUPERUSER>:<PASSWORD> -H "Content-Type: application/json" -X POST https://<NIOS_GRID_IP>/wapi/<WAPI_VERSION>/permission -d '{"group": "terraform-sa", "permission": "WRITE", "resource_type": "NETWORK", "object": "<NETWORK_VIEW_REF>"}'

for RT in ZONE HOST A AAAA CNAME PTR; do
  curl -k -u <SUPERUSER>:<PASSWORD> -H "Content-Type: application/json" -X POST https://<NIOS_GRID_IP>/wapi/<WAPI_VERSION>/permission -d '{"group": "terraform-sa", "permission": "WRITE", "resource_type": "'"$RT"'", "object": "<DNS_VIEW_REF>"}'
done
```

To scope a permission to a single zone or network instead, use that object's reference in `object`.

In Grid Manager, you can create the group and its permissions on the **Administration** tab > **Administrators** tab. For details, refer to the Managing Administrators topic in the [Infoblox NIOS Documentation](https://docs.infoblox.com/).

### 3. Create the admin user

Create the user in the group with the `infoblox_admin_user` resource, using a provider configured with superuser credentials:

```terraform
resource "infoblox_admin_user" "terraform_sa" {
  nios = {
    name         = "terraform"
    password     = var.terraform_sa_password
    admin_groups = ["terraform-sa"]
    comment      = "Terraform service account"
  }
}
```

### 4. Use the account

Configure the `nios` block with the service account's credentials:

```terraform
provider "infoblox" {
  nios = {
    host_url = "<NIOS_HOST_URL>"
    username = "terraform"
    password = "<SERVICE_ACCOUNT_PASSWORD>"
  }
}
```

## Troubleshooting

| Error | Permission to add |
|---|---|
| `Write permission for the network '...' in the network view '...' is required` | `NETWORK` `WRITE` on the network view |
| `Write permission for Reserved range '...' in network view '...' is required` | `RANGE` `WRITE` on the parent network |
| `Write permission for fixed address '...' in network view '...' is required` | `FIXED_ADDRESS` `WRITE` on the parent network |
| `Write permission for zone '...' in the DNS view '...' is required` | `ZONE` `WRITE` on the DNS view |
| `Write permission for the resource record type 'HOST' in the zone '...' is required` | `HOST` `WRITE` on the named zone. For a reverse zone such as `10.0.0.0/24`, add it on that zone as well. |
| `The search parameters {"network": ...} for object network did not return any result` (with `dynamic_allocation`) | `READ` on the network |
| `You have no permission to delete this object recursively` | See [Deleting Zones and Network Containers](#deleting-zones-and-network-containers) |
| A data source returns no results for an object that exists | `READ` on the object type |
