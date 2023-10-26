# MX-record Data Source

Use the data source to retrieve the following information for MX-record from the corresponding object in NIOS:

* `dns_view`: the DNS view which the record's zone belongs to.
* `fqdn`: the DNS zone (as a fully qualified domain name) which a mail exchange host is assigned to. Example: `samplemx.demo.com`
* `mail_exchanger`: the mail exchange host's fully qualified domain name. Example: `mx1.secure-mail-provider.net`
* `preference`: the preference number (0-65535) for this MX-record.
* `zone`: the zone which the record belongs to.
* `ttl`: the "time to live" value of the record, in seconds. Example: `1800`.
* `comment`: the description of the record. This is a regular comment. Example: `spare node for the service`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as stirng of JSON map. Example: `"{\"Owner\":\"State Library\", \"Expires\":\"never\"}"`.

For usage of filters, add the fields as keys and appropriate values to be passed to the keys like `name`, `view` corresponding to object.
From the below list of supported arguments for filters,  use only the searchable fields for retriving the matching records.

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

!> Any of the combination from searchable fields in supported arguments list for fields are allowed.

!> Please consider using only fields as the keys in terraform datasource filters, kindly don't use alias names as keys from the above table.

### Example for using the filters:
 ```hcl
 data "infoblox_mx_record" "mx_filter" {
    filters = {
        name = "samplemx.demo.com"
        mail_exchanger = "mx1.secure-mail-provider.net"
        view = "nondefault_dnsview" // associated DNS view
    }
 }
 ```

!> From the above example, if the 'view' alias 'dns_view' value is not specified, if same record exists in one or more different DNS views, those
all records will be fetched in results.

!> If `null` or empty filters are passed, then all the records or objects associated with datasource like here `infoblox_mx_record` will be fetched in results.

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
  filters = {
    dns_view = "nondefault_dnsview1"
    fqdn = "rec2.example2.org"
    mail_exchanger = "sample.test.com"
  }

  // This is just to ensure that the record has been be created
  // using 'infoblox_mx_record' resource block before the data source will be queried.
  depends_on = [infoblox_mx_record.rec2]
}

output "mx_rec_res" {
  value = data.infoblox_mx_record.ds2
}

// accessing individual field in results
output "mx_rec_name" {
  value = data.infoblox_mx_record.ds2.results.0.fqdn //zero represents index of json object from results list
}

// accessing MX-Record through EA's
data "infoblox_mx_record" "mx_rec_ea" {
  filters = {
    "*Location" = "California"
    "*TestEA" = "automate"
  }
}

// throws matching MX-Records with EA, if any
output "mx_rec_out" {
  value = data.infoblox_mx_record.mx_rec_ea
}
```
