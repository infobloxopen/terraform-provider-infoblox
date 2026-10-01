---
page_title: "Host Record Management"
subcategory: "Guides"
description: |-
  Allocate IP addresses and associate DHCP settings on a NIOS host record using infoblox_record_host and infoblox_ip_association.
---

# Host Record Management

A NIOS host record binds a hostname to one or more IP addresses and can carry both DNS and DHCP configuration for them. The provider manages it through two resources, which correspond to the two operations performed against a host record:

| Operation | Resource | Purpose |
|---|---|---|
| Allocation | `infoblox_record_host` | Creates the host record and its addresses, either static or allocated from a network, and manages the DNS configuration. The address is marked as used in NIOS. |
| Association | `infoblox_ip_association` | Attaches DHCP identifiers - a MAC address for IPv4, a DUID for IPv6 - to the addresses of an existing host record. |

`infoblox_record_host` owns the lifecycle of the host record. `infoblox_ip_association` never creates or destroys one, it only updates the DHCP settings of a host record that already exists.

Both resources are NIOS-only and are configured in the `nios = { ... }` block.

## Allocating an Address

An address is either static, set through `ipv4addr` or `ipv6addr`, or allocated dynamically through `dynamic_allocation`. The two are mutually exclusive within an address entry.

````terraform
// Static address
resource "infoblox_record_host" "static" {
  nios = {
    name              = "host1.example.com"
    view              = "default"
    configure_for_dns = true
    ipv4addrs = [
      {
        ipv4addr = "10.0.0.20"
      }
    ]
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Next available address from a network
resource "infoblox_record_host" "dynamic" {
  nios = {
    name              = "host2.example.com"
    view              = "default"
    configure_for_dns = true
    ipv4addrs = [
      {
        dynamic_allocation = {
          network      = "10.0.0.0/24"
          network_view = "default"
          exclude      = ["10.0.0.1", "10.0.0.2"]
        }
      }
    ]
  }
}
````

Within `dynamic_allocation`, use either `network` to name the network directly or `filter_params` to select one by extensible attribute. These are mutually exclusive.

## Associating DHCP Settings

`infoblox_ip_association` references the host record through `ref`, which takes the `id` of the `infoblox_record_host` resource.

````terraform
resource "infoblox_record_host" "dual_stack" {
  nios = {
    name              = "host3.example.com"
    view              = "default"
    configure_for_dns = true
    ipv4addrs = [
      {
        ipv4addr = "10.0.0.21"
      }
    ]
    ipv6addrs = [
      {
        ipv6addr = "2002:1f93::12:2"
      }
    ]
  }
}

resource "infoblox_ip_association" "dual_stack" {
  nios = {
    ref                = infoblox_record_host.dual_stack.id
    mac                = "aa:bb:cc:11:22:44"
    duid               = "00:03:00:01:aa:bb:cc:11:22:44"
    match_client       = "DUID"
    configure_for_dhcp = true
  }
}
````

| Attribute | Notes |
|---|---|
| `ref` | Reference to the host record, normally `infoblox_record_host.<name>.id`. Required. |
| `mac` | MAC address the IPv4 address is leased to. |
| `duid` | DUID the IPv6 address is leased to. |
| `match_client` | Identifier the IPv6 address is leased to: `DUID` matches the DUID, `MAC_ADDRESS` matches the MAC address. Defaults to `DUID`. |
| `configure_for_dhcp` | Enables the DHCP configuration for the associated addresses. When `true`, at least one of `mac` or `duid` is required. |

## Lifecycle and Behavior

- **Create the allocation first.** The host record must exist before the association is applied. If the host record cannot be found, the association fails with a message directing you to create `infoblox_record_host` first.
- **Destroying the association clears the DHCP settings** - `mac`, `duid`, and `configure_for_dhcp` - and leaves the host record and its addresses in place. To remove the host record itself, destroy `infoblox_record_host`, which removes the DHCP settings along with it.
- **On import,** import the `infoblox_record_host` resource before the `infoblox_ip_association` that references it.
