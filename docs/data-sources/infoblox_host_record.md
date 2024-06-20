# Host-record Data Source

Use the `infoblox_host_record` data source to retrieve the following information for a Host-Record if any, which is managed by a NIOS server:

* `dns_view`: the DNS view which the record's zone belongs to. Example: `default`
* `fqdn`: the fully qualified domain name which the IP address is assigned to. `blues.test.com`
* `ipv4_addr`: the IPv4 address associated with the Host-record. Example: `10.0.0.32`
* `ipv6_addr`: the IPv6 address associated with the Host-record. Example: `2001:1890:1959:2710::32`
* `mac_addr`: the MAC address associated with the Host-record. Example: `aa:bb:cc:dd:ee:11`
* `zone`: the zone that contains the record in the specified DNS view. Example: `test.com`.
* `ttl`: the "time to live" value of the record, in seconds. Example: `1800`.
* `duid`: the DHCP Unique Identifier of the record. Example: `34:df:37:1a:d9:7f`.
* `enable_dns`: the flag to enable or disable the DNS record. Example: `true`.
* `enable_dhcp`: the flag to enable or disable the DHCP record. Example: `true`.
* `comment`: the description of the record. This is a regular comment. Example: `Temporary A-record`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as string of JSON map. Example: `"{\"TestEA\":56,\"TestEA1\":\"kickoff\"}"`

To retrieve information about host records that match the specified filters, use the `filters` argument and specify the parameters mentioned in the below table. These are the searchable parameters of the corresponding object in Infoblox NIOS WAPI. If you do not specify any parameter, the data source retrieves information about all host records in the NIOS Grid.

The following table describes the parameters you can define in an `infoblox_host_record` data source block:

### Supported Arguments for filters

-----
| Field        | Alias        | Type   | Searchable |
|--------------|--------------|--------|------------|
| name         | fqdn         | string | yes        |
| view         | dns_view     | string | yes        |
| network_view | network_view | string | yes        |
| zone         | zone         | string | yes        |
| comment      | comment      | string | yes        |

!> Aliases are the parameter names used in the prior releases of Infoblox IPAM Plug-In for Terraform. Do not use the alias names for parameters in the data source blocks. Using them can result in error scenarios.

### Example for using the filters:
```hcl
data "infoblox_host_record" "host_rec_filter" {
  filters = {
    name = "host1.example.org"
  }
}
```

!> If `null` or empty filters are passed, then all the records or objects associated with datasource like here `infoblox_host_record` will be fetched in results.

### Example of an Host-record Data Source Block

This example defines a data source of type `infoblox_host_record` and the name "host_rec_temp", which is configured in a Terraform file.
You can reference this resource and retrieve information about it.

```hcl
// This is just to ensure that the record has been be created
resource "infoblox_zone_auth" "zone1" {
  fqdn = "example.org"
  view = "default"
}

resource "infoblox_ip_allocation" "allocation1" {
  dns_view = "default"
  enable_dns = true
  fqdn = "host1.example.org"
  ipv4_addr = "10.10.0.7"
  ipv6_addr = "1::1"
  ext_attrs = jsonencode({"Location" = "USA"})
  
  depends_on = [infoblox_zone_auth.zone1]
}

resource "infoblox_ip_association" "association1" {
  internal_id = infoblox_ip_allocation.allocation1.id
  mac_addr = "12:00:43:fe:9a:8c"
  duid = "12:00:43:fe:9a:81"
  enable_dhcp = false
  depends_on = [infoblox_ip_allocation.allocation1]
}

data "infoblox_host_record" "host_rec_temp" {
  filters = {
    name = "host1.example.org"
  }
  // This is just to ensure that the record has been be created
  // using 'infoblox_host_record' resource block before the data source will be queried.
  depends_on = [infoblox_ip_association.association1]
}

// accessing Host-record through name
output "host_rec_res" {
  value = data.infoblox_host_record.host_rec_temp
}

// fetching Host-Records through EAs
data "infoblox_host_record" "host_rec_ea" {
  filters = {
    "*Location" = "USA"
  }
}

output "host_ea_out" {
  value = data.infoblox_host_record.host_rec_ea
}
```

