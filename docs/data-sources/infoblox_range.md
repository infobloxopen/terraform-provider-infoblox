# Range Data Source

Use the `infoblox_range` data source to retrieve the following information for an Range if any, which is managed by a NIOS server:

* `name`: specifies the display name. Example: `network-range`.
* `comment`: comment for the range, maximum 256 characters. Example: `test range`.
* `network`: The network to which this range belongs, in IPv4 Address/CIDR format. Example: `21.20.2.0/24`.
* `network_view`: The name of the network view in which this range resides. Example: `default`.
* `start_addr`: The IPv4 Address starting address of the range. Example: `21.20.2.20`.
* `end_addr`: The IPv4 Address end address of the range. Example: `21.20.2.40`
* `disable`: Determines whether a range is disabled or not. When this is set to False, the range is enabled. Default value: `false`.
* `extattrs`: Extensible attributes associated with the object. Example: `"{\"*Site\":\"Antarctica\"}"`
* `failover_association`: The name of the failover association: the server in this failover association will serve the IPv4 range in case the main server is out of service. `server_association_type` must be set to `FAILOVER` or `FAILOVER_MS` if you want the failover association specified here to serve the range.
* `server_association_type`: The type of server that is going to serve the range. Valid values are `FAILOVER`,`MEMBER`,`MS_FAILOVER`,`MS_SERVER`,`NONE`. Default value: `NONE`.
* `options`: An array of DHCP option structs that lists the DHCP options associated with the object.
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
* `member`: The member that will provide service for this range.
```terraform
member = jsonencode({
     ipv4addr = "10.197.81.111"
   })
```
* `template`: If set on creation, the range will be created according to the values specified in the named template. Example: `template1`

### Supported Arguments for filters

-----
| Field                   | Alias                   | Type   | Searchable |
|-------------------------|-------------------------|--------|------------|
| end_addr                | end_addr                | string | yes        |
| failover_association    | failover_association    | string | yes        |
| member                  | member                  | string | yes        |
| network                 | network                 | string | yes        |
| comment                 | comment                 | string | yes        |
| network_view            | network_view            | string | yes        |
| server_association_type | server_association_type | string | yes        |
| start_addr              | start_addr              | string | yes        |

!> Any of the combination from searchable fields in supported arguments list for fields are allowed.

!> Please consider using only fields as the keys in terraform datasource filters, kindly don't use alias names as keys from the above table.

### Example for using the filters:
 ```hcl
 data "infoblox_range" "range_rec_temp" {
  filters = {
    start_addr = "12.4.0.146"
  }
}
 ```
!> If `null` or empty filters are passed, then all the records or objects associated with datasource like here `infoblox_range` will be fetched in results.
You can reference this resource and retrieve information about it.

```hcl
resource "infoblox_network_view" "netview_range" {
  name = "custom_network_view"
}
resource "infoblox_ipv4_network" "net_range" {
  cidr = "17.0.0.0/24"
  network_view = infoblox_network_view.netview_range.name
}
resource "infoblox_range" "range" {
  start_addr = "17.0.0.221"
  end_addr   = "17.0.0.240"
  options {
    name         = "dhcp-lease-time"
    value        = "43200"
    vendor_class = "DHCP"
    num          = 51
    use_option   = true
  }
  network              = infoblox_ipv4_network.net_range.cidr
  network_view = infoblox_ipv4_network.net_range.network_view
  comment              = "test comment"
  name                 = "test_range"
  disable              = false
  member = jsonencode({
    name = "infoblox.localdomain"
  })
  server_association_type= "MEMBER"
  ext_attrs = jsonencode({
    "Site" = "Blr"
  })
  use_options = true
}

data "infoblox_range" "range_rec_temp" {
  filters = {
    start_addr = "17.0.0.221"
  }
  depends_on = [infoblox_range.range]
}

output "range_rec_res" {
  value = data.infoblox_range.range_rec_temp
}

//accessing range through EA
data "infoblox_range" "range_rec_temp_ea" {
  filters = {
    "*Site" = "Blr"
  }
}

output "range_rec_res1" {
  value = data.infoblox_range.range_rec_temp_ea
}
```