# Range Template Resource

The `infoblox_ipv4_range_template` resource enables you to perform `create`, `update` and `delete` operations on IPV4 Range Template in a NIOS appliance.
The resource represents the ‘rangetemplate’ WAPI object in NIOS.

The following list describes the parameters you can define in the resource block of the Range Template object:

* `name`: required, specifies the display name of the Range Template. Example: `test-rangetemplate`.
* `number_of_addresses`: required, specifies the number of addresses for this range. Example: `100`.
* `offset`: required, specifies the start address offset for the range. Example: `30`.
* `use_options`: optional, specifies the use flag for options. Example: `true`.
* `cloud_api_compatible`: optional, specifies the flag controls whether this template can be used to create network objects in a cloud-computing deployment. Default: `false`. If the Terraform internal ID extensible attribute is set up for cloud API access, then you must set this parameter to true.
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
* `comment`: optional, specifies the description of the record. This is a regular comment. Example: `Temporary Range Template`.
* `ext_attrs`: optional, specifies the set of extensible attributes of the record, if any. The content is formatted as string of JSON map. Example: `"{\"Site\":"Nagoya"}"`
* `server_association_type`: optional, specifies the type of server that is going to serve the range. Valid values are: `FAILOVER`, `MEMBER`, `MS_FAILOVER`, `MS_SERVER`, `NONE` .Example: `NONE`.
* `failover_association`: optional, specifies the name of the failover association: the server in this failover association will serve the IPv4 range in case the main server is out of service. Example: `dhcp_failover`.
* `ms_server`: optional, specifies the Microsoft server that will provide service for this range. `server_association_type` needs to be set to `MS_SERVER` if you want the server specified here to serve the range. Example: `10.23.23.2`.
* `member`: optional, specifies the member that will provide service for this range. `server_association_type` needs to be set to `MEMBER` if you want the server specified here to serve the range. `member` has the following three fields `name`, `ipv4addr` and `ipv6addr`. Any one these `name`, `ipv4addr`, `ipv6addr` should be specified. The description of the fields of `member` is as follows:
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

!> If the user is a cloud-user, then they need Terraform internal ID with cloud permission and enable cloud delegation for the user to create a range template.

!> if the user is a non cloud-user, they need to have  Terraform internal ID without cloud permission.

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

### Examples of a Range Template Block

```hcl
// creating a Range Template record with minimal set of parameters
resource "infoblox_ipv4_range_template" "range_template_minimal_parameters" {
  name = "range-template1"
  number_of_addresses = 10
  cloud_api_compatible = true
  options {
    name         = "dhcp-lease-time"
    value        = "43200"
    vendor_class = "DHCP"
    num          = 51
    use_option   = false  
  }
  offset = 20
}

// creating a Range Template record with full set of parameters
resource "infoblox_ipv4_range_template" "range_template_full_set_parameters" {
  name = "range-template2"
  number_of_addresses = 40
  offset = 30
  comment = "Temporary Range Template"
  cloud_api_compatible = true
  use_options = true
  ext_attrs = jsonencode({
    "Site" = "Kobe"
  })
  options {
    name = "domain-name-servers"
    value = "11.22.33.44"
    vendor_class = "DHCP"
    num = 6
    use_option = true
  }
  options {
    name         = "dhcp-lease-time"
    value        = "43200"
    vendor_class = "DHCP"
    num          = 51
    use_option   = false
  }
  member {
    ipv4addr = "10.197.81.146"
    ipv6addr = "2403:8600:80cf:e10c:3a00::1192"
    name = "infoblox.localdomain"
  }
  failover_association = "failover1"
  server_association_type = "FAILOVER"
}
```
