# A-record Data Source

Use the `infoblox_a_record` data source to retrieve the following information for an A-record, which is managed by a NIOS server:

* `zone`: the zone that contains the record in the specified DNS view. Example: `test.com`.
* `ttl`: the "time to live" value of the record, in seconds. Example: `1800`.
* `comment`: the description of the record. This is a regular comment. Example: `Temporary A-record`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as a JSON map. Example: `{"Owner": "State Library", "Expires": "never"}`.

To get information about an A-record, specify a combination of the DNS view, IPv4 address that the record points to, and the FQDN that corresponds to the IP address.

The following list describes the parameters you must define in an `infoblox_a_record` data source block (all of them are required):

* `dns_view`: the DNS view in which the zone exists.
* `ip_addr`: the IPv4 address associated with the A-record.
* `fqdn`: the fully qualified domain name which the IP address is assigned to.

### Example of an A-record Data Source Block

This example defines a data source of type `infoblox_a_record` and the name "vip_host", which is configured in a Terraform file.
You can reference this resource and retrieve information about it. For example, `data.infoblox_a_record.vip_host.comment` returns
a text as is a comment for the A-record.

```hcl
data "infoblox_a_record" "vip_host" {
  dns_view="default"
  fqdn="very-interesting-host.example.com"
  ip_addr="10.3.1.65"
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
