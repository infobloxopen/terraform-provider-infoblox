# DTC-Pool Data Source 

Use the `infoblox_dtc_pool` data source to retrieve the following information for an DTC Pool Record if any, which is managed by a NIOS server:

* `auto_consolidated_monitors`: Flag for enabling auto managing DTC Consolidated Monitors in DTC Pool. Default value: `false`
* `availability`: A resource in the pool is available if `ANY`, at least `QUORUM`, or `ALL` monitors for the pool say that it is up. Default value: `ALL`
* `comment`: The comment for the DTC Pool; maximum 256 characters. Example: `pool creation`
* `consolidated_monitors`: List of monitors and associated members statuses of which are shared across members and consolidated in server availability determination.

  `monitor_name`: name of the monitor. Example:` https`

  `montior_type`: Type of the monitor. Example: `http`

  `members`: Members whose monitor statuses are shared across other members in a pool. Example: `["infoblox.localdomain"]`

  `availability`: Servers assigned to a pool with monitor defined are healthy if ANY or ALL members report healthy status. Valid values are `ALL` and `ANY`

  `full_health_communication`: Flag for switching health performing and sharing behavior to perform health checks on each DTC grid member that serves related LBDN(s) and send them across all DTC grid members from both selected and non-selected lists.

```terraform
consolidated_monitors{
    monitor_name = "http"
    monitor_type = "http"
    members = ["infoblox.localdomain"]
    availability= "ALL"
    full_health_communication= true
  }
```
* `disable`: Determines whether the DTC Pool is disabled or not. When this is set to False, the fixed address is enabled. Default value: `false`
* `extattrs`: Extensible attributes associated with the object. Example: `"{\"*Site\":\"Antarctica\"}"`
* `lb_alternate_method`: The alternate load balancing method. Use this to select a method type from the pool if the preferred method does not return any results. Valid values are `ALL_AVAILABLE` , `DYNAMIC_RATIO` , `GLOBAL_AVAILABILITY` , `NONE` , `RATIO` , `ROUND_ROBIN` , `SOURCE_IP_HASH` , `TOPOLOGY`.
* `lb_alternate_topology`: The alternate topology for load balancing. The name of the topology ruleset. Example: `topology_name`
* `lb_dynamic_ratio_alternate`: The DTC Pool settings for dynamic ratio when it’s selected as alternate method.
  The fields to define alternate dynamic ratio are `method` , `monitor_metric` , `monitor_weighing` , `monitor_name` , `monitor_type` and `invert_monitor_metric`.

  `method`: The method of the DTC dynamic ratio load balancing. Valid values are `MONITOR` and `ROUND_TRIP_DELAY`

  `monitor_metric`: The metric of the DTC SNMP monitor that will be used for dynamic weighing Type : string. Example: `.1.2`

  `monitor_weighing`: The DTC monitor weight. ‘PRIORITY’ means that all clients will be forwarded to the least loaded server. ‘RATIO’ means that distribution will be calculated based on dynamic weights. Valid values are `PRIORITY` and `RATIO` . Default value is `RATIO`

  `invert_monitor_metric`: Determines whether the inverted values of the DTC SNMP monitor metric will be used. Default value: `false`

  `monitor_name`: The name of the monitor . Example: `https`

  `montior_type`: The type of the monitor . Example: `http`

```terraform
lb_dynamic_ratio_alternate = jsonencode({
    "monitor_name"="snmp"
    "monitor_type"="snmp"
    "method"="MONITOR"
    "monitor_metric"=".1.2"
    "monitor_weighing"="PRIORITY"
    "invert_monitor_metric"=true
  })
```
* `lb_dynamic_ratio_preferred`: The DTC Pool settings for dynamic ratio when it’s selected as preferred method.
  The fields to define alternate dynamic ratio are `method` , `monitor_metric` , `monitor_weighing` , `monitor_name` , `monitor_type` and `invert_monitor_metric`.

  `method`: The method of the DTC dynamic ratio load balancing. Valid values are `MONITOR` and `ROUND_TRIP_DELAY`

  `monitor_metric`: The metric of the DTC SNMP monitor that will be used for dynamic weighing. Type: string. Example: `.1.2`

  `monitor_weighing`: The DTC monitor weight. ‘PRIORITY’ means that all clients will be forwarded to the least loaded server. ‘RATIO’ means that distribution will be calculated based on dynamic weights. Valid values are `PRIORITY` and `RATIO` . Default value is `RATIO`

  `invert_monitor_metric`: Determines whether the inverted values of the DTC SNMP monitor metric will be used. Default value: `false`

  `monitor_name`: The name of the monitor. Example: `https`

  `montior_type`: The type of the monitor. Example: `http`

