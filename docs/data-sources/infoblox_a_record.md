# A-record Datasource

Use the `infoblox_a_record` data source to retrieve the following information for an A-record, which is managed by a NIOS server:

* `zone`: the zone that contains the record in the specified DNS view. Example: `test.com`.
* `ttl`: the time to live value of the record, in seconds. Example: `1800`.
* `comment`: the description of the record. This is a regular comment. Example: `Temporary A-record`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as a JSON map. Example: `{“Owner”: “State Library”, “Expires”: “never”}`.

To get information about an A-record, specify a combination of the DNS view, IPv4 address that the record points to and the FQDN that corresponds to the IP address.

The following list describes the parameters you must define in an `infoblox_a_record` data source block:

* `dns_view`: the DNS view in which the zone exists. If a value is not specified, the default DNS view is considered.
* `ip_addr`: the IPv4 address associated with the A-record.
* `fqdn`: the fully qualified domain name to which the IP address is assigned.

### Example of the A-record Datasource Block

This example defines a data source of type `infoblox_a_record` and the name "vip_host", which is configured in a Terraform file. You can reference this resource and retrieve information about it. For example, `data.infoblox_a_record.vip_host.comment` returns a text as is a comment for the A-record.

```hcl
data "infoblox_a_record" "vip_host" {
  dns_view="default"
  fqdn="very-interesting-host.example.com"
  ip_addr="10.3.1.65"
}

output "id" {
  value = data.infoblox_a_record.vip_host
}


output "zone" {
  value = data.infoblox_a_record.vip_host.zone
}


output "ttl" {
  value = data.infoblox_a_record.vip_host.ttl
}


output "comment" {
  value = data.infoblox_a_record.vip_host.comment
}
```
