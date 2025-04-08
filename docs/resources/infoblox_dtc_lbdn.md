# DTC LBDN Resource

The `infoblox_dtc_lbdn` resource enables you to perform `create`, `update` and `delete` operations on DTC LBDN in a NIOS appliance.
The resource represents the ‘dtc:lbdn’ WAPI object in NIOS.

The following list describes the parameters you can define in the resource block of the DTC LBDN object:

- `name`: required, specifies the display name of the DTC LBDN. Example: `test-lbdn`.
- `lb_method`: required, specifies the load balancing method. Used to select pool. Example: `ROUND_ROBIN`. Valid values are `"GLOBAL_AVAILABILITY", "RATIO", "ROUND_ROBIN", "SOURCE_IP_HASH", "TOPOLOGY"`.
- `auto_consolidated_monitors`: optional, specifies the flag for enabling auto managing DTC Consolidated Monitors on related DTC Pools. Default value: `false`.
- `topology`: optional, specifies the topology rules for TOPOLOGY method. Example: `test-topo`.
- `disable`: optional, specifies whether the DTC LBDN is disabled or not. When this is set to False, the fixed address is enabled. Default value: `false`.
- `patterns`: optional, LBDN wildcards for pattern match. Example: `["*.example.com", "*test.com"]`.
- `persistence`: optional, specifies the maximum time, in seconds, for which client specific LBDN responses will be cached. Zero specifies no caching. Default value: `0`.
- `priority`: optional, specifies the LBDN pattern match priority for “overlapping” DTC LBDN objects. LBDNs are “overlapping” if they are simultaneously assigned to a zone and have patterns
  that can match the same FQDN. The matching LBDN with highest priority (lowest ordinal) will be used. Default value: `1`.
- `ttl`: optional, specifies the Time To Live (TTL) value for the DTC LBDN. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached).
  Zero indicates that the record should not be cached. Example: `60`
- `types`: required, specifies the list of resource record types supported by LBDN. Valid values are: `"A", "AAAA", "CNAME", "NAPTR", "SRV"`. Atleast one of the valid values must be given.
- `pools`: optional, specifies the pools associated with the LBDN. `pools` has the following two fields `pool` and `ratio`.The description of the fields of `pools` is as follows:
  - `pool`: required, specifies the name of the pool. Example: `pool1`.
  - `ratio`: required, specifies the weight of the pool. Example: `2`.

Example for `pools`:

```hcl
pools {
  pool  = "pool1"
  ratio = "2"
}
```

- `auth_zones`: optional, specifies the list of linked auth zones. `auth_zones` has the following two fields `fqdn` and `dns_view`. The description of the fields of `auth_zones` is as follows:
  - `fqdn`: required, specifies the name of the auth-zone to link with. Example: `example.com`.
  - `dns_view`: required, specifies the DNS view on which the auth-zone is available. Example: `default`.

Example for `auth_zones`:

```hcl
auth_zones {
  fqdn     = "example.com"
  dns_view = "default"
}
```

- `comment`: optional, description of the DTC LBDN. Example: `custom DTC LBDN`.
- `ext_attrs`: optional, set of the Extensible attributes of the LBDN, as a map in JSON format. Example: `jsonencode({})`.

### Examples of a DTC LBDN Block

```hcl
// creating a DTC LBDN record with minimal set of parameters
resource "infoblox_dtc_lbdn" "lbdn_minimal_parameters" {
  name      = "testLbdn2"
  lb_method = "ROUND_ROBIN"
  topology  = "test-topo"
  types     = ["A"]
}

// creating a DTC LBDN record with full set of parameters
resource "infoblox_dtc_lbdn" "lbdn_full_set_parameters" {
  name = "testLbdn11"

  auth_zones {
    fqdn     = "test.com"
    dns_view = "default"
  }
  auth_zones {
    fqdn     = "test.com"
    dns_view = "default.dnsview1"
  }
  auth_zones {
    fqdn     = "info.com"
    dns_view = "default.view2"
  }

  comment = "test"

  ext_attrs = jsonencode({
    "Location" = "65.8665701230204, -37.00791763398113"
  })

  lb_method = "TOPOLOGY"
  patterns  = ["test.com", "info.com*"]

  pools {
    pool  = "pool2"
    ratio = 2
  }
  pools {
    pool  = "rrpool"
    ratio = 3
  }
  pools {
    pool  = "test-pool"
    ratio = 6
  }

  ttl         = 120
  topology    = "test-topo"
  disable     = true
  types       = ["A", "AAAA", "CNAME"]
  persistence = 60
  priority    = 1
}
```
