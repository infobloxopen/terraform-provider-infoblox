# DTC LBDN Data Source

Use the `infoblox_dtc_lbdn` data source to retrieve the following information for a DTC LBDN if any, which is managed by a NIOS server:

* `name`: The name of th DTC LBDN. Example: `test-lbdn`.
* `auth_zones`: List of linked auth zones with their respective views.
* `auto_consolidated_monitors`: Flag for enabling auto managing DTC Consolidated Monitors on related DTC Pools.
* `disable`: Flag to determine whether the DTC LBDN is disabled or not. When this is set to False, the fixed address is enabled.
* `lb_method`: The load balancing method. Used to select pool. Valid values are GLOBAL_AVAILABILITY, RATIO, ROUND_ROBIN, SOURCE_IP_HASH and TOPOLOGY.
* `pools`: Pools associated with an LBDN are collections of load-balanced servers.
* `Patterns`: LBDN wildcards for pattern match.
* `persistence`: Maximum time, in seconds, for which client specific LBDN responses will be cached. Zero specifies no caching.
* `priority`: The LBDN pattern match priority for overlapping DTC LBDN objects.
* `ttl`: The Time To Live (TTL) value for the DTC LBDN. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached.
* `topology`: The topology rules for TOPOLOGY method.
* `types`: The list of resource record types supported by LBDN. Valid values are `"A", "AAAA", "CNAME", "NAPTR", "SRV"`. Default value is `"A"` and `"AAAA"`.
* `health`: The LBDN health information.
* `comment`: The description of the DTC LBDN. This is a regular comment. Example `this is some text`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as string of JSON map. Example: `"{\"Site\":\"Kapu\"}"`.

For usage of filters, add the fields as keys and appropriate values to be passed to the keys like `name`, `comment` corresponding to object.
From the below list of supported arguments for filters,  use only the searchable fields for retrieving the matching records.

### Supported Arguments for filters

-----
| Field   | Alias   | Type   | Searchable |
|---------|---------|--------|------------|
| name    | name    | string | yes        |
| comment | comment | string | yes        |

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

!> If `null` or empty filters are passed, then all the objects associated with datasource like here `infoblox_dtc_lbdn` will be fetched in results.

### Example of DTC LBDN Data Source Block

```hcl
resource "infoblox_dtc_lbdn" "lbdn" {
  name = "testLbdn"
  lb_method = "ROUND_ROBIN"
  comment = "test LBDN"
  topology = "test-topo"
  ext_attrs = jsonencode({
    "Site" = "Antarctica"
  })
}

data "infoblox_dtc_lbdn" "lbdn" {
  filters = {
    name = "testLbdn"
    comment = "test LBDN"
  }
  
  // This is just to ensure that the record has been be created
  // using 'infoblox_dtc_lbdn' resource block before the data source will be queried.
  depends_on = [infoblox_dtc_lbdn.lbdn]
}

output "lbdn_res" {
  value = data.infoblox_dtc_lbdn.lbdn
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