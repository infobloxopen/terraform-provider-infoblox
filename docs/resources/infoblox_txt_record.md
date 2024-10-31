# TXT-record Resource

The `infoblox_txt_record` resource associates a text value with a domain name.

The following list describes the parameters you can define in the resource block of the record:

* `fqdn`: required, specifies the fully qualified domain name which you want to assign the text value for. Example: `host43.zone12.org`
* `text`: required, specifies the text value for the TXT-record. It can contain substrings of up to 255 bytes and a total of up to 512 bytes. If you enter leading, trailing, or embedded spaces in the text string, enclose the entire string within `\"` characters to preserve the spaces.
* `dns_view`: optional, specifies the DNS view which the zone exists in. If a value is not specified, the name `default` is used for DNS view. Example: `dns_view_1`
* `ttl`: optional, specifies the "time to live" value for the record. There is no default value for this parameter. If a value is not specified, then in NIOS, the value is inherited from the parent zone of the DNS record for this resource. A TTL value of 0 (zero) means caching should be disabled for this record. Example: `600`
* `comment`: optional, describes the record. Example: `auto-created test record #1`
* `ext_attrs`: optional, a set of NIOS extensible attributes that are attached to the record. Example: `jsonencode({})`

## Examples

```hcl
# TXT-Record, minimal set of parameters
resource "infoblox_txt_record" "rec1" {
  fqdn = "sample1.example.org"
  text = "\"this is just a sample\""
}

# Some parameters for a TXT-Record
resource "infoblox_txt_record" "rec2" {
  dns_view = "default" // may be omitted
  fqdn = "sample2.example.org"
  text = "\"data for TXT-record #2\""
  ttl = 120 // 120s
}

# All the parameters for a TXT-Record
resource "infoblox_txt_record" "rec3" {
  dns_view = "nondefault_dnsview1"
  fqdn = "example3.example2.org"
  text = "\"data for TXT-record #3\""
  ttl = 300
  comment = "example TXT record #3"
  ext_attrs = jsonencode({
    "Location" = "65.8665701230204, -37.00791763398113"
  })
}
```
