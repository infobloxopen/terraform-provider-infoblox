# TXT-record Data Source

Use the data source to retrieve the following information for an TXT-record from the corresponding object in NIOS:

* `text`: the text value for the TXT-record. An empty value is not allowed.
* `zone`: the zone which the record belongs to.
* `ttl`: the "time to live" value of the record, in seconds. Example: `1800`.
* `comment`: the description of the record. This is a regular comment. Example: `spare node for the service`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as a JSON map. Example: `{"Owner”: "State Library”, "Expires”: "never”}`.

The following list describes the parameters you must define in an `infoblox_txt_record` data source block:

* `dns_view`: optional, specifies the DNS view which the record's zone belongs to. If a value is not specified, the name `default` is used as the DNS view.
* `fqdn`: required, specifies the fully qualified domain name which a textual value is assigned to. Exmample: `big-big-company.com`

### Supported Arguments for filters

-----

| Field   | Alias    | Type   | Searchable |
|---------|----------|--------|------------|
| name    | fqdn     | string | yes        |
| text    | ---      | string | yes        |
| view    | dns_view | string | yes        |
| zone    | ---      | string | yes        |
| ttl     | ---      | uint   | no         |
| comment | ---      | string | yes        |

Note: Please consider using only fields as the keys in terraform datasource, kindly don't use alias names as keys from the above table.

### Example of the TXT-record Data Source Block

```hcl
resource "infoblox_txt_record" "rec3" {
  dns_view = "nondefault_dnsview1"
  fqdn = "example3.example2.org"
  text = "data for TXT-record #3"
  ttl = 300
  comment = "example TXT record #3"
  ext_attrs = jsonencode({
    "Location" = "65.8665701230204, -37.00791763398113"
  })
}

data "infoblox_txt_record" "ds3" {
  dns_view = "nondefault_dnsview1"
  fqdn = "example3.example2.org"

  // This is just to ensure that the record has been be created
  // using 'infoblox_txt_record' resource block before the data source will be queried.
  depends_on = [infoblox_txt_record.rec3]
}

output "txt_rec3_text" {
  value = data.infoblox_txt_record.ds3.text
}

output "txt_rec3_zone" {
  value = data.infoblox_txt_record.ds3.zone
}

output "txt_rec3_ttl" {
  value = data.infoblox_txt_record.ds3.ttl
}

output "txt_rec3_comment" {
  value = data.infoblox_txt_record.ds3.comment
}

output "txt_rec3_ext_attrs" {
  value = data.infoblox_txt_record.ds3.ext_attrs
}
```
