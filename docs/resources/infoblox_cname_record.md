# CNAME-record Resource

A CNAME-record maps one domain name to another (canonical) one. The `infoblox_cname_record` resource allows managing such domain name mappings in a NIOS server for CNAME records.

The following list describes the parameters you can define in the `infoblox_cname_record` resource block:

* `alias`: required, specifies the alias name in the FQDN format. Example: `alias1.example.com`.
* `canonical`: required, specifies the canonical name in the FQDN format. Example: `main.example.com`.
* `ttl`: optional, specifies the time to live value for the CNAME-record. There is no default value for this parameter. If you do not specify a value, the TTL value is inherited from Grid DNS properties. A TTL value of 0 (zero) means caching should be disabled for this record. Example: `3600`.
* `dns_view`: optional, specifies the DNS view in which the zone exists. If a value is not specified, the default DNS view is considered. Example: `dns_view_1`.
* `comment`: optional, describes the CNAME record. Example: `an example CNAME-record`.
* `ext_attrs`: optional, specifies the set of NIOS extensible attributes that are attached to the CNAME-record. Example: `jsonencode({})`.

### Example of the CNAME-record Resource

```hcl
resource "infoblox_cname_record" "ib_cname_record"{
  dns_view = "default" # the same as not specifying the attribute
  canonical = "CanonicalTestName.xyz.com"
  alias = "AliasTestName.xyz.com"
  ttl = 3600

  comment = "an example CNAME record"
  ext_attrs = jsonencode({
     "Tenant ID" = "tf-plugin"
     "CMP Type" = "Terraform"
     "Cloud API Owned" = "True"
  })
}
```
