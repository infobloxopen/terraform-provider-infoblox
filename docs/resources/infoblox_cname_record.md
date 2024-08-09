# CNAME-record Resource

A CNAME-record maps one domain name to another (canonical) one. The `infoblox_cname_record` resource allows managing such domain name mappings in a NIOS server for CNAME records.

The following list describes the parameters you can define in the `infoblox_cname_record` resource block:

* `alias`: required, specifies the alias name in the FQDN format. Example: `alias1.example.com`.
* `canonical`: required, specifies the canonical name in the FQDN format. Example: `main.example.com`.
* `ttl`: optional, specifies the "time to live" value for the CNAME-record. There is no default value for this parameter. If a value is not specified, then in NIOS, the value is inherited from the parent zone of the DNS record for this resource. A TTL value of 0 (zero) means caching should be disabled for this record. Example: `3600`.
* `dns_view`: optional, specifies the DNS view in which the zone exists. If a value is not specified, the name `default` is set as the DNS view. Example: `dns_view_1`.
* `comment`: optional, describes the CNAME-record. Example: `an example CNAME-record`.
* `ext_attrs`: optional, specifies the set of NIOS extensible attributes that are attached to the CNAME-record. Example: `jsonencode({})`.

### Example of a CNAME-record Resource

```hcl
// CNAME-record, minimal set of parameters
resource "infoblox_cname_record" "cname_rec1" {
  canonical = "bla-bla-bla.somewhere.in.the.net"
  alias     = "hq-server.example1.org"
}

// CNAME-record, full set of parameters
resource "infoblox_cname_record" "cname_rec2" {
  dns_view  = "default.nondefault_netview"
  canonical = "strange-place.somewhere.in.the.net"
  alias     = "alarm-server.example3.org"
  comment   = "we need to keep an eye on this strange host"
  ttl       = 0 // disable caching
  ext_attrs = jsonencode({
    Site     = "unknown"
    Location = "TBD"
  })
}
```
