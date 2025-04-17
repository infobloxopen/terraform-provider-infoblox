# Resource Zone Delegated

The `infoblox_zone_delegated` resource enables you to perform the create, update, and delete operations on the delegated zones in a NIOS appliance. The resource represents the ‘zone_delegated’ WAPI object in NIOS.

A delegated zone must be a subzone of an authoritative zone.

The following list describes the parameters you can define in the `infoblox_zone_delegated` resource block:

- `fqdn`: required, specifies the name (in FQDN format) of the delegated DNS zone. For a reverse mapping zone, specify the IP address in CIDR format. For other zones, specify the value in FQDN format. This value can be in Unicode format.
  Example: `10.1.0.0/24` for reverse zone and `zone1.com` for forward zone.
- `view`: optional, specifies The name of the DNS view in which the zone resides. If value is not specified, `default` will be considered as default DNS view. Example: `external`.
- `zone_format`: optional, determines the format of corresponding zone. Valid values are `FORWARD`, `IPV4` and `IPV6`. Default value: `FORWARD`.
- `ns_group`: required if `delegate_to` field is not set, specifies the name server group that serves DNS for this zone. Example: `demoGroup`.
- `disable`: optional, specifies whether the zone is disabled. Default value: `false`.
- `delegated_ttl`: optional, specifies the TTL value for the delegated zone. The default value is `ttlUndef`.
- `comment`: optional, describes the delegated DNS zone. Example: `random delegated zone`.
- `ext_attrs`: optional, specifies the set of NIOS extensible attributes that will be attached to the delegated zone.
- `locked`: optional, determines whether the other administrators must be restricted from making conflicting changes.
  When you set this parameter to true, other administrators are restricted from making changes. The default value is false. Note that this flag is for administration purposes only. The zone will continue to serve DNS data even when it is locked.
- `delegate_to`: required if ns_group is not configured. Specifies the information of the remote name server that maintains the data for the delegated zone. Example:

```hcl
delegate_to {
  name = "te32.dz.ex.com"
  address = "10.0.0.1"
}
```

!> For a reverse zone, the corresponding 'zone_format' value should be set. And 'fqdn' once set cannot be updated.

> **Note**: Either define delegate_to or ns_group.

### Examples of a Zone Delegated Block

```hcl
// zone delegated, with fqdn and delegate_to
resource "infoblox_zone_delegated" "subdomain" {
  fqdn = "subdomain.example.com"

  delegate_to {
    name    = "ns-1488.awsdns-58.org"
    address = "10.1.1.1"
  }
  delegate_to {
    name    = "ns-2034.awsdns-62.co.uk"
    address = "10.10.1.1"
  }
}

// zone delegated, with fqdn and ns_group
resource "infoblox_zone_delegated" "zone_delegated2" {
  fqdn     = "min_params.ex.org"
  ns_group = "test"
}

// zone delegated with full set of parameters
resource "infoblox_zone_delegated" "zone_delegated3" {
  fqdn        = "max_params.ex.org"
  view        = "nondefault_view"
  zone_format = "FORWARD"
  comment     = "test sample delegated zone"

  delegate_to {
    name    = "te32.dz.ex.com"
    address = "10.0.0.1"
  }

  locked        = true
  delegated_ttl = 60

  ext_attrs = jsonencode({
    "Site" = "LA"
  })

  disable = true
}

// zone delegated IPV6 reverse mapping zone
resource "infoblox_zone_delegated" "zone_delegated4" {
  fqdn        = "3001:db8::/64"
  comment     = "zone delegated IPV6"
  zone_format = "IPV6"

  delegate_to {
    name = "test22.dz.ex.com"
    address = "10.0.0.1"
  }
}
```
