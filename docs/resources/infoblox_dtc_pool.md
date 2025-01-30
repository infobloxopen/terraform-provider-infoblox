# DTC-Pool Resource

DTC pools are load balanced servers that distribute client requests to multiple servers using a load balancing pool.

The `infoblox_dtc_pool` resource, enables you to create, update, or delete an DTC pools in a NIOS appliance.

The following list describes the parameters you can define in the `infoblox_dtc_pool` resource block:

* `auto_consolidated_monitors`: optional, Flag for enabling auto managing DTC Consolidated Monitors in DTC Pool. Default value: `false`
* `availability`: optional, A resource in the pool is available if `ANY`, at least `QUORUM`, or `ALL` monitors for the pool say that it is up. Default value: `ALL`
* `comment`: optional, The comment for the DTC Pool; maximum 256 characters. Example: `pool creation`
* `consolidated_monitors`: optional, List of monitors and associated members statuses of which are shared across members and consolidated in server availability determination.

  `monitor_name`: name of the monitor Example:` https`

  `montior_type`: Type of the monitor Example: `http`

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
* `disable`: optional, Determines whether the DTC Pool is disabled or not. When this is set to False, the fixed address is enabled. Default value: `false`
* `extattrs`: optional, Extensible attributes associated with the object. Example: `"{\"*Site\":\"Antarctica\"}"`
* `lb_alternate_method`: optional, The alternate load balancing method. Use this to select a method type from the pool if the preferred method does not return any results. Valid values are `ALL_AVAILABLE` , `DYNAMIC_RATIO` , `GLOBAL_AVAILABILITY` , `NONE` , `RATIO` , `ROUND_ROBIN` , `SOURCE_IP_HASH` , `TOPOLOGY`.
* `lb_alternate_topology`: optional, The alternate topology for load balancing. The name of the topology ruleset. Example: `topology_name`
* `lb_dynamic_ratio_alternate`: optional, The DTC Pool settings for dynamic ratio when it’s selected as alternate method.
  The fields to define alternate dynamic ratio are `method` , `monitor_metric` , `monitor_weighing` , `monitor_name` , `monitor_type` and `invert_monitor_metric`.

  `method`: The method of the DTC dynamic ratio load balancing. Valid values are `MONITOR` and `ROUND_TRIP_DELAY`

  `monitor_metric`: The metric of the DTC SNMP monitor that will be used for dynamic weighing Type : string . Example: `.1.2`

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
* `lb_dynamic_ratio_preferred`: optional, The DTC Pool settings for dynamic ratio when it’s selected as preferred method.
  The fields to define alternate dynamic ratio are `method` , `monitor_metric` , `monitor_weighing` , `monitor_name` , `monitor_type` and `invert_monitor_metric`.

  `method`: The method of the DTC dynamic ratio load balancing. Valid values are `MONITOR` and `ROUND_TRIP_DELAY`

  `monitor_metric`: The metric of the DTC SNMP monitor that will be used for dynamic weighing Type: string. Example: `.1.2`

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
* `lb_preferred_method`: required, The preferred load balancing method. Use this to select a method type from the pool. Valid values are `ALL_AVAILABLE` , `DYNAMIC_RATIO` , `GLOBAL_AVAILABILITY` , `NONE` , `RATIO` , `ROUND_ROBIN` , `SOURCE_IP_HASH` , `TOPOLOGY`.
* `lb_preferred_topology`: optional, The preferred topology for load balancing. The name of the topology ruleset. Example: `topology_name`
* `monitors`: optional, The monitors related to pool. An array of the following objects: `dtc:monitor:http`, `dtc:monitor:icmp`, `dtc:monitor:tcp`, `dtc:monitor:pdp`, `dtc:monitor:sip`, `dtc:monitor:snmp`.

  `monitor_name`: The name of the monitor. Example: `https`

  `montior_type`: The type of the monitor. Example: `http`

```terraform
monitors{
        monitor_name = "snmp"
        monitor_type="snmp"
      }
```
* `name`: required,  The DTC Pool display name. Example: `dtc_pool`
* `quorum`: optional, For `availability` mode QUORUM, at least this many monitors must report the resource as up for it to be available Example: `2`
* `servers`: optional, The servers related to the pool.

  `server`: Name of the server. Example: `dummy-server.com`

  `ratio`: The weight of server. Example: `3`

```terraform
servers{
    server = "dummy-server.com"
    ratio=3
  }
```

* `ttl`: optional, The Time To Live (TTL) value for the DTC Pool. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. Example: `600`

### Examples of an DTC-Pool Block

```hcl
//Pool creation with minimal parameters 
resource "infoblox_dtc_pool" "test_pool1" {
  name = "Pool63"
  lb_preferred_method = "ROUND_ROBIN"
}

//Pool creation with maximal parameters 
//parameters for DTC pool when preferred load balancing is TOPOLOGY and alternate load balancing is DYNAMIC_RATIO
resource "infoblox_dtc_pool" "pool" {
  name                  = "terraform_pool.com"
  comment               = "testing pool terraform"
  lb_preferred_method   = "TOPOLOGY"
  lb_preferred_topology = "topology_ruleset"
  ext_attrs = jsonencode({
    "Site" = "Blr"
  })
  servers {
    server = "server.com"
    ratio  = 3
  }
  servers {
    server = "terraform_server.com"
    ratio  = 3
  }
  servers {
    server = "terraform_server1.com"
    ratio  = 4
  }
  monitors {
    monitor_name = "http"
    monitor_type = "http"
  }
  monitors {
    monitor_name = "snmp"
    monitor_type = "snmp"
  }
  lb_alternate_method = "DYNAMIC_RATIO"
  lb_dynamic_ratio_alternate = jsonencode({
    "monitor_name"          = "snmp"
    "monitor_type"          = "snmp"
    "method"                = "MONITOR"
    "monitor_metric"        = ".1.2"
    "monitor_weighing"      = "PRIORITY"
    "invert_monitor_metric" = true
  })
  availability               = "QUORUM"
  quorum                     = 2
  ttl                        = 120
  consolidated_monitors{
    monitor_name = "http"
    monitor_type = "http"
    members = ["infoblox.localdomain"]
    availability= "ALL"
    full_health_communication= true
  }
  disble = true
}

//parameters for DTC pool when preferred load balancing is DYNAMIC_RATIO
resource "infoblox_dtc_pool" "test_pool3" {
  name = "Pool64"
  monitors{
    monitor_name = "snmp"
    monitor_type="snmp"
  }
  lb_preferred_method = "DYNAMIC_RATIO"
  lb_dynamic_ratio_preferred = jsonencode({
    "monitor_name"="snmp"
    "monitor_type"="snmp"
    "method"="MONITOR"
    "monitor_metric"=".1.2"
    "monitor_weighing"="PRIORITY"
    "invert_monitor_metric"=true
  })
}

```
