# CNAME-record Data Source

Use the `infoblox_cname_record` data resource for the CNAME object to retrieve the following information for CNAME records:

* `zone`: the zone that contains the record in the specified DNS view. Example: `test.com`.
* `ttl`: the "time to live" value of the record, in seconds. Example: `3600`.
* `comment`: the text describing the record. This is a regular comment. Example: `Temporary CNAME-record`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as a JSON map. Example: `{"Owner”: "State Library”, "Expires”: "never”}`

To get information about a CNAME-record, specify a combination of the DNS view, canonical name, and an alias that the record points to.

The following list describes the parameters you must define in an `infoblox_cname_record` data source block (all of them are required):

* `dns_view`: optional, specifies the DNS view which the record's zone belongs to. If a value is not specified, the name `default` is used as the DNS view.
* `canonical`: specifies the canonical name of the record in the FQDN format.
* `alias`: specifies the alias name of the record in the FQDN format.

### Example of the CNAME-record Data Source Block

```hcl
data "infoblox_cname_record" "foo"{
  dns_view="default"
  alias="foo.test.com"
  canonical="main.test.com"
}

output "foo_ttl" {
  value = data.infoblox_cname_record.foo.ttl
}

output "foo_zone" {
  value = data.infoblox_cname_record.foo.zone
}

output "foo_comment" {
  value = data.infoblox_cname_record.foo.comment
}

output "foo_ext_attrs" {
  value = data.infoblox_cname_record.foo.ext_attrs
}
```
