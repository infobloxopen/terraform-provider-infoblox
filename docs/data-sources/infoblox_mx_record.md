# MX-record Data Source

Use the data source to retrieve the following information for an MX-record from the corresponding object in NIOS:

* `zone`: the zone which the record belongs to.
* `ttl`: the "time to live" value of the record, in seconds. Example: `1800`.
* `comment`: the description of the record. This is a regular comment. Example: `spare node for the service`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as a JSON map. Example: `{"Owner”: "State Library”, "Expires”: "never”}`.

The following list describes the parameters you must define in an `infoblox_mx_record` data source block:

* `dns_view`: optional, specifies the DNS view which the record's zone belongs to. If a value is not specified, the name `default` is used as the DNS view.
* `fqdn`: required, specifies the DNS zone (as a fully qualified domain name) which a mail exchange host is assigned to. Example: `big-big-company.com`
* `mail_exchanger`: required, specifies the mail exchange host's fully qualified domain name. Example: `mx1.secure-mail-provider.net`
* `preference`: required, specifies the preference number (0-65535) for this MX-record.

### Supported Arguments for filters

-----
| Field          | Alias    | Type   | Searchable |
|----------------|----------|--------|------------|
| name           | fqdn     | string | yes        |
| mail_exchanger | ---      | string | yes        |
| preference     | ---      | uint32 | yes        |
| view           | dns_view | string | yes        |
| ttl            | ---      | uint32 | no         |
| comment        | ---      | string | yes        |
| zone           | ---      | string | yes        |

Note: Please consider using only fields as the keys in terraform datasource, kindly don't use alias names as keys from the above table.

### Example of the MX-record Data Source Block

```hcl
resource "infoblox_mx_record" "rec2" {
  dns_view = "nondefault_dnsview1"
  fqdn = "rec2.example2.org"
  mail_exchanger = "sample.test.com"
  preference = 40
  comment = "example MX-record"
  ttl = 120
  ext_attrs = jsonencode({
    "Location" = "Las Vegas"
  })
}

data "infoblox_mx_record" "ds2" {
  dns_view = "nondefault_dnsview1"
  fqdn = "rec2.example2.org"
  mail_exchanger = "sample.test.com"
  preference = 40

  // This is just to ensure that the record has been be created
  // using 'infoblox_mx_record' resource block before the data source will be queried.
  depends_on = [infoblox_mx_record.rec2]
}

output "mx_rec2_zone" {
  value = data.infoblox_mx_record.ds2.zone
}

output "mx_rec2_ttl" {
  value = data.infoblox_mx_record.ds2.ttl
}

output "mx_rec2_comment" {
  value = data.infoblox_mx_record.ds2.comment
}

output "mx_rec2_ext_attrs" {
  value = data.infoblox_mx_record.ds2.ext_attrs
}
```
