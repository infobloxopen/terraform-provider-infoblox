# DTC LBDN Data Source

Use the `infoblox_dtc_lbdn` data source to retrieve the following information for a DTC LBDN if any, which is managed by a NIOS server:

* `name`: The name of th DTC LBDN. Example: `testLbdn`.
* `auth_zones`: List of linked auth zones with their respective views. `auth_zones` has the following two fields `fqdn` and `dns_view`. Example:
```terraform
auth_zones {
    fqdn = "example.com"
    dns_view = "default"
  }
```
* `fqdn`: The name of the auth-zone to link with. Example: `example.com`.
* `dns_view`: The DNS view on which the auth-zone is available. Example: `default`.
* `auto_consolidated_monitors`: Flag for enabling auto managing DTC Consolidated Monitors on related DTC Pools. Example: `false`.
* `disable`: Flag to determine whether the DTC LBDN is disabled or not. When this is set to False, the fixed address is enabled. Example: `false`.
* `lb_method`: The load balancing method. Used to select pool. Valid values are GLOBAL_AVAILABILITY, RATIO, ROUND_ROBIN, SOURCE_IP_HASH and TOPOLOGY. Example: `ROUND_ROBIN`.
* `pools`: Pools associated with an LBDN are collections of load-balanced servers. `pools` has the following two fields `pool` and `ratio`. The description of the fields of `pools` is as follows:
  
  `pool`: The name of the pool. Example: `pool1`.
  
  `ratio`: The weight of the pool. Example: `2`.
```terraform
pools {
    pool = "pool1"
    ratio = "2"
  }
```
* `Patterns`: LBDN wildcards for pattern match. Example: `["*.example.com","*test.com"]`.
* `persistence`: Maximum time, in seconds, for which client specific LBDN responses will be cached. Zero specifies no caching. Example: `60`.
* `priority`: The LBDN pattern match priority for overlapping DTC LBDN objects. Example: `1`.
* `ttl`: The Time To Live (TTL) value for the DTC LBDN. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. Example: `60`.
* `topology`: The topology rules for TOPOLOGY method. Example: `test-topo`.
* `types`: The list of resource record types supported by LBDN. Example: `["A","AAAA","CNAME","NAPTR","SRV"]`.
* `health`: The LBDN health information. The description of the fields of `health` is as follows:

  `availability`: The availability color status. Default value: `NONE`. Valid values are one of these: `BLUE`, `GREEN`, `GRAY`, `NONE`, `RED` and `YELLOW`.

  `description`: The textual description of the LBDN object’s status. Default value: `""`. Example: `test health`.

  `enabled_state`: The enabled state of the LBDN. Default value: `ENABLED`. Valid values are one of these: `DISABLED`, `DISABLED_BY_PARENT`, `ENABLED` and `NONE`.
```terraform
health { 
  availability = "NONE"
  description = ""
  enabled_state = "DISABLED"
}
```
* `comment`: The description of the DTC LBDN. This is a regular comment. Example: `test LBDN`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as string of JSON map. Example: `"{\"*Site\":\"Antarctica\"}"`

For usage of filters, add the fields as keys and appropriate values to be passed to the keys like `name`, `comment`, `fqdn` and `status_member` corresponding to object.
From the below list of supported arguments for filters,  use only the searchable fields for retrieving the matching records.

### Supported Arguments for filters

-----
| Field   | Alias         | Type   | Searchable |
|---------|---------------|--------|------------|
| name    | name          | string | yes        |
| comment | comment       | string | yes        |

These fields are used only for searching. They are not actual members of the object and therefore the server does not return 
these fields with this name unless they are nested return fields.
-----
| Field   | Alias         | Type   | Searchable |
|---------|---------------|--------|------------|
| -       | fqdn          | string | yes        |
| -       | status_member | string | yes        |

!> Any of the combination from searchable fields in supported arguments list for fields are allowed.

!> Please consider using only fields as the keys in terraform datasource filters, kindly don't use alias names as keys from the above table.

### Example for using the filters:
```hcl
data "infoblox_dtc_lbdn" "lbdn_filter" {   
  filters = {
    name = "test-lbdn"
    comment = "sample LBDN"
  }
}
```

```hcl
data "infoblox_dtc_lbdn" "lbdn_filter" {   
  filters = {
    fqdn = "test.com"
    status_member = "infoblox.localdomain"
  }
}
```

!> If `null` or empty filters are passed, then all the objects associated with datasource like here `infoblox_dtc_lbdn` will be fetched in results.

### Example of DTC LBDN Data Source Block

```hcl
resource "infoblox_dtc_lbdn" "lbdn_record" {
  name = "testLbdn"
  lb_method = "ROUND_ROBIN"
  comment = "test LBDN"
  topology = "test-topo"
  types = ["A","AAAA","CNAME","NAPTR","SRV"]
  auth_zones {
    fqdn = "example.com"
    dns_view = "default"
  }
  ext_attrs = jsonencode({
    "Site" = "Antarctica"
  })
}

data "infoblox_dtc_lbdn" "lbdn_read" {
  filters = {
    name = infoblox_dtc_lbdn.lbdn_record.name
    comment = infoblox_dtc_lbdn.lbdn_record.comment
    fqdn = "example.com"
    status_member = "infoblox.localdomain"
  }
}

output "lbdn_res" {
  value = data.infoblox_dtc_lbdn.lbdn_read
}

// accessing individual field in results
output "lbdn_name" {
  value = data.infoblox_dtc_lbdn.lbdn_res.results.0.name //zero represents index of json object from results list
}

// accessing DTC LBDN through EA's
data "infoblox_dtc_lbdn" "lbdn_ea" {
  filters = {
    "*Site" = "Antarctica"
  }
}

output "lbdn_out" {
  value = data.infoblox_dtc_lbdn.lbdn_ea
}
```