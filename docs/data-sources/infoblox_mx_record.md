# MX-record Data Source

Use the data source to retrieve the following information for an MX-record from the corresponding object in NIOS:

* `zone`: the zone which the record belongs to.
* `ttl`: the "time to live" value of the record, in seconds. Example: `1800`.
* `comment`: the description of the record. This is a regular comment. Example: `spare node for the service`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as a JSON map. Example: `{"Owner”: "State Library”, "Expires”: "never”}`.

The following list describes the parameters you must define in an `infoblox_mx_record` data source block:

* `dns_view`: optional, specifies the DNS view in which the reverse mapping zone exists. If a value is not specified, the name `default` is used as the DNS view.
* `fqdn`: required, required, specifies the fully qualified domain name which a mail exchange host is assigned to. Example: `big-big-company.com`
* `mail_exchanger`: required, specifies the mail exchange host's fully qualified domain name. Example: `mx1.secure-mail-provider.net`
* `preference`: required, specifies the preference number (0-65535) for this MX-record.

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

  depends_on = [infoblox_mx_record.rec2]
}
```
