# PTR-record Data Source

Use the `infoblox_ptr_record` data source to retrieve the following information for an PTR-record, which is managed by a NIOS server:

* `ttl`: the time to live value of the record, in seconds. Example: `1800`.
* `comment`: the description of the record. This is a regular comment. Example: `manager's PC`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as a JSON map. Example: `{“Owner”: “State Library”, “Expires”: “never”}`.

To get information about a PTR-record, specify a combination of the DNS view,
IPv4/IPv6 address that the record points from and the FQDN that corresponds to the IP address.
Instead of an IP-address you may specify a record's name in FQDN format.
For example, instead of IP-address 10.0.0.1 you may specify record's name 1.0.0.10.in-addr.arpa.

The following list describes the parameters you must define in an `infoblox_ptr_record` data source block:

* `dns_view`: the DNS view in which appropriate reverse zone exists. If a value is not specified, the default DNS view is considered.
* `ip_addr`: the IP address associated with the PTR-record, either IPv4 or IPv6.
* `record_name`: the name of the DNS PTR-record in FQDN format; may be used instead of an IP-address. Example: 1.0.0.10.in-addr.arpa.
* `ptrdname`: the fully qualified domain name which PTR-record points to.

### Example of the PTR-record Data Source Block

This example defines a data source of type `infoblox_ptr_record` and the name "vip_host", which is configured in a Terraform file.
You can reference this resource and retrieve information about it. For example,
`data.infoblox_ptr_record.vip_host.comment` returns a textual content of comment field for the PTR-record.

```hcl
data "infoblox_ptr_record" "vip_host" {
  dns_view="default"
  fqdn="very-interesting-host.example.com"
  ipv6_addr="2a05:d014:275:cb00:ec0d:12e2:df27:aa60"
}

output "id" {
  value = data.infoblox_ptr_record.vip_host.id
}

output "ttl" {
  value = data.infoblox_ptr_record.vip_host.ttl
}

output "comment" {
  value = data.infoblox_ptr_record.vip_host.comment
}

output "ext_attrs" {
  value = data.infoblox_ptr_record.vip_host.ext_attrs
}
```
