# A-record Data Source

Use the `infoblox_a_record` data source to retrieve the following information for an A-record, which is managed by a NIOS server:

* `zone`: the zone that contains the record in the specified DNS view. Example: `test.com`.
* `ttl`: the "time to live" value of the record, in seconds. Example: `1800`.
* `comment`: the description of the record. This is a regular comment. Example: `Temporary A-record`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as a JSON map. Example: `{"Owner": "State Library", "Expires": "never"}`.

To get information about an A-record, specify a combination of the DNS view, IPv4 address that the record points to, and the FQDN that corresponds to the IP address.

The following list describes the parameters you must define in an `infoblox_a_record` data source block (all of them are required):

* `dns_view`: optional, specifies the DNS view which the record's zone belongs to. If a value is not specified, the name `default` is used as the DNS view.
* `ip_addr`: the IPv4 address associated with the A-record.
* `fqdn`: the fully qualified domain name which the IP address is assigned to.

### Supported Arguments for filters

-----
| Field    | Alias    | Type   | Searchable |
|----------|----------|--------|------------|
| name     | fqdn     | string | yes        |
| view     | dns_view | string | yes        |
| zone     | ---      | string | yes        |
| ttl      | ---      | uint   | no         |
| comment  | ---      | string | yes        |
| ipv4addr | ip_addr  | string | yes        |

Note: Please consider using only fields as the keys in terraform datasource, kindly don't use alias names as keys from the above table.

### Example of an A-record Data Source Block

This example defines a data source of type `infoblox_a_record` and the name "vip_host", which is configured in a Terraform file.
You can reference this resource and retrieve information about it. For example, `data.infoblox_a_record.vip_host.comment` returns
a text as is a comment for the A-record.

```hcl
resource "infoblox_a_record" "vip_host" {
  fqdn = "very-interesting-host.example.com"
  ip_addr = "10.3.1.65"
  comment = "special host"
  dns_view = "nondefault_dnsview2"
  ttl = 120 // 120s
  ext_attrs = jsonencode({
    "Location" = "65.8665701230204, -37.00791763398113"
  })
}


data "infoblox_a_record" "vip_host" {
  dns_view="nondefault_dnsview2"
  fqdn="very-interesting-host.example.com"
  ip_addr="10.3.1.65"
  
  // This is just to ensure that the record has been be created
  // using 'infoblox_a_record' resource block before the data source will be queried.
  depends_on = [infoblox_a_record.vip_host]
}

output "vip_host_id" {
  value = data.infoblox_a_record.vip_host.id
}

output "vip_host_zone" {
  value = data.infoblox_a_record.vip_host.zone
}

output "vip_host_ttl" {
  value = data.infoblox_a_record.vip_host.ttl
}

output "vip_host_comment" {
  value = data.infoblox_a_record.vip_host.comment
}

output "vip_host_ext_attrs" {
  value = data.infoblox_a_record.vip_host.ext_attrs
}
```

