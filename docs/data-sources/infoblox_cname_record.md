# CNAME-record Data Source

Use the `infoblox_cname_record` data resource for the CNAME object to retrieve the following information for CNAME records:

* `zone`: the zone that contains the record in the specified DNS view. Example: `test.com`.
* `ttl`: the time to live value of the record, in seconds. Example: `3600`.
* `comment`: the text describing the record. This is a regular comment. Example: `Temporary A-record`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as a JSON map. Example: `{“Owner”: “State Library”, “Expires”: “never”}`

To get information about a CNAME-record, specify a combination of the DNS view, canonical name, and an alias that the record points to.

The following list describes the parameters you must define in an `infoblox_cname_record` data source block:

* `dns_view`: the DNS view in which the zone exists. If a value is not specified, the default DNS view defined in NIOS is considered.
* `canonical`: the canonical name of the record in the FQDN format.
* `alias`: the alias name of the record in the FQDN format.

### Example of the CNAME-record Data Source Block

```hcl
data "infoblox_cname_record" "foo"{
  dns_view="default"
  alias="foo.test.com"
  canonical="main.test.com"
}

output "ttl" {
  value = data.infoblox_cname_record.foo.ttl
}

output "zone" {
  value = data.infoblox_cname_record.foo.zone
}

output "comment" {
  value = data.infoblox_cname_record.foo.comment
}

output "ext_attrs" {
  value = data.infoblox_cname_record.foo.ext_attrs
}
```