```terraform
lb_dynamic_ratio_preferred = jsonencode({
    "monitor_name"="snmp"
    "monitor_type"="snmp"
    "method"="MONITOR"
    "monitor_metric"=".1.2"
    "monitor_weighing"="PRIORITY"
    "invert_monitor_metric"=true
  })
```
* `lb_preferred_method`: The preferred load balancing method. Use this to select a method type from the pool. Valid values are `ALL_AVAILABLE` , `DYNAMIC_RATIO` , `GLOBAL_AVAILABILITY` , `NONE` , `RATIO` , `ROUND_ROBIN` , `SOURCE_IP_HASH` , `TOPOLOGY`.
* `lb_preferred_topology`: The preferred topology for load balancing. The name of the topology ruleset. Example : `topology_name`
* `monitors`: The monitors related to pool. An array of the following objects: `dtc:monitor:http`, `dtc:monitor:icmp`, `dtc:monitor:tcp`, `dtc:monitor:pdp`, `dtc:monitor:sip`, `dtc:monitor:snmp`.

  `monitor_name`: The name of the monitor. Example: `https`

  `montior_type`: The type of the monitor. Example: `http`

```terraform
monitors{
        monitor_name = "snmp"
        monitor_type="snmp"
      }
```
* `name`: The DTC Pool display name. Example: `dtc_pool`
* `quorum`: For `availability` mode QUORUM, at least this many monitors must report the resource as up for it to be available. Example : `2`
* `servers`: The servers related to the pool.

  `server`: Name of the server. Example: `dummy-server.com`

  `ratio`: The weight of server. Example: `3`

```terraform
servers{
    server = "dummy-server.com"
    ratio=3
  }
```

* `ttl`: The Time To Live (TTL) value for the DTC Pool. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. Example: `600`
* `health`: The Pool health information. Example:

  `availability`: The availability color status. Default value: `NONE`. Valid values are one of these: `BLUE`, `GREEN`, `GRAY`, `NONE`, `RED` and `YELLOW`.

  `description`: The textual description of the Pool object's status. Default value: `""`. Example: `test health`.

  `enabled_state`: The enabled state of the Pool. Default value: `ENABLED`. Valid values are one of these: `DISABLED`, `DISABLED_BY_PARENT`, `ENABLED` and `NONE`.
```terraform
health {
  availability = "NONE"
  description = ""
  enabled_state = "DISABLED"
}
```

For usage of filters, add the fields as keys and appropriate values to be passed to the keys like `name`,`comment` and  `status_member` corresponding to object.
From the below list of supported arguments for filters,  use only the searchable fields for retrieving the matching records.

### Supported Arguments for filters

-----
| Field   | Alias         | Type   | Searchable |
|---------|---------------|--------|------------|
| name    | name          | string | yes        |
| comment | comment       | string | yes        |

-----

These fields are used only for searching. 
They are not actual members of the object and therefore the server does not return these fields with this name unless they are nested return fields.

-----
| Field | Alias         | Type   | Searchable |
|-------|---------------|--------|------------|
| -     | status_member | string | yes        |
-----

!> Any of the combination from searchable fields in supported arguments list for fields are allowed.

!> Please consider using only fields as the keys in terraform datasource filters, kindly don't use alias names as keys from the above table.

### Example for using the filters:
```hcl
 data "infoblox_dtc_pool" "pool_filter" {
    filters = {
        name = "pool-test.com"
        comment = "pool creation"
    }
 }
 ```

```hcl
data "infoblox_dtc_pool" "pool_filter" {
  filters = {
    status_member = "infoblox.localdomain"
  }
}
 ```
### Example of the DTC Pool Data Source Block

```hcl
resource "infoblox_dtc_pool" "pool"{
  name="pool-test.com"	  
  comment="pool creation"	  
  lb_preferred_method="TOPOLOGY"	  
  lb_preferred_topology="topology_ruleset1"	  
  monitors{	    
    monitor_name = "snmp"
    monitor_type="snmp"
  }	  
  monitors{	    
    monitor_name = "http"
    monitor_type="http"
  }	  
  lb_alternate_method="DYNAMIC_RATIO"	  
  lb_dynamic_ratio_alternate = jsonencode({
    "monitor_name"="snmp"
    "monitor_type"="snmp"
    "method"="MONITOR"
    "monitor_metric"=".1.2"
    "monitor_weighing"="PRIORITY"
    "invert_monitor_metric"=true
  })	  
  servers{
    server = "dummy-server.com"
    ratio=3
  }	  
  servers{
    server = "server-test.com"
    ratio=3
  }
  servers{
    server = "server-test1.com"
    ratio= 4
  }
  availability = "QUORUM"	  
  quorum = 2	  
  ttl = 120	  
  ext_attrs = jsonencode({
    "Site" = "Blr"
  })
  consolidated_monitors{
    monitor_name = "http"
    monitor_type = "http"
    members = ["infoblox.localdomain"]
    availability= "ALL"
    full_health_communication= true
  }
  disable = true 
}


data "infoblox_dtc_pool" "testPool_read" {
  filters = {
    name = infoblox_dtc_pool.pool.name
    status_member = "infoblox.localdomain"
  }
}

output "dtc_rec_res" {
  value = data.infoblox_dtc_pool.testPool_read
}

data "infoblox_dtc_pool" "dtcPool_ea" {
  filters = {
    "*Site" = "Blr"
  }
}

// throws matching pools with EA, if any
output "dtcPool_out" {
  value = data.infoblox_dtc_pool.dtcPool_ea
}
```

