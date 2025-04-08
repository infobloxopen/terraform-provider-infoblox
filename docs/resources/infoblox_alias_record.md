# Alias-record Resource

The `infoblox_alias_record` resource enables you to perform `create`, `update` and `delete` operations on Alias Record in a NIOS appliance.
The resource represents the ‘record:alias’ WAPI object in NIOS.

The following list describes the parameters you can define in the `infoblox_alias_record` resource block:

- `name`: required, specifies the alias name in the FQDN format. Example: `alias1.example.com`.
- `target_name`: required, specifies the target name in the FQDN format. Example: `main.example.com`.
- `target_type`: required, specifies the type of the target object. Valid values are: `A`, `AAAA`, `MX`, `NAPTR`, `PTR`, `SPF`, `SRV` and `TXT`.
- `ttl`: optional, specifies the "time to live" value for the alias-record. There is no default value for this parameter. If a value is not specified, then in NIOS, the value is inherited from the parent zone of the DNS record for this resource. A TTL value of 0 (zero) means caching should be disabled for this record. Example: `3600`.
- `disable`: optional, specifies whether the alias record is disabled or not. Default value is `false`.
- `dns_view`: required, specifies the DNS view in which the zone exists. If a value is not specified, the name `default` is set as the DNS view. Example: `dns_view_1`.
- `comment`: optional, describes the alias-record. Example: `an example alias-record`.
- `ext_attrs`: optional, specifies the set of NIOS extensible attributes that are attached to the alias-record. Example: `jsonencode({"Site":"Singapore"})`.

### Example of an Alias-record Resource

```hcl
// Alias-record, minimal set of parameters
resource "infoblox_alias_record" "alias_record_minimum_params" {
  name        = "alias-record1.test.com"
  target_name = "aa.bb.com"
  target_type = "PTR"
}

// Alias-record, full set of parameters
resource "infoblox_alias_record" "alias_record_full_params" {
  name = "alias-record2.test.com"
  target_name = "kk.ll.com"
  target_type = "AAAA"
  comment     = "example alias record"
  dns_view    = "view2"
  disable     = false
  ttl         = 120

  ext_attrs = jsonencode({
    "Site" = "MOROCCO"
  })
}
```
