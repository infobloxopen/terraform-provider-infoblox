# AAAA-record Data Source

Use the `infoblox_aaaa_record` data source to retrieve the following information for an AAAA-record, which is managed by a NIOS server:

* `zone`: the zone that contains the record in the specified DNS view. Example: `test.com`.
* `ttl`: the "time to live" value of the record, in seconds. Example: `1800`.
* `comment`: the description of the record. This is a regular comment. Example: `Temporary AAAA-record`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as a JSON map. Example: `{"Owner": "State Library", "Expires": "never"}`.

To get information about an AAAA-record, specify a combination of the DNS view, IPv6 address that the record points to, and the FQDN that corresponds to the IP address.

The following list describes the parameters you must define in an `infoblox_aaaa_record` data source block (all except `dns_view` are required):

* `dns_view`: optional, specifies the DNS view which the record's zone belongs to. If a value is not specified, the name `default` is used as the DNS view.
* `ipv6_addr`: the IPv6 address associated with the AAAA-record.
* `fqdn`: the fully qualified domain name which the IP address is assigned to.

### Example of an AAAA-record Data Source Block

This example defines a data source of type `infoblox_aaaa_record` and the name "vip_host", which is configured in a Terraform file.
You can reference this resource and retrieve information about it. For example, `data.infoblox_aaaa_record.vip_host.comment` returns
a text as is a comment for the AAAA-record.

```hcl
data "infoblox_aaaa_record" "vip_host" {
  dns_view="default"
  fqdn="very-interesting-host.example.com"
  ipv6_addr="2a05:d014:275:cb00:ec0d:12e2:df27:aa60"
}

output "vip_host_id" {
  value = data.infoblox_aaaa_record.vip_host.id
}

output "vip_host_zone" {
  value = data.infoblox_aaaa_record.vip_host.zone
}

output "vip_host_ttl" {
  value = data.infoblox_aaaa_record.vip_host.ttl
}

output "vip_host_comment" {
  value = data.infoblox_aaaa_record.vip_host.comment
}

output "vip_host_ext_attrs" {
  value = data.infoblox_aaaa_record.vip_host.ext_attrs
}
```
