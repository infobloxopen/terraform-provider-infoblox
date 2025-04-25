# Range Data Source

Use the `infoblox_ipv4_range` data source to retrieve the following information for an Range if any, which is managed by a NIOS server:

* `name`: specifies the display name. Example: `network-range`.
* `comment`: comment for the range, maximum 256 characters. Example: `test range`.
* `network`: The network to which this range belongs, in IPv4 Address/CIDR format. Example: `21.20.2.0/24`.
* `network_view`: The name of the network view in which this range resides. Example: `default`.
* `start_addr`: The IPv4 Address starting address of the range. Example: `21.20.2.20`.
* `end_addr`: The IPv4 Address end address of the range. Example: `21.20.2.40`
* `disable`: Determines whether a range is disabled or not. When this is set to False, the range is enabled. Default value: `false`.
* `ext_attrs`: Extensible attributes associated with the object. Example: `"{\"*Site\":\"Antarctica\"}"`
* `failover_association`: The name of the failover association: the server in this failover association will serve the IPv4 range in case the main server is out of service. Example: `dhcp_failover`.
* `server_association_type`: The type of server that is going to serve the range. Valid values are `FAILOVER`,`MEMBER`,`MS_FAILOVER`,`MS_SERVER`,`NONE`. Default value: `NONE`.
* `options`: An array of DHCP option structs that lists the DHCP options associated with the object. The description of the fields of `options` is as follows:
  * `name`: The Name of the DHCP option. Example: `domain-name-servers`.
  * `num`: The code of the DHCP option. Example: `6`.
  * `value`: The value of the option. Example: `11.22.33.44`.
  * `vendor_class`: The name of the space this DHCP option is associated to. Default value is `DHCP`.
  * `use_option`:Only applies to special options that are displayed separately from other options and have a use flag. These options are `router`,
    `router-templates`, `domain-name-servers`, `domain-name`, `broadcast-address`, `broadcast-address-offset`, `dhcp-lease-time`, and `dhcp6.name-servers`.
```terraform
options {
    name         = "dhcp-lease-time"
    value        = "43200"
    vendor_class = "DHCP"
    num          = 51
    use_option   = true
  }
```
* `use_options`: Use option is a flag that indicates whether the options field are used or not. The default value is false. Example: `false`
* `ms_server`: The Microsoft server that will provide service for this range. server_association_type needs to be set to MS_SERVER if you want the server specified here to serve the range. Example: `10.23.23.2`
* `member`: The member that will provide service for this range. `server_association_type` needs to be set to `MEMBER` if you want the server specified here to serve the range. `member` has the following three fields `name`, `ipv4addr` and `ipv6addr`.The description of the fields of `member` is as follows:
  * `name`: The name of the Grid member. Example: `infoblox.localdomain`.
  * `ipv4addr`: The IPv4 Address of the Grid Member. Example: `11.10.1.0`.
  * `ipv6addr`: The IPv6 address of the member. Example: `2403:8600:80cf:e10c:3a00::1192`.

Example for `member`:
```terraform
member = { 
    name = "infoblox.localdomain"
    ipv4addr = "11.22.33.44"
    ipv6addr = "2403:8600:80cf:e10c:3a00::1192"
  }
```

### Supported Arguments for filters

-----
| Field                   | Alias                   | Type   | Searchable |
|-------------------------|-------------------------|--------|------------|
| end_addr                | end_addr                | string | yes        |
| failover_association    | failover_association    | string | yes        |
| network                 | network                 | string | yes        |
| comment                 | comment                 | string | yes        |
| network_view            | network_view            | string | yes        |
| server_association_type | server_association_type | string | yes        |
| start_addr              | start_addr              | string | yes        |

!> Any of the combination from searchable fields in supported arguments list for fields are allowed.

!> Please consider using only fields as the keys in terraform datasource filters, kindly don't use alias names as keys from the above table.

!> The search functionality using the filters argument is not supported for member and ms_server fields.

### Example for using the filters:
 ```hcl
 data "infoblox_ipv4_range" "range_rec_temp" {
  filters = {
    start_addr = "12.4.0.146"
  }
}
 ```
!> If `null` or empty filters are passed, then all the records or objects associated with datasource like here `infoblox_range` will be fetched in results.
You can reference this resource and retrieve information about it.

```hcl
resource "infoblox_ipv4_range" "range" {
  start_addr = "17.0.0.221"
  end_addr   = "17.0.0.240"
  options {
    name         = "dhcp-lease-time"
    value        = "43200"
    vendor_class = "DHCP"
    num          = 51
    use_option   = true
  }
  network              = "17.0.0.0/24"
  network_view = "default"
  comment              = "test comment"
  name                 = "test_range"
  disable              = false
  member = {
    name = "infoblox.localdomain"
    ipv4addr = "10.197.2.19"
  }
  server_association_type= "MEMBER"
  ext_attrs = jsonencode({
    "Site" = "Blr"
  })
  use_options = true
}

data "infoblox_ipv4_range" "range_rec_temp" {
  filters = {
    start_addr = "17.0.0.221"
  }
  depends_on = [infoblox_ipv4_range.range]
}

output "range_rec_res" {
  value = data.infoblox_ipv4_range.range_rec_temp
}

//accessing range through EA
data "infoblox_ipv4_range" "range_rec_temp_ea" {
  filters = {
    "*Site" = "Blr"
  }
}

output "range_rec_res1" {
  value = data.infoblox_range.range_rec_temp_ea
}
```