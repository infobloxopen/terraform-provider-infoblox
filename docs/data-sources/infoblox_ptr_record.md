# PTR-record Data Source

Use the data source to retrieve the following information for a PTR-record from the corresponding object in NIOS:

* `zone`: the zone which the record belongs to.
* `ttl`: the "time to live" value of the record, in seconds. Example: `1800`.
* `comment`: the description of the record. This is a regular comment. Example: `manager's PC`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as a JSON map. Example: `{"Owner”: "State Library”, "Expires”: "never”}`.

To get information about an PTR-record, specify a combination of the DNS view, IPv4 address that the record points to
or the record name in FQDN format, and the FQDN that corresponds to the IP address.

The following list describes the parameters you must define in an `infoblox_ptr_record` data source block:

* `dns_view`: optional, specifies the DNS view in which the reverse mapping zone exists. If a value is not specified, the name `default` is used as the DNS view.
* `ip_addr`: required only if record_name is not set; specifies the IPv4 or IPv6 address associated with the PTR-record. 
  If both `ip_addr` and `record_name` are set, `ip_addr` takes precedence.
* `record_name`: required only if ip_addr is not set; specifies the name of the PTR-record in FQDN format,
  which can be used instead of an IP address. Example: `1.0.0.10.in-addr.arpa`.
* `ptrdname`: required; specifies the fully qualified domain name that the PTR-record points to.

### Example of the PTR-record Data Source Block

This example defines a data source of type `infoblox_ptr_record` and the name "vip_host", which is configured in a Terraform file.
You can reference this resource and retrieve information about it. For example,
`data.infoblox_ptr_record.vip_host.comment` returns a textual content of comment field for the PTR-record.

```hcl
data "infoblox_ptr_record" "host1" {
  dns_view="default"
  ptrdname="host.example.org"
  ip_addr="2a05:d014:275:cb00:ec0d:12e2:df27:aa60"
}

output "host1_id" {
  value = data.infoblox_ptr_record.host1.id
}

output "host1_ip_addr" {
  value = data.infoblox_ptr_record.host1.ip_addr
}

output "host1_record_name" {
  value = data.infoblox_ptr_record.host1.record_name
}

output "host1_ttl" {
  value = data.infoblox_ptr_record.host1.ttl
}

output "host1_comment" {
  value = data.infoblox_ptr_record.host1.comment
}

output "host1_ext_attrs" {
  value = data.infoblox_ptr_record.host1.ext_attrs
}

data "infoblox_ptr_record" "host2" {
  dns_view="default"
  ptrdname="host.example.org"
  record_name="0.6.a.a.7.2.f.d.2.e.2.1.d.0.c.e.0.0.b.c.5.7.2.0.4.1.0.d.5.0.a.2.ip6.arpa"
}

output "host2_id" {
  value = data.infoblox_ptr_record.host2.id
}

output "host2_ip_addr" {
  value = data.infoblox_ptr_record.host2.ip_addr
}

output "host2_record_name" {
  value = data.infoblox_ptr_record.host2.record_name
}

output "host2_ttl" {
  value = data.infoblox_ptr_record.host2.ttl
}

output "host2_comment" {
  value = data.infoblox_ptr_record.host2.comment
}

output "host2_ext_attrs" {
  value = data.infoblox_ptr_record.host2.ext_attrs
}
```
