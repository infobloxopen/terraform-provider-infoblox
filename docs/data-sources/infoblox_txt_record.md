# TXT-record Data Source

Use the data source to retrieve the following information for TXT-record from the corresponding object in NIOS:

- `dns_view`: the DNS view which the record's zone belongs to.
- `fqdn`: the fully qualified domain name which a textual value is assigned to. Example: `sampletxt.demo.com`
- `text`: the text value for the TXT-record. Example: `some random next`
- `zone`: the zone which the record belongs to.
- `ttl`: the "time to live" value of the record, in seconds. Example: `1800`.
- `comment`: the description of the record. This is a regular comment. Example: `spare node for the service`.
- `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as string of JSON map. Example: `"{\"Owner\":\"State Library\", \"Expires\":\"never\"}"`.

For usage of filters, add the fields as keys and appropriate values to be passed to the keys like `name`, `view` corresponding to object.
From the below list of supported arguments for filters, use only the searchable fields for retriving the matching records.

### Supported Arguments for filters

---

| Field   | Alias    | Type   | Searchable |
| ------- | -------- | ------ | ---------- |
| name    | fqdn     | string | yes        |
| text    | text     | string | yes        |
| view    | dns_view | string | yes        |
| zone    | zone     | string | yes        |
| ttl     | ttl      | uint   | no         |
| comment | comment  | string | yes        |

!> Any of the combination from searchable fields in supported arguments list for fields are allowed.

!> Please consider using only fields as the keys in terraform datasource filters, kindly don't use alias names as keys from the above table.

### Example for using the filters

```hcl
data "infoblox_txt_record" "txt_rec_filter" {
  filters = {
    name = "sampletxt.demo.com"
    text = "\"some random text\""
    view = "default" // associated DNS view
  }
}
```

!> From the above example, if the 'view' alias 'dns_view' value is not specified, if same record exists in one or more different DNS views, those
all records will be fetched in results.

!> If `null` or empty filters are passed, then all the records or objects associated with datasource like here `infoblox_txt_record` will be fetched in results.

### Example of the TXT-record Data Source Block

```hcl
resource "infoblox_txt_record" "rec3" {
  dns_view = "nondefault_dnsview1"
  fqdn     = "example3.example2.org"
  text     = "\"data for TXT-record #3\""
  ttl      = 300
  comment  = "example TXT record #3"

  ext_attrs = jsonencode({
    "Location" = "65.8665701230204, -37.00791763398113"
  })
}

data "infoblox_txt_record" "ds3" {
  filters = {
    view = "nondefault_dnsview1"
    name = "example3.example2.org"
  }

  // This is just to ensure that the record has been be created
  // using 'infoblox_txt_record' resource block before the data source will be queried.
  depends_on = [infoblox_txt_record.rec3]
}

output "txt_rec_res" {
  value = data.infoblox_txt_record.ds3
}

// accessing individual field in results
output "txt_rec_mes" {
  value = data.infoblox_txt_record.ds3.results.0.text //zero represents index of json object from results list
}

// accessing TXT-Record through EA's
data "infoblox_txt_record" "txt_rec_ea" {
  filters = {
    "*Location" = "Unknown"
  }
}

// throws matching TXT-Records with EA, if any
output "txt_rec_out" {
  value = data.infoblox_txt_record.txt_rec_ea
}
```
