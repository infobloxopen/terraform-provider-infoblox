---
page_title: "NIOS Service Account Permissions"
subcategory: "Guides"
description: |-
  Grant a dedicated, non-superuser NIOS admin account the permissions it needs to manage objects with the Infoblox provider.
---

# NIOS Service Account Permissions

The provider can connect to NIOS with a dedicated admin account that is not a superuser. The account belongs to an admin group, and the permissions of that group decide which objects Terraform can read and change. This guide lists the permissions each common resource needs and shows how to set up the account.

This guide applies to NIOS only. Access through the `uddi` block is controlled by the Infoblox Portal API key.

A NIOS permission has three parts:

| Part | Values | Example |
|---|---|---|
| Resource type | The kind of object the permission covers | `NETWORK`, `ZONE`, `HOST` |
| Permission | `READ` or `WRITE`. `WRITE` includes `READ`. | `WRITE` |
| Object | The parent object the permission applies to. The permission covers the matching objects inside it. | A network view, a DNS view or a zone |

When a permission is missing, NIOS rejects the request with an error that names it, for example `Write permission for the network '10.0.0.0/24' in the network view 'default' is required for this operation.`

## Requirements

- **The Terraform Internal ID extensible attribute must exist.** A superuser creates the definition once, before the service account is used. After that, the service account needs no permission on extensible attribute definitions. For the steps, see [Management of Provider Operations with Terraform Internal ID](tf-internal-id-management.md#creating-the-terraform-internal-id-extensible-attribute).
- **The admin group must allow API access.** Set `access_method` to include `API`. Access to Grid Manager is not needed.

## Permissions by Resource

Grant the narrowest scope that covers the objects Terraform manages.

| Resource | Resource type and permission | Object |
|---|---|---|
| `infoblox_network` | `NETWORK`, `WRITE` | Network view |
| `infoblox_network_container` | `NETWORK`, `WRITE` | Network view |
| `infoblox_range` | `RANGE`, `WRITE` | Parent network. `NETWORK`, `WRITE` on the network view also covers it. |
| `infoblox_fixed_address` | `FIXED_ADDRESS`, `WRITE` | Parent network. `NETWORK`, `WRITE` on the network view also covers it. |
| `infoblox_zone_auth` | `ZONE`, `WRITE` | DNS view |
| `infoblox_record_a` | `A`, `WRITE` | Zone or DNS view |
| `infoblox_record_aaaa` | `AAAA`, `WRITE` | Zone or DNS view |
| `infoblox_record_cname` | `CNAME`, `WRITE` | Zone or DNS view |
| `infoblox_record_ptr` | `PTR`, `WRITE` | Reverse zone or DNS view |
| `infoblox_record_host` | `HOST`, `WRITE` | Forward zone and reverse zone of the address, or DNS view |

- **`NETWORK` is the only resource type that can be granted on a network view.** It covers the networks, network containers, ranges and fixed addresses in that view.
- **A host record also needs its reverse zone.** NIOS writes the matching PTR entry when a reverse zone exists for the address, so `HOST`, `WRITE` is needed on both zones. A permission on the DNS view covers both.
- **Dynamic allocation needs `READ` on the network.** An `infoblox_record_host` address that uses `dynamic_allocation` also needs `READ` on the network it allocates from, or `NETWORK`, `READ` on the network view. Without it, NIOS reports that the network was not found.

## Data Sources

Data sources need `READ` on the resource type, at the same scope as the matching resource. For example, `data "infoblox_record_host"` needs `HOST`, `READ` on the zone or DNS view.

When the account cannot read an object, NIOS leaves the object out of the results instead of returning an error, so the data source returns no results. If a configuration uses a lookup to decide whether to create an object, a missing `READ` permission can lead to duplicate objects. Grant `READ` on everything the configuration looks up.

## Deleting Zones and Network Containers

NIOS deletes authoritative zones and network containers recursively, even when they are empty. The Grid setting **Allow recursive deletion** (`allow_recursive_deletion` on the WAPI `grid` object) decides who can delete them:

| Value | Who can delete zones and network containers |
|---|---|
| `ALL` | Any admin with `WRITE` permission on them |
| `SUPERUSERS` | Superusers only |
| `NOBODY` | Nobody |

With `SUPERUSERS`, `terraform destroy` of these resources fails for the service account with `You have no permission to delete this object recursively`. To let the service account delete them, a Grid administrator sets **Allow recursive deletion** to `ALL`. Otherwise, delete zones and network containers with a superuser account.

## Creating the Service Account

Run the following steps with a superuser account. The examples create an admin group named `terraform-sa` with access to the `default` network view and the `default` DNS view.

### Creating the Admin Group

```bash
curl -k -u <SUPERUSER>:<PASSWORD> -H "Content-Type: application/json" -X POST https://<NIOS_GRID_IP>/wapi/<WAPI_VERSION>/admingroup -d '{"name": "terraform-sa", "access_method": ["API"], "comment": "Terraform service account"}'
```

### Adding Permissions

Look up the reference of each object to grant permissions on:

```bash
curl -k -u <SUPERUSER>:<PASSWORD> "https://<NIOS_GRID_IP>/wapi/<WAPI_VERSION>/networkview?name=default"
curl -k -u <SUPERUSER>:<PASSWORD> "https://<NIOS_GRID_IP>/wapi/<WAPI_VERSION>/view?name=default"
```

Add one permission for each resource type. The following commands grant `NETWORK` on the network view, and `ZONE` and the record types on the DNS view:

```bash
curl -k -u <SUPERUSER>:<PASSWORD> -H "Content-Type: application/json" -X POST https://<NIOS_GRID_IP>/wapi/<WAPI_VERSION>/permission -d '{"group": "terraform-sa", "permission": "WRITE", "resource_type": "NETWORK", "object": "<NETWORK_VIEW_REF>"}'

for RT in ZONE HOST A AAAA CNAME PTR; do
  curl -k -u <SUPERUSER>:<PASSWORD> -H "Content-Type: application/json" -X POST https://<NIOS_GRID_IP>/wapi/<WAPI_VERSION>/permission -d '{"group": "terraform-sa", "permission": "WRITE", "resource_type": "'"$RT"'", "object": "<DNS_VIEW_REF>"}'
done
```

To limit a permission to a single zone or network, use the reference of that object in `object`.

You can also create the admin group and its permissions in Grid Manager, on the **Administration** tab > **Administrators** tab. For steps, refer to the Managing Administrators topic in the [Infoblox NIOS Documentation](https://docs.infoblox.com/).

### Creating the Admin User

Create the user in the admin group with the `infoblox_admin_user` resource, using a provider configured with superuser credentials:

````terraform
resource "infoblox_admin_user" "terraform_sa" {
  nios = {
    name         = "terraform"
    password     = var.terraform_sa_password
    admin_groups = ["terraform-sa"]
    comment      = "Terraform service account"
  }
}
````

### Configuring the Provider

Set the credentials of the service account in the `nios` block:

````terraform
provider "infoblox" {
  nios = {
    host_url = "<NIOS_HOST_URL>"
    username = "terraform"
    password = var.terraform_sa_password
  }
}
````

## Troubleshooting

| Error | Permission to add |
|---|---|
| `Write permission for the network '...' in the network view '...' is required` | `NETWORK`, `WRITE` on the network view |
| `Write permission for Reserved range '...' in network view '...' is required` | `RANGE`, `WRITE` on the parent network |
| `Write permission for fixed address '...' in network view '...' is required` | `FIXED_ADDRESS`, `WRITE` on the parent network |
| `Write permission for zone '...' in the DNS view '...' is required` | `ZONE`, `WRITE` on the DNS view |
| `Write permission for the resource record type 'HOST' in the zone '...' is required` | `HOST`, `WRITE` on the zone in the message. For a reverse zone such as `10.0.0.0/24`, add it on that zone. |
| `The search parameters {"network": ...} for object network did not return any result` | `READ` on the network used by `dynamic_allocation` |
| `You have no permission to delete this object recursively` | See [Deleting Zones and Network Containers](#deleting-zones-and-network-containers) |
| A data source returns no results for an object that exists | `READ` on the resource type |

## Reference

For more information about admin groups, permissions and the Grid security settings, refer to the [Infoblox NIOS Documentation](https://docs.infoblox.com/). For the `admingroup` and `permission` objects, refer to the [Infoblox WAPI documentation](https://docs.infoblox.com/space/NIOS/35400616/NIOS).
