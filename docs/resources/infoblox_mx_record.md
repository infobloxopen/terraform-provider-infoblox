# MX-record Resource

The `infoblox_mx_record` resource corresponds to MX-record (mail exchanger record) on NIOS side,
and it associates a mail exchange host to a domain name.

The following list describes the parameters you can define in the resource block of the record:

* `fqdn`: required, specifies the fully qualified domain name which you want to assign a mail exchange host for. Example: `big-big-company.com`
* `mail_exchanger`: required, specifies the mail exchange host's fully qualified domain name. Example: `mx1.secure-mail-provider.net`
* `preference`: required, specifies the preference number (0-65535) for this MX-record.
* `dns_view`: optional, specifies the DNS view which the zone exists in. If a value is not specified, the name `default` is used for DNS view. Example: `dns_view_1`
* `ttl`: optional, specifies the "time to live" value for the record. There is no default value for this parameter. If a value is not specified, then in NIOS, the value is inherited from the parent zone of the DNS record for this resource. A TTL value of 0 (zero) means caching should be disabled for this record. Example: `600`
* `comment`: optional, describes the record. Example: `auto-created test record #1`
* `ext_attrs`: optional, a set of NIOS extensible attributes that are attached to the record. Example: `jsonencode({})`

## Examples

```hcl
// MX-record, minimal set of parameters
resource "infoblox_mx_record" "rec1" {
  fqdn = "big-big-company.com"
  mail_exchanger = "mx1.secure-mail-provider.net"
  preference = 30
}

// MX-record, full set of parameters
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
```
