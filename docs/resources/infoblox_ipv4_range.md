# Range Resource

The `infoblox_ipv4_range` resource enables you to perform `create`, `update` and `delete` operations on Network Range in a NIOS appliance.
The resource represents the ‘range’ WAPI object in NIOS.

The following list describes the parameters you can define in the resource block of the Network Range object:

* `name`: optional, specifies the display name. Example: `network-range`.
* `comment`: optional, comment for the range, maximum 256 characters. Example: `test range`.
* `network`: optional, The network to which this range belongs, in IPv4 Address/CIDR format. Example: `21.20.2.0/24`.
* `network_view`: optional, The name of the network view in which this range resides. Example: `default`.
* `start_addr`: required, The IPv4 Address starting address of the range. Example: `21.20.2.20`.
* `end_addr`: required, The IPv4 Address end address of the range. Example: `21.20.2.40`
* `disable`: optional, Determines whether a range is disabled or not. When this is set to False, the range is enabled. Default value: `false`. 
* `ext_attrs`: optional, Extensible attributes associated with the object. Example: `"{\"*Site\":\"Antarctica\"}"`
* `failover_association`: optional, The name of the failover association: the server in this failover association will serve the IPv4 range in case the main server is out of service. `server_association_type` must be set to `FAILOVER` or `FAILOVER_MS` if you want the failover association specified here to serve the range.
* `server_association_type`: optional, The type of server that is going to serve the range. Valid values are `FAILOVER`,`MEMBER`,`MS_FAILOVER`,`MS_SERVER`,`NONE`. Default value: `NONE`.
* `ms_server`: optional, specifies the IP address of the Microsoft server that will provide service for this range. server_association_type needs to be set to MS_SERVER if you want the server specified here to serve the range. Example: `10.23.23.2`
* `options`: optional, specifies an array of DHCP option structs that lists the DHCP options associated with the object. The description of the fields of `options` is as follows:
  * `name`: required, specifies the Name of the DHCP option. Example: `domain-name-servers`.
  * `num`: required, specifies the code of the DHCP option. Example: `6`.
  * `value`: required, specifies the value of the option. Example: `11.22.33.44`.
  * `vendor_class`: optional, specifies the name of the space this DHCP option is associated to. Default value is `DHCP`.
  * `use_option`: optional, only applies to special options that are displayed separately from other options and have a use flag. These options are `router`,
    `router-templates`, `domain-name-servers`, `domain-name`, `broadcast-address`, `broadcast-address-offset`, `dhcp-lease-time`, and `dhcp6.name-servers`.
```terraform
options {
    name         = "dhcp-lease-time"
    value        = "43200"
    vendor_class = "DHCP"
    num          = 51
    use_option   = false
  }
```
* `use_options`: optional, Use option is a flag that indicates whether the options field are used or not. The default value is false. Example: `false`
* `member`: optional, specifies the member that will provide service for this range. `server_association_type` needs to be set to `MEMBER` if you want the server specified here to serve the range. `member` has the following three fields `name`, `ipv4addr` and `ipv6addr`. At least one of `name`, `ipv4addr`, or `ipv6addr` is required in the `member` block.
  The description of the fields of `member` is as follows:
  * `name`: optional, specifies the name of the Grid member. Example: `infoblox.localdomain`.
  * `ipv4addr`: optional, specifies the IPv4 Address of the Grid Member. Example: `11.10.1.0`.
  * `ipv6addr`: optional, specifies the IPv6 address of the member. Example: `2403:8600:80cf:e10c:3a00::1192`.

Example for `member`:
```terraform
member = {
  name = "infoblox.localdomain"
  ipv4addr = "11.10.1.0"
  ipv6addr = "2403:8600:80cf:e10c:3a00::1192"
}
```
* `template` : optional, If set on creation, the range will be created according to the values specified in the named template. Example: `range_template`

!> When configuring the options parameter, you must define the default option dhcp-lease-time to avoid the undesirable changes that can occur when the next terraform apply command runs. The sub parameters name, num, and value are required. An example block is as follows:
```terraform
options {
  name         = "dhcp-lease-time"
  value        = "43200"
  vendor_class = "DHCP"
  num          = 51
  use_option   = false
}
```

### Examples of a Network Range Block
```hcl
// creating a Network Range
resource "infoblox_ipv4_range" "range3" {
  start_addr = "17.0.0.221"
  end_addr   = "17.0.0.240"
  options {
    name         = "dhcp-lease-time"
    value        = "43200"
    vendor_class = "DHCP"
    num          = 51
    use_option   = false
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
```
